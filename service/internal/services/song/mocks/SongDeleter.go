// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SongDeleter is an autogenerated mock type for the SongDeleter type
type SongDeleter struct {
	mock.Mock
}

// DeleteSong provides a mock function with given fields: ctx, id
func (_m *SongDeleter) DeleteSong(ctx context.Context, id uint64) (uint64, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSong")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (uint64, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) uint64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSongDeleter creates a new instance of SongDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSongDeleter(t interface {
	mock.TestingT
	Cleanup(func())
}) *SongDeleter {
	mock := &SongDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
