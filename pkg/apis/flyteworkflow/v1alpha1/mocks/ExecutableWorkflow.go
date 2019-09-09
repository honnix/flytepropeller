// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableWorkflow is an autogenerated mock type for the ExecutableWorkflow type
type ExecutableWorkflow struct {
	mock.Mock
}

// FindSubWorkflow provides a mock function with given fields: subID
func (_m *ExecutableWorkflow) FindSubWorkflow(subID string) v1alpha1.ExecutableSubWorkflow {
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

// FromNode provides a mock function with given fields: name
func (_m *ExecutableWorkflow) FromNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnnotations provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetAnnotations() map[string]string {
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

// GetConnections provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetConnections() *v1alpha1.Connections {
	ret := _m.Called()

	var r0 *v1alpha1.Connections
	if rf, ok := ret.Get(0).(func() *v1alpha1.Connections); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Connections)
		}
	}

	return r0
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetCreationTimestamp() v1.Time {
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
func (_m *ExecutableWorkflow) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
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
func (_m *ExecutableWorkflow) GetExecutionStatus() v1alpha1.ExecutableWorkflowStatus {
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

// GetID provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetK8sWorkflowID() types.NamespacedName {
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
func (_m *ExecutableWorkflow) GetLabels() map[string]string {
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
func (_m *ExecutableWorkflow) GetName() string {
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
func (_m *ExecutableWorkflow) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNode provides a mock function with given fields: nodeID
func (_m *ExecutableWorkflow) GetNode(nodeID string) (v1alpha1.ExecutableNode, bool) {
	ret := _m.Called(nodeID)

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNode); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNodeExecutionStatus provides a mock function with given fields: id
func (_m *ExecutableWorkflow) GetNodeExecutionStatus(id string) v1alpha1.ExecutableNodeStatus {
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

// GetNodes provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetNodes() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetOnFailureNode provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetOnFailureNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

// GetOutputBindings provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetOutputBindings() []*v1alpha1.Binding {
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

// GetOutputs provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetOutputs() *v1alpha1.OutputVarMap {
	ret := _m.Called()

	var r0 *v1alpha1.OutputVarMap
	if rf, ok := ret.Get(0).(func() *v1alpha1.OutputVarMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.OutputVarMap)
		}
	}

	return r0
}

// GetOwnerReference provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *ExecutableWorkflow) GetServiceAccountName() string {
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
func (_m *ExecutableWorkflow) GetTask(id string) (v1alpha1.ExecutableTask, error) {
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

// StartNode provides a mock function with given fields:
func (_m *ExecutableWorkflow) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}