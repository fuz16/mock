package mock

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatch2(t *testing.T) {
	assert.Nil(t, Patch(yes, no))
	assert.False(t, yes())
	Unpatch(yes)
	assert.True(t, yes())
}

func TestPatch3(t *testing.T) {
	assert.Nil(t, Patch(yes, no))
	assert.False(t, yes())
	assert.Nil(t, Patch(yes, no))
	assert.False(t, yes())
	UnpatchAll()
	assert.True(t, yes())
}

type Payload struct{}

func (p Payload) Yes() bool {
	return true
}

func TestPatch4(t *testing.T) {
	p := Payload{}
	assert.Nil(t, Patch(p.Yes, no))
	assert.True(t, p.Yes())
	assert.True(t, Payload{}.Yes())
	Unpatch(p.Yes)
	assert.True(t, p.Yes())
}

func TestPatchMethod(t *testing.T) {
	p := Payload{}
	assert.Nil(t, PatchMethod(reflect.TypeOf(p), "Yes", no))
	assert.False(t, p.Yes())
	assert.Nil(t, UnpatchMethod(reflect.TypeOf(p), "Yes"))
	assert.True(t, p.Yes())
}
