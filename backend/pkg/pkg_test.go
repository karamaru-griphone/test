package pkg_test

import (
	"test/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoge(t *testing.T) {
	_, err := pkg.Hoge()
	assert.NoError(t, err)
}
