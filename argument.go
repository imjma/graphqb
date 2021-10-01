package graphqb

import (
	"fmt"
	"strings"
)

type Argument struct {
	Name  string
	Value interface{}
}

type ArgumentVariable string

func NewArgument(name string, value interface{}) *Argument {
	return &Argument{
		Name:  name,
		Value: value,
	}
}

func (a *Argument) Stringify() string {
	var str string
	switch v := a.Value.(type) {
	case bool:
		str = fmt.Sprintf("%t", v)
	case int:
		str = fmt.Sprintf("%d", v)
	case string:
		str = fmt.Sprintf(`"%s"`, v)
	case ArgumentVariable:
		str = fmt.Sprintf(`$%s`, v)
	default:
		return ""
	}

	var sb strings.Builder

	sb.WriteString(a.Name)
	sb.WriteString(qColumn)
	sb.WriteString(str)

	return sb.String()
}
