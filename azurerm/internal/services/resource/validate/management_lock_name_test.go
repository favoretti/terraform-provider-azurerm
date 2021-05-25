package validate

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestManagementLockName(t *testing.T) {
	str := acceptance.RandString(259)
	testCases := []struct {
		input       string
		shouldError bool
	}{
		{"ab", false},
		{"ABC", false},
		{"abc", false},
		{"abc123ABC", false},
		{"123abcABC", false},
		{"ABC123abc", false},
		{"abc-123", false},
		{"abc_123", false},
		{str, false},
		{str + "h", true},
	}

	for _, test := range testCases {
		_, es := ManagementLockName(test.input, "name")

		if test.shouldError && len(es) == 0 {
			t.Fatalf("Expected validating name %q to fail", test.input)
		}
	}
}
