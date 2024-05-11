package green

import (
	"testing"
	 "github.com/gpltaylor/copper-monkey/cmd/green/cmd"
)

// Add unit tests for cobra green command
// AddClientPaymentRequest

func Test_ExecuteCommand(t *testing.T) {
	cmdRoot := green.NewRootCmd()
	cmdRoot.SetArgs([]string{"AddClientPaymentRequest", "--amount", "12.99"})
	cmdRoot.Execute()

	// throw and exception
	t.Error("Test failed")
}

