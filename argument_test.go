package graphqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgument_Stringify(t *testing.T) {
	assert.Equal(t, "argBool:true", NewArgument("argBool", true).Stringify())
	assert.Equal(t, "argInt:123", NewArgument("argInt", 123).Stringify())
	assert.Equal(t, `argString:"str"`, NewArgument("argString", "str").Stringify())

}
