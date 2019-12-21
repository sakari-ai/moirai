package promise

import (
	"github.com/capitalone/go-future-context"
)

type Promise struct {
	promise future.Interface
	err     error
}

func Resolve(data interface{}) *Promise {
	return &Promise{
		promise: future.New(func() (i interface{}, e error) {
			return data, nil
		}),
	}
}

func (p *Promise) Then(a func(interface{}) (i interface{}, e error)) *Promise {
	p.promise = p.promise.Then(func(data interface{}) (i interface{}, e error) {
		return a(data)
	})
	return p
}

func (p *Promise) Catch(a func(err error)) *Promise {
	_, err := p.promise.Get()
	if err != nil && p.err == nil {
		a(err)
		p.err = err
	}
	return p
}

func (p *Promise) Await() (interface{}, error) {
	return p.promise.Get()
}
