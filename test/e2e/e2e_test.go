package e2e

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestE2E(t *testing.T) {
	const libdir = "../data/library/multi-level"
	dir, err := os.ReadDir(libdir)
	assert.Nil(t, err)
	for _, file := range dir {
		fmt.Println(file.Name(), file.IsDir())
	}
}
