package graphqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField_Stringify(t *testing.T) {
	f := &Field{
		Name: "field",
	}
	assert.Equal(t, "field", f.Stringify())

	f.SetFields(NewField("f1"), NewField("f2"))
	f.SetArguments(NewArgument("argBool", true), NewArgument("argInt", 123), NewArgument("argString", "str"))
	assert.Equal(t, `field(argBool:true,argInt:123,argString:"str"){f1,f2}`, f.Stringify())
}
