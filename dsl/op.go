package dsl

func CastAs(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("CAST(%s AS %s)", expr1, expr2)
}

type caseExpression struct {
	Expression Expression
	Whens      []caseWhen
	Else       Expression
}

type caseWhen struct {
	When Expression
	Then Expression
}

func Case(expr Expression) {
	panic("Unimplemented")
}
func CaseWhen() {
	panic("Unimplemented")
}

func NullIf(a interface{}, b interface{}) *FuncExpr {
	return &FuncExpr{"NULLIF", WrapExprs(a, b)}
}
func Coalesce(exprs ...interface{}) *FuncExpr {
	return &FuncExpr{"COALESCE", WrapExprs(exprs...)}
}
func Greatest(exprs ...interface{}) *FuncExpr {
	return &FuncExpr{"GREATEST", WrapExprs(exprs...)}
}
func Least(exprs ...interface{}) *FuncExpr {
	return &FuncExpr{"LEAST", WrapExprs(exprs...)}
}

// Comparison Operators
// https://www.postgresql.org/docs/current/functions-comparison.html

func (c *Exprf) Eq(expr interface{}) *Exprf  { return NewExprf("%s = %s", c, expr) }
func (c *Exprf) Ne(expr interface{}) *Exprf  { return NewExprf("%s != %s", c, expr) }
func (c *Exprf) Gt(expr interface{}) *Exprf  { return NewExprf("%s > %s", c, expr) }
func (c *Exprf) Lt(expr interface{}) *Exprf  { return NewExprf("%s < %s", c, expr) }
func (c *Exprf) Gte(expr interface{}) *Exprf { return NewExprf("%s >= %s", c, expr) }
func (c *Exprf) Lte(expr interface{}) *Exprf { return NewExprf("%s <= %s", c, expr) }
func (c *Exprf) Between(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN %s AND %s", c, expr1, expr2)
}
func (c *Exprf) NotBetween(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN %s AND %s", c, expr1, expr2)
}
func (c *Exprf) BetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c *Exprf) NotBetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c *Exprf) IsDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS DISTINCT FROM %s", c, expr1)
}
func (c *Exprf) IsNotDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS NOT DISTINCT FROM %s", c, expr1)
}
func (c *Exprf) IsNull() *Exprf       { return NewExprf("%s IS NULL", c) }
func (c *Exprf) IsNotNull() *Exprf    { return NewExprf("%s IS NOT NULL", c) }
func (c *Exprf) IsTrue() *Exprf       { return NewExprf("%s IS TRUE", c) }
func (c *Exprf) IsNotTrue() *Exprf    { return NewExprf("%s IS NOT TRUE", c) }
func (c *Exprf) IsFalse() *Exprf      { return NewExprf("%s IS FALSE", c) }
func (c *Exprf) IsNotFalse() *Exprf   { return NewExprf("%s IS NOT FALSE", c) }
func (c *Exprf) IsUnknown() *Exprf    { return NewExprf("%s IS UNKNOWN", c) }
func (c *Exprf) IsNotUnknown() *Exprf { return NewExprf("%s IS NOT UNKNOWN", c) }

// https://www.postgresql.org/docs/current/functions-math.html
func (c *Exprf) Add(expr interface{}) *Exprf { return NewExprf("%s + %s", c, expr) }
func (c *Exprf) Sub(expr interface{}) *Exprf { return NewExprf("%s - %s", c, expr) }
func (c *Exprf) Mul(expr interface{}) *Exprf { return NewExprf("%s * %s", c, expr) }
func (c *Exprf) Div(expr interface{}) *Exprf { return NewExprf("%s / %s", c, expr) }
func (c *Exprf) Mod(expr interface{}) *Exprf { return NewExprf("%s %% %s", c, expr) }
func (c *Exprf) Exp(expr interface{}) *Exprf { return NewExprf("%s ^ %s", c, expr) }

// String and pattern maching
// https://www.postgresql.org/docs/current/functions-matching.html

// Concat does String concatenation
func (c *Exprf) Concat(expr interface{}) *Exprf  { return NewExprf("%s || %s", c, expr) }
func (c *Exprf) Like(expr interface{}) *Exprf    { return NewExprf("%s LIKE %s", c, expr) }
func (c *Exprf) NotLike(expr interface{}) *Exprf { return NewExprf("%s NOT LIKE %s", c, expr) }
func (c *Exprf) SimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s SIMILAR TO %s", c, expr)
}
func (c *Exprf) NotSimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s NOT SIMILAR TO %s", c, expr)
}

// Comparison Operators
// https://www.postgresql.org/docs/current/functions-comparison.html

func (c ColumnReference) Eq(expr interface{}) *Exprf  { return NewExprf("%s = %s", c, expr) }
func (c ColumnReference) Ne(expr interface{}) *Exprf  { return NewExprf("%s != %s", c, expr) }
func (c ColumnReference) Gt(expr interface{}) *Exprf  { return NewExprf("%s > %s", c, expr) }
func (c ColumnReference) Lt(expr interface{}) *Exprf  { return NewExprf("%s < %s", c, expr) }
func (c ColumnReference) Gte(expr interface{}) *Exprf { return NewExprf("%s >= %s", c, expr) }
func (c ColumnReference) Lte(expr interface{}) *Exprf { return NewExprf("%s <= %s", c, expr) }
func (c ColumnReference) Between(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN %s AND %s", c, expr1, expr2)
}
func (c ColumnReference) NotBetween(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN %s AND %s", c, expr1, expr2)
}
func (c ColumnReference) BetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c ColumnReference) NotBetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c ColumnReference) IsDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS DISTINCT FROM %s", c, expr1)
}
func (c ColumnReference) IsNotDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS NOT DISTINCT FROM %s", c, expr1)
}
func (c ColumnReference) IsNull() *Exprf       { return NewExprf("%s IS NULL", c) }
func (c ColumnReference) IsNotNull() *Exprf    { return NewExprf("%s IS NOT NULL", c) }
func (c ColumnReference) IsTrue() *Exprf       { return NewExprf("%s IS TRUE", c) }
func (c ColumnReference) IsNotTrue() *Exprf    { return NewExprf("%s IS NOT TRUE", c) }
func (c ColumnReference) IsFalse() *Exprf      { return NewExprf("%s IS FALSE", c) }
func (c ColumnReference) IsNotFalse() *Exprf   { return NewExprf("%s IS NOT FALSE", c) }
func (c ColumnReference) IsUnknown() *Exprf    { return NewExprf("%s IS UNKNOWN", c) }
func (c ColumnReference) IsNotUnknown() *Exprf { return NewExprf("%s IS NOT UNKNOWN", c) }

// https://www.postgresql.org/docs/current/functions-math.html
func (c ColumnReference) Add(expr interface{}) *Exprf { return NewExprf("%s + %s", c, expr) }
func (c ColumnReference) Sub(expr interface{}) *Exprf { return NewExprf("%s - %s", c, expr) }
func (c ColumnReference) Mul(expr interface{}) *Exprf { return NewExprf("%s * %s", c, expr) }
func (c ColumnReference) Div(expr interface{}) *Exprf { return NewExprf("%s / %s", c, expr) }
func (c ColumnReference) Mod(expr interface{}) *Exprf { return NewExprf("%s %% %s", c, expr) }
func (c ColumnReference) Exp(expr interface{}) *Exprf { return NewExprf("%s ^ %s", c, expr) }

// String and pattern maching
// https://www.postgresql.org/docs/current/functions-matching.html

// Concat does String concatenation
func (c ColumnReference) Concat(expr interface{}) *Exprf  { return NewExprf("%s || %s", c, expr) }
func (c ColumnReference) Like(expr interface{}) *Exprf    { return NewExprf("%s LIKE %s", c, expr) }
func (c ColumnReference) NotLike(expr interface{}) *Exprf { return NewExprf("%s NOT LIKE %s", c, expr) }
func (c ColumnReference) SimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s SIMILAR TO %s", c, expr)
}
func (c ColumnReference) NotSimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s NOT SIMILAR TO %s", c, expr)
}

// Comparison Operators
// https://www.postgresql.org/docs/current/functions-comparison.html

func (c *FuncExpr) Eq(expr interface{}) *Exprf  { return NewExprf("%s = %s", c, expr) }
func (c *FuncExpr) Ne(expr interface{}) *Exprf  { return NewExprf("%s != %s", c, expr) }
func (c *FuncExpr) Gt(expr interface{}) *Exprf  { return NewExprf("%s > %s", c, expr) }
func (c *FuncExpr) Lt(expr interface{}) *Exprf  { return NewExprf("%s < %s", c, expr) }
func (c *FuncExpr) Gte(expr interface{}) *Exprf { return NewExprf("%s >= %s", c, expr) }
func (c *FuncExpr) Lte(expr interface{}) *Exprf { return NewExprf("%s <= %s", c, expr) }
func (c *FuncExpr) Between(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN %s AND %s", c, expr1, expr2)
}
func (c *FuncExpr) NotBetween(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN %s AND %s", c, expr1, expr2)
}
func (c *FuncExpr) BetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c *FuncExpr) NotBetweenSymmetric(expr1 interface{}, expr2 interface{}) *Exprf {
	return NewExprf("%s NOT BETWEEN SYMMETRIC %s AND %s", c, expr1, expr2)
}
func (c *FuncExpr) IsDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS DISTINCT FROM %s", c, expr1)
}
func (c *FuncExpr) IsNotDistinctFrom(expr1 interface{}) *Exprf {
	return NewExprf("%s IS NOT DISTINCT FROM %s", c, expr1)
}
func (c *FuncExpr) IsNull() *Exprf       { return NewExprf("%s IS NULL", c) }
func (c *FuncExpr) IsNotNull() *Exprf    { return NewExprf("%s IS NOT NULL", c) }
func (c *FuncExpr) IsTrue() *Exprf       { return NewExprf("%s IS TRUE", c) }
func (c *FuncExpr) IsNotTrue() *Exprf    { return NewExprf("%s IS NOT TRUE", c) }
func (c *FuncExpr) IsFalse() *Exprf      { return NewExprf("%s IS FALSE", c) }
func (c *FuncExpr) IsNotFalse() *Exprf   { return NewExprf("%s IS NOT FALSE", c) }
func (c *FuncExpr) IsUnknown() *Exprf    { return NewExprf("%s IS UNKNOWN", c) }
func (c *FuncExpr) IsNotUnknown() *Exprf { return NewExprf("%s IS NOT UNKNOWN", c) }

// https://www.postgresql.org/docs/current/functions-math.html
func (c *FuncExpr) Add(expr interface{}) *Exprf { return NewExprf("%s + %s", c, expr) }
func (c *FuncExpr) Sub(expr interface{}) *Exprf { return NewExprf("%s - %s", c, expr) }
func (c *FuncExpr) Mul(expr interface{}) *Exprf { return NewExprf("%s * %s", c, expr) }
func (c *FuncExpr) Div(expr interface{}) *Exprf { return NewExprf("%s / %s", c, expr) }
func (c *FuncExpr) Mod(expr interface{}) *Exprf { return NewExprf("%s %% %s", c, expr) }
func (c *FuncExpr) Exp(expr interface{}) *Exprf { return NewExprf("%s ^ %s", c, expr) }

// String and pattern maching
// https://www.postgresql.org/docs/current/functions-matching.html

// Concat does String concatenation
func (c *FuncExpr) Concat(expr interface{}) *Exprf  { return NewExprf("%s || %s", c, expr) }
func (c *FuncExpr) Like(expr interface{}) *Exprf    { return NewExprf("%s LIKE %s", c, expr) }
func (c *FuncExpr) NotLike(expr interface{}) *Exprf { return NewExprf("%s NOT LIKE %s", c, expr) }
func (c *FuncExpr) SimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s SIMILAR TO %s", c, expr)
}
func (c *FuncExpr) NotSimilarTo(expr interface{}) *Exprf {
	return NewExprf("%s NOT SIMILAR TO %s", c, expr)
}
