package app

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	// setup test suite & execution environment
	ts := &testsuite.WorkflowTestSuite{}
	env := ts.NewTestWorkflowEnvironment()

	// mock ComposeGreeting when run with anything, string will return "Hello World", nil
	env.OnActivity(ComposeGreeting, mock.Anything, "World").Return("Hello World", nil)

	// prepare result
	var greetings string

	// run the workflow
	env.ExecuteWorkflow(GreetingWorkflow, "World")

	// assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	require.NoError(t, env.GetWorkflowResult(&greetings))
	require.Equal(t, "Hello World", greetings)
}
