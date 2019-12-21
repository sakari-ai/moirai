package concurrency

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
)

const (
	MaxWorkers      = 8
	MaxQueueSize    = 256
	MasterQueueSize = 8 * MaxQueueSize * MaxWorkers
)

type Pipeline struct {
	workers map[int]*worker
	chain   chan interface{}
}

func (p *Pipeline) Start(ctx context.Context) {
	go func(pipe *Pipeline) {
		for {
			expectationWorkers := len(pipe.chain) % MaxWorkers
			if expectationWorkers >= MaxWorkers {
				expectationWorkers = 0
			}
			select {
			case <-ctx.Done():
				return
			case val, ok := <-pipe.chain:
				if !ok {
					return
				}
				go pipe.workers[expectationWorkers].stream(val)
			}
		}
	}(p)
}

func (p *Pipeline) Dispatch(msg interface{}) {
	p.chain <- msg
}

type DispatcherBuilder func(*Pipeline) Dispatcher

func NewPipeline(d DispatcherBuilder, idle uint32, debug bool) *Pipeline {
	ch := make(chan interface{}, MasterQueueSize)
	wk := make(map[int]*worker)
	p := &Pipeline{workers: wk, chain: ch}
	for i := 0; i < MaxWorkers; i++ {
		wk[i] = &worker{
			index:      uint32(i + 1),
			chain:      make(chan interface{}, MaxQueueSize),
			mutex:      new(sync.RWMutex),
			debug:      debug,
			idle:       idle,
			Dispatcher: d(p),
		}
	}
	return p
}

type Dispatcher interface {
	Before(context.Context) error
	After() error
	Process(interface{}) error
	ErrorProcessor(interface{}, error)
}

type worker struct {
	index   uint32
	mutex   *sync.RWMutex
	running bool
	chain   chan interface{}
	debug   bool
	idle    uint32
	Dispatcher
}

func (c *worker) stream(val interface{}) {
	c.chain <- val
	if !c.running {
		c.mutex.RLock()
		c.running = true
		var idle uint32 = 0
		ctx, cancel := context.WithCancel(context.Background())
		var processedError error
		defer func(w *worker, cancel context.CancelFunc, processed error, msg interface{}) {
			err := c.After()
			if err != nil {
				log.Error("can not finish track issue", field.Error(err))
			}
			w.running = false
			if processedError != nil {
				w.ErrorProcessor(msg, processedError)
			}
			cancel()
			w.mutex.RUnlock()
		}(c, cancel, processedError, val)
		err := c.Before(ctx)

		if err != nil {
			processedError = err
			return
		}
		for {
			select {
			case msg, ok := <-c.chain:
				if !ok {
					return
				}
				idle = 0
				if msg != nil {
					err := c.Process(msg)
					if err != nil {
						log.Error("can not process message",
							field.Any("msg", &msg),
							field.Error(err),
						)
						val = msg
						processedError = err
					}
				}
			default:
				idle++
				if idle > c.idle {
					if len(c.chain) == 0 {
						if c.debug {
							log.Info("Worker leaving", field.Any("index", c.index), field.Any("idle", idle))
						}
						return
					}
					atomic.StoreUint32(&idle, 0)
				}
			}
		}
	}
}
