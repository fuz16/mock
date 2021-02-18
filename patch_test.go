package mock

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func no() bool  { return false }
func yes() bool { return true }

func TestPatch(t *testing.T) {
	p, err := newPatch(reflect.ValueOf(yes), reflect.ValueOf(no))
	assert.Nil(t, err)
	p.apply()
	assert.False(t, yes())
	p.undo()
	assert.True(t, yes())
}
