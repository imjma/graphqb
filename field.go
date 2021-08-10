package graphqb

import "strings"

type Field struct {
	Name      string
	Arguments []*Argument
	Fields    []*Field
	Err       error
}

func NewField(name string) *Field {
	return &Field{Name: name}
}

func (f *Field) Stringify() string {
	var sb strings.Builder

	sb.WriteString(f.Name)

	if len(f.Arguments) > 0 {
		sb.WriteString(qLP)

		for k, argument := range f.Arguments {
			if k != 0 {
				sb.WriteString(qComma)
			}
			sb.WriteString(argument.Stringify())
		}

		sb.WriteString(qRP)
	}

	if len(f.Fields) > 0 {
		sb.WriteString(qLB)

		for k, field := range f.Fields {
			if k != 0 {
				sb.WriteString(qComma)
			}
			sb.WriteString(field.Stringify())
		}

		sb.WriteString(qRB)
	}

	return sb.String()
}

func (f *Field) SetArguments(args ...*Argument) *Field {
	f.Arguments = args
	return f
}

func (f *Field) SetFields(fs ...*Field) *Field {
	f.Fields = fs
	return f
}
