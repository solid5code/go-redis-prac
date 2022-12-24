// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	controller "go-redis-prac/src/controller"

	mock "github.com/stretchr/testify/mock"
)

// IController is an autogenerated mock type for the IController type
type IController struct {
	mock.Mock
}

// Redis provides a mock function with given fields:
func (_m *IController) Redis() controller.IRedisController {
	ret := _m.Called()

	var r0 controller.IRedisController
	if rf, ok := ret.Get(0).(func() controller.IRedisController); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(controller.IRedisController)
		}
	}

	return r0
}

type mockConstructorTestingTNewIController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIController creates a new instance of IController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIController(t mockConstructorTestingTNewIController) *IController {
	mock := &IController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}