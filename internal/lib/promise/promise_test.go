package promise

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThen(t *testing.T) {
	type args struct {
		data      interface{}
		calculate *calculate
	}
	tests := []struct {
		name        string
		args        args
		wantResult  int
		wantIsError bool
	}{
		{
			name: "#1 Then Process Calculate Plus success",
			args: args{
				data: 3,
				calculate: &calculate{
					err: nil,
					processAddition: func(data interface{}) (i interface{}, e error) {
						return data.(int) + 2, nil
					},
				},
			},
			wantResult:  5,
			wantIsError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promise := Resolve(tt.args.data).Then(tt.args.calculate.calculateAddition).Catch(tt.args.calculate.logError)
			result, err := promise.Await()
			assert.NoError(t, tt.args.calculate.err, "No processAddition catch")
			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantIsError, err != nil)
		})
	}
}

func TestCatch(t *testing.T) {
	type args struct {
		data      interface{}
		calculate *calculate
	}
	tests := []struct {
		name        string
		args        args
		wantResult  interface{}
		wantIsError bool
	}{
		{
			name: "#1 Then Process Calculate Plus success",
			args: args{
				data: 3,
				calculate: &calculate{
					err: nil,
					processAddition: func(i2 interface{}) (i interface{}, e error) {
						return nil, errors.New("have error processAddition")
					},
				},
			},
			wantResult:  nil,
			wantIsError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promise := Resolve(tt.args.data).Then(tt.args.calculate.calculateAddition).Catch(tt.args.calculate.logError)
			result, err := promise.Await()
			assert.Error(t, tt.args.calculate.err, "processAddition catch success")
			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantIsError, err != nil)
		})
	}
}

func TestChain(t *testing.T) {
	type args struct {
		data      interface{}
		calculate *calculate
	}
	tests := []struct {
		name        string
		args        args
		wantResult  interface{}
		wantIsError bool
	}{
		{
			name: "#1 Chain of responsibility pass processAddition then processSub then processMultiplication" +
				" then processDivision and not switch catch",
			args: args{
				data: 50,
				calculate: &calculate{
					err: nil,
					processAddition: func(data interface{}) (i interface{}, e error) {
						return data.(int) + 2, nil
					},
					processSub: func(data interface{}) (i interface{}, e error) {
						return data.(int) - 10, nil
					},
					processMultiplication: func(data interface{}) (i interface{}, e error) {
						return data.(int) * 2, nil
					},
					processDivision: func(data interface{}) (i interface{}, e error) {
						return data.(int) / 2, nil
					},
				},
			},
			wantIsError: false,
			wantResult:  (((50 + 2) - 10) * 2) / 2,
		},
		{
			name: "#2 Chain of responsibility pass processAddition then processSub then processMultiplication" +
				" then processDivision and it will switch catch at process Multiplication",
			args: args{
				data: 50,
				calculate: &calculate{
					err: nil,
					processAddition: func(data interface{}) (i interface{}, e error) {
						return data.(int) + 2, nil
					},
					processSub: func(data interface{}) (i interface{}, e error) {
						return data.(int) - 10, nil
					},
					processMultiplication: func(data interface{}) (i interface{}, e error) {
						return "catch here", errors.New("have error processAddition")
					},
					processDivision: func(data interface{}) (i interface{}, e error) {
						return data.(int) / 2, nil
					},
				},
			},
			wantIsError: true,
			wantResult:  "catch here",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promise := Resolve(tt.args.data).
				Then(tt.args.calculate.calculateAddition).
				Catch(tt.args.calculate.logError).
				Then(tt.args.calculate.calculateSub).
				Catch(tt.args.calculate.logError).
				Then(tt.args.calculate.calculateMultiplication).
				Catch(tt.args.calculate.logError).
				Then(tt.args.calculate.calculateDivision).
				Catch(tt.args.calculate.logError)
			result, err := promise.Await()
			assert.Equal(t, tt.wantIsError, tt.args.calculate.err != nil)
			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantIsError, err != nil)

		})
	}
}

type calculate struct {
	err                   error
	processAddition       func(interface{}) (interface{}, error)
	processSub            func(interface{}) (interface{}, error)
	processMultiplication func(interface{}) (interface{}, error)
	processDivision       func(interface{}) (interface{}, error)
}

func (c *calculate) calculateAddition(data interface{}) (interface{}, error) {
	return c.processAddition(data)
}

func (c *calculate) calculateSub(data interface{}) (interface{}, error) {
	return c.processSub(data)
}

func (c *calculate) calculateMultiplication(data interface{}) (interface{}, error) {
	return c.processMultiplication(data)
}

func (c *calculate) calculateDivision(data interface{}) (interface{}, error) {
	return c.processDivision(data)
}

func (c *calculate) logError(err error) {
	c.err = err
}
