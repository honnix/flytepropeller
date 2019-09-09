// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableBranchNodeStatus is an autogenerated mock type for the ExecutableBranchNodeStatus type
type ExecutableBranchNodeStatus struct {
	mock.Mock
}

// GetFinalizedNode provides a mock function with given fields:
func (_m *ExecutableBranchNodeStatus) GetFinalizedNode() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// GetPhase provides a mock function with given fields:
func (_m *ExecutableBranchNodeStatus) GetPhase() v1alpha1.BranchNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.BranchNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.BranchNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.BranchNodePhase)
	}

	return r0
}