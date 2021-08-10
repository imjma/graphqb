package graphqb

import "strings"

const (
	qComma  = ","
	qLB     = "{"
	qRB     = "}"
	qLP     = "("
	qRP     = ")"
	qColumn = ":"
)

type Query struct {
	Type   string
	Name   string
	Fields []*Field
	Err    error
}

func (q *Query) Stringify() string {
	var sb strings.Builder

	// default is query
	if len(q.Type) > 0 {
		sb.WriteString(q.Type)
	}
	sb.WriteString(qLB)
	for k, field := range q.Fields {
		if k != 0 {
			sb.WriteString(qComma)
		}
		sb.WriteString(field.Stringify())
	}
	sb.WriteString(qRB)

	return sb.String()
}

func (q *Query) SetFields(fs ...*Field) *Query {
	q.Fields = fs
	return q
}