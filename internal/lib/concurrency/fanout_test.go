package concurrency

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/sakari-ai/moirai/proto"

	"github.com/stretchr/testify/mock"
)

type mockDispatcher struct {
	mock.Mock
}

func (m *mockDispatcher) Before(ctx context.Context) error {
	arg := m.Called(ctx)
	return arg.Error(0)
}

func (m *mockDispatcher) After() error {
	arg := m.Called()
	return arg.Error(0)
}

func (m *mockDispatcher) Process(msg interface{}) error {
	arg := m.Called(msg)
	return arg.Error(0)
}

func (m *mockDispatcher) ErrorProcessor(msg interface{}, err error) {
	m.Called(msg, err)
}

func Test_worker_stream(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#1: Dispatcher receive error during BeforeProcessing",
			args: args{
				val: &proto.Tagging{Topic: "error-tagging"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorIssue := fmt.Errorf("before error: #%d", 1)
			errorBeforeDispatcher := new(mockDispatcher)
			errorBeforeDispatcher.On("Before", mock.Anything).Return(errorIssue)
			errorBeforeDispatcher.On("ErrorProcessor", mock.Anything, mock.Anything)
			errorBeforeDispatcher.On("After").Return(nil)
			c := &worker{
				mutex:      new(sync.RWMutex),
				running:    false,
				chain:      make(chan interface{}, MaxQueueSize),
				debug:      true,
				idle:       100,
				Dispatcher: errorBeforeDispatcher,
			}
			c.stream(tt.args.val)
			errorBeforeDispatcher.AssertCalled(t, "ErrorProcessor",
				mock.MatchedBy(func(msg interface{}) bool {
					return reflect.DeepEqual(msg, tt.args.val)
				}),
				mock.MatchedBy(func(err error) bool {
					return reflect.DeepEqual(err, errorIssue)
				}),
			)
		})
	}
}
