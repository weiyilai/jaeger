// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

package mocks

import (
	"github.com/bsm/sarama-cluster"
	mock "github.com/stretchr/testify/mock"
)

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

type Consumer_Expecter struct {
	mock *mock.Mock
}

func (_m *Consumer) EXPECT() *Consumer_Expecter {
	return &Consumer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function for the type Consumer
func (_mock *Consumer) Close() error {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func() error); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Consumer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Consumer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Consumer_Expecter) Close() *Consumer_Close_Call {
	return &Consumer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Consumer_Close_Call) Run(run func()) *Consumer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Consumer_Close_Call) Return(err error) *Consumer_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Consumer_Close_Call) RunAndReturn(run func() error) *Consumer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// MarkPartitionOffset provides a mock function for the type Consumer
func (_mock *Consumer) MarkPartitionOffset(topic string, partition int32, offset int64, metadata string) {
	_mock.Called(topic, partition, offset, metadata)
	return
}

// Consumer_MarkPartitionOffset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkPartitionOffset'
type Consumer_MarkPartitionOffset_Call struct {
	*mock.Call
}

// MarkPartitionOffset is a helper method to define mock.On call
//   - topic string
//   - partition int32
//   - offset int64
//   - metadata string
func (_e *Consumer_Expecter) MarkPartitionOffset(topic interface{}, partition interface{}, offset interface{}, metadata interface{}) *Consumer_MarkPartitionOffset_Call {
	return &Consumer_MarkPartitionOffset_Call{Call: _e.mock.On("MarkPartitionOffset", topic, partition, offset, metadata)}
}

func (_c *Consumer_MarkPartitionOffset_Call) Run(run func(topic string, partition int32, offset int64, metadata string)) *Consumer_MarkPartitionOffset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 string
		if args[0] != nil {
			arg0 = args[0].(string)
		}
		var arg1 int32
		if args[1] != nil {
			arg1 = args[1].(int32)
		}
		var arg2 int64
		if args[2] != nil {
			arg2 = args[2].(int64)
		}
		var arg3 string
		if args[3] != nil {
			arg3 = args[3].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
			arg3,
		)
	})
	return _c
}

func (_c *Consumer_MarkPartitionOffset_Call) Return() *Consumer_MarkPartitionOffset_Call {
	_c.Call.Return()
	return _c
}

func (_c *Consumer_MarkPartitionOffset_Call) RunAndReturn(run func(topic string, partition int32, offset int64, metadata string)) *Consumer_MarkPartitionOffset_Call {
	_c.Run(run)
	return _c
}

// Partitions provides a mock function for the type Consumer
func (_mock *Consumer) Partitions() <-chan cluster.PartitionConsumer {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Partitions")
	}

	var r0 <-chan cluster.PartitionConsumer
	if returnFunc, ok := ret.Get(0).(func() <-chan cluster.PartitionConsumer); ok {
		r0 = returnFunc()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan cluster.PartitionConsumer)
		}
	}
	return r0
}

// Consumer_Partitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Partitions'
type Consumer_Partitions_Call struct {
	*mock.Call
}

// Partitions is a helper method to define mock.On call
func (_e *Consumer_Expecter) Partitions() *Consumer_Partitions_Call {
	return &Consumer_Partitions_Call{Call: _e.mock.On("Partitions")}
}

func (_c *Consumer_Partitions_Call) Run(run func()) *Consumer_Partitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Consumer_Partitions_Call) Return(partitionConsumerCh <-chan cluster.PartitionConsumer) *Consumer_Partitions_Call {
	_c.Call.Return(partitionConsumerCh)
	return _c
}

func (_c *Consumer_Partitions_Call) RunAndReturn(run func() <-chan cluster.PartitionConsumer) *Consumer_Partitions_Call {
	_c.Call.Return(run)
	return _c
}
