// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
import mock "github.com/stretchr/testify/mock"

// ExecutableTask is an autogenerated mock type for the ExecutableTask type
type ExecutableTask struct {
	mock.Mock
}

// CoreTask provides a mock function with given fields:
func (_m *ExecutableTask) CoreTask() *core.TaskTemplate {
	ret := _m.Called()

	var r0 *core.TaskTemplate
	if rf, ok := ret.Get(0).(func() *core.TaskTemplate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskTemplate)
		}
	}

	return r0
}

// TaskType provides a mock function with given fields:
func (_m *ExecutableTask) TaskType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}