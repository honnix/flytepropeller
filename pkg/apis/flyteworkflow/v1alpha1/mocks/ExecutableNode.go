// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableNode is an autogenerated mock type for the ExecutableNode type
type ExecutableNode struct {
	mock.Mock
}

// GetBranchNode provides a mock function with given fields:
func (_m *ExecutableNode) GetBranchNode() v1alpha1.ExecutableBranchNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableBranchNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableBranchNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableBranchNode)
		}
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *ExecutableNode) GetConfig() *v1.ConfigMap {
	ret := _m.Called()

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func() *v1.ConfigMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	return r0
}

// GetID provides a mock function with given fields:
func (_m *ExecutableNode) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetInputBindings provides a mock function with given fields:
func (_m *ExecutableNode) GetInputBindings() []*v1alpha1.Binding {
	ret := _m.Called()

	var r0 []*v1alpha1.Binding
	if rf, ok := ret.Get(0).(func() []*v1alpha1.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Binding)
		}
	}

	return r0
}

// GetKind provides a mock function with given fields:
func (_m *ExecutableNode) GetKind() v1alpha1.NodeKind {
	ret := _m.Called()

	var r0 v1alpha1.NodeKind
	if rf, ok := ret.Get(0).(func() v1alpha1.NodeKind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.NodeKind)
	}

	return r0
}

// GetOutputAlias provides a mock function with given fields:
func (_m *ExecutableNode) GetOutputAlias() []v1alpha1.Alias {
	ret := _m.Called()

	var r0 []v1alpha1.Alias
	if rf, ok := ret.Get(0).(func() []v1alpha1.Alias); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.Alias)
		}
	}

	return r0
}

// GetResources provides a mock function with given fields:
func (_m *ExecutableNode) GetResources() *v1.ResourceRequirements {
	ret := _m.Called()

	var r0 *v1.ResourceRequirements
	if rf, ok := ret.Get(0).(func() *v1.ResourceRequirements); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ResourceRequirements)
		}
	}

	return r0
}

// GetRetryStrategy provides a mock function with given fields:
func (_m *ExecutableNode) GetRetryStrategy() *v1alpha1.RetryStrategy {
	ret := _m.Called()

	var r0 *v1alpha1.RetryStrategy
	if rf, ok := ret.Get(0).(func() *v1alpha1.RetryStrategy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.RetryStrategy)
		}
	}

	return r0
}

// GetTaskID provides a mock function with given fields:
func (_m *ExecutableNode) GetTaskID() *string {
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

// GetWorkflowNode provides a mock function with given fields:
func (_m *ExecutableNode) GetWorkflowNode() v1alpha1.ExecutableWorkflowNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowNode)
		}
	}

	return r0
}

// IsEndNode provides a mock function with given fields:
func (_m *ExecutableNode) IsEndNode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsStartNode provides a mock function with given fields:
func (_m *ExecutableNode) IsStartNode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
