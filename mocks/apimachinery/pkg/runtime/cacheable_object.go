// Code generated by mockery v2.19.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// CacheableObject is an autogenerated mock type for the CacheableObject type
type CacheableObject struct {
	mock.Mock
}

// CacheEncode provides a mock function with given fields: id, encode, w
func (_m *CacheableObject) CacheEncode(id runtime.Identifier, encode func(runtime.Object, io.Writer) error, w io.Writer) error {
	ret := _m.Called(id, encode, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(runtime.Identifier, func(runtime.Object, io.Writer) error, io.Writer) error); ok {
		r0 = rf(id, encode, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetObject provides a mock function with given fields:
func (_m *CacheableObject) GetObject() runtime.Object {
	ret := _m.Called()

	var r0 runtime.Object
	if rf, ok := ret.Get(0).(func() runtime.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Object)
		}
	}

	return r0
}

type mockConstructorTestingTNewCacheableObject interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheableObject creates a new instance of CacheableObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheableObject(t mockConstructorTestingTNewCacheableObject) *CacheableObject {
	mock := &CacheableObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
