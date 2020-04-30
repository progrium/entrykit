package codep

import (
	"testing"
)

func TestCodepNonExistent(t *testing.T) {
	Codep(map[string]string{"non-existent1": "non-existent1", "non-existent2": "non-existent2"})
}
