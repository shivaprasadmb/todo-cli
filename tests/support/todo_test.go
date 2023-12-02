package support

import (
	"testing"

	"github.com/shivaprasadmb/todo-cli"
)

func TestCompleteFunction(t *testing.T) {
	// Create a sample Todos slice
	todos := todo.Todos{
		{Task: "Task 1", Done: false},
		{Task: "Task 2", Done: false},
		{Task: "Task 3", Done: false},
	}

	// Call the Complete method
	err := todos.Complete(2)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the task is marked as done
	if !todos[1].Done {
		t.Error("Expected task to be marked as done, but it's not.")
	}
	t.Logf("passed!!!")
}
