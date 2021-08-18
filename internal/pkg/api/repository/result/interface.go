package result

import (
	"context"

	"github.com/kubeshop/kubetest/pkg/api/kubetest"
)

type Repository interface {
	// Get gets execution result by id
	Get(ctx context.Context, id string) (kubetest.ScriptExecution, error)
	// GetByName gets execution result by name
	GetByNameAndScript(ctx context.Context, name, script string) (kubetest.ScriptExecution, error)
	// GetNewestExecutions gets top X newest executions
	GetNewestExecutions(ctx context.Context, limit int) ([]kubetest.ScriptExecution, error)
	// GetScriptExecutions gets executions for given script ID
	GetScriptExecutions(ctx context.Context, scriptID string) ([]kubetest.ScriptExecution, error)
	// Insert inserts new execution result
	Insert(ctx context.Context, result kubetest.ScriptExecution) error
	// Update updates execution result
	Update(ctx context.Context, result kubetest.ScriptExecution) error
}