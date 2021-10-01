package graphqb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Stringify(t *testing.T) {
	q := &Query{
		Type: "query",
	}

	q.SetFields(
		NewField("f1").SetFields(
			NewField("f1a"),
			NewField("f1b"),
		),
		NewField("f2").SetArguments(
			&Argument{Name: "arg", Value: "str"},
		).SetFields(
			NewField("f2a"),
			NewField("f2b"),
		),
	)

	assert.Equal(t, `query{f1{f1a,f1b},f2(arg:"str"){f2a,f2b}}`, q.Stringify())
}

func TestQuery_Variable_Stringify(t *testing.T) {
	q := &Query{
		Type: "query",
		Name: "var",
	}

	var arg ArgumentVariable = "argVar"

	q.SetFields(
		NewField("f1").SetFields(
			NewField("f1a"),
			NewField("f1b"),
		),
		NewField("f2").SetArguments(
			&Argument{Name: "arg", Value: arg},
		).SetFields(
			NewField("f2a"),
			NewField("f2b"),
		),
	)

	assert.Equal(t, `query var{f1{f1a,f1b},f2(arg:$argVar){f2a,f2b}}`, q.Stringify())
}
