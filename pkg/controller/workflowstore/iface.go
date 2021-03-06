package workflowstore

import (
	"context"

	"github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

type PriorityClass int

const (
	PriorityClassCritical PriorityClass = iota
	PriorityClassRegular
)

type FlyteWorkflow interface {
	Get(ctx context.Context, namespace, name string) (*v1alpha1.FlyteWorkflow, error)
	UpdateStatus(ctx context.Context, workflow *v1alpha1.FlyteWorkflow, priorityClass PriorityClass) error
	Update(ctx context.Context, workflow *v1alpha1.FlyteWorkflow, priorityClass PriorityClass) error
}
