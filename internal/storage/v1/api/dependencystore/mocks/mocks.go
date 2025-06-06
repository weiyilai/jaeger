// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

package mocks

import (
	"context"
	"time"

	"github.com/jaegertracing/jaeger-idl/model/v1"
	mock "github.com/stretchr/testify/mock"
)

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

type Reader_Expecter struct {
	mock *mock.Mock
}

func (_m *Reader) EXPECT() *Reader_Expecter {
	return &Reader_Expecter{mock: &_m.Mock}
}

// GetDependencies provides a mock function for the type Reader
func (_mock *Reader) GetDependencies(ctx context.Context, endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	ret := _mock.Called(ctx, endTs, lookback)

	if len(ret) == 0 {
		panic("no return value specified for GetDependencies")
	}

	var r0 []model.DependencyLink
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, time.Time, time.Duration) ([]model.DependencyLink, error)); ok {
		return returnFunc(ctx, endTs, lookback)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, time.Time, time.Duration) []model.DependencyLink); ok {
		r0 = returnFunc(ctx, endTs, lookback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DependencyLink)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, time.Time, time.Duration) error); ok {
		r1 = returnFunc(ctx, endTs, lookback)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Reader_GetDependencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDependencies'
type Reader_GetDependencies_Call struct {
	*mock.Call
}

// GetDependencies is a helper method to define mock.On call
//   - ctx context.Context
//   - endTs time.Time
//   - lookback time.Duration
func (_e *Reader_Expecter) GetDependencies(ctx interface{}, endTs interface{}, lookback interface{}) *Reader_GetDependencies_Call {
	return &Reader_GetDependencies_Call{Call: _e.mock.On("GetDependencies", ctx, endTs, lookback)}
}

func (_c *Reader_GetDependencies_Call) Run(run func(ctx context.Context, endTs time.Time, lookback time.Duration)) *Reader_GetDependencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 time.Time
		if args[1] != nil {
			arg1 = args[1].(time.Time)
		}
		var arg2 time.Duration
		if args[2] != nil {
			arg2 = args[2].(time.Duration)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *Reader_GetDependencies_Call) Return(dependencyLinks []model.DependencyLink, err error) *Reader_GetDependencies_Call {
	_c.Call.Return(dependencyLinks, err)
	return _c
}

func (_c *Reader_GetDependencies_Call) RunAndReturn(run func(ctx context.Context, endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error)) *Reader_GetDependencies_Call {
	_c.Call.Return(run)
	return _c
}
