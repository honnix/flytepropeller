// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// WorkflowMetaExtended is an autogenerated mock type for the WorkflowMetaExtended type
type WorkflowMetaExtended struct {
	mock.Mock
}

// FindSubWorkflow provides a mock function with given fields: subID
func (_m *WorkflowMetaExtended) FindSubWorkflow(subID string) v1alpha1.ExecutableSubWorkflow {
	ret := _m.Called(subID)

	var r0 v1alpha1.ExecutableSubWorkflow
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableSubWorkflow); ok {
		r0 = rf(subID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableSubWorkflow)
		}
	}

	return r0
}

// GetAnnotations provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

// GetExecutionID provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowExecutionIdentifier
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowExecutionIdentifier)
	}

	return r0
}

// GetExecutionStatus provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetExecutionStatus() v1alpha1.ExecutableWorkflowStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowStatus)
		}
	}

	return r0
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetK8sWorkflowID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

// GetLabels provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNamespace provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *WorkflowMetaExtended) GetServiceAccountName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTask provides a mock function with given fields: id
func (_m *WorkflowMetaExtended) GetTask(id string) (v1alpha1.ExecutableTask, error) {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableTask
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewControllerRef provides a mock function with given fields:
func (_m *WorkflowMetaExtended) NewControllerRef() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}
