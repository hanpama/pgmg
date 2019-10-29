package dsl

func With(alias string) cteAlias { return cteAlias(alias) }

func WithRecursive(alias string) {}

type cteAlias string

func (a cteAlias) As(query SelectQuery) CTEList {
	return nil
}

type CTEList []CTE

func (c CTEList) Select(selection ...interface{}) *SelectQuery { return &SelectQuery{} }

type CTE struct {
	Alias     string
	Recursive bool
	Query     SelectQuery
}
