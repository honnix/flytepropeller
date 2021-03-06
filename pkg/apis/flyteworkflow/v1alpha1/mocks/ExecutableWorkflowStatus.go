// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/lyft/flytestdlib/storage"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableWorkflowStatus is an autogenerated mock type for the ExecutableWorkflowStatus type
type ExecutableWorkflowStatus struct {
	mock.Mock
}

// ConstructNodeDataDir provides a mock function with given fields: ctx, constructor, name
func (_m *ExecutableWorkflowStatus) ConstructNodeDataDir(ctx context.Context, constructor storage.ReferenceConstructor, name string) (storage.DataReference, error) {
	ret := _m.Called(ctx, constructor, name)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context, storage.ReferenceConstructor, string) storage.DataReference); ok {
		r0 = rf(ctx, constructor, name)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.ReferenceConstructor, string) error); ok {
		r1 = rf(ctx, constructor, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataDir provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetDataDir() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// GetLastUpdatedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetLastUpdatedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetMessage provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetMessage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNodeExecutionStatus provides a mock function with given fields: id
func (_m *ExecutableWorkflowStatus) GetNodeExecutionStatus(id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

// GetOutputReference provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetOutputReference() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// GetPhase provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetPhase() v1alpha1.WorkflowPhase {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowPhase
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowPhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowPhase)
	}

	return r0
}

// GetStartedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetStartedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetStoppedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetStoppedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// IncFailedAttempts provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) IncFailedAttempts() {
	_m.Called()
}

// IsTerminated provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) IsTerminated() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *ExecutableWorkflowStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetMessage provides a mock function with given fields: msg
func (_m *ExecutableWorkflowStatus) SetMessage(msg string) {
	_m.Called(msg)
}

// SetOutputReference provides a mock function with given fields: reference
func (_m *ExecutableWorkflowStatus) SetOutputReference(reference storage.DataReference) {
	_m.Called(reference)
}

// UpdatePhase provides a mock function with given fields: p, msg
func (_m *ExecutableWorkflowStatus) UpdatePhase(p v1alpha1.WorkflowPhase, msg string) {
	_m.Called(p, msg)
}
