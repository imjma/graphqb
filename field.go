package graphqb

import "strings"

type Field struct {
	Name      string
	Alias     string
	Arguments []*Argument
	Fields    []*Field
	Err       error
}

func NewField(name string) *Field {
	return &Field{Name: name}
}

func (f *Field) Stringify() string {
	var sb strings.Builder

	if len(f.Alias) > 0 {
		sb.WriteString(f.Alias + qColumn)
	}

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

func (f *Field) AddField(fs *Field) *Field {
	f.Fields = append(f.Fields, fs)
	return f
}

func (f *Field) SetAlias(alias string) *Field {
	f.Alias = alias
	return f
}
