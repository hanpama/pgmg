package introspect

type Schema struct {
	SchemaName  string        `json:"schema_name"`
	Tables      []*Table      `json:"tables"`
	Columns     []*Column     `json:"columns"`
	Keys        []*Key        `json:"keys"`
	ForeignKeys []*ForeignKey `json:"foreign_keys"`
	Types       []*Type       `json:"types"`
}

type Table struct {
	TableName  string `json:"table_name"`
	Insertable bool   `json:"insertable"`
}

type Column struct {
	TableName   string `json:"table_name"`
	ColumnName  string `json:"column_name"`
	DataType    string `json:"data_type"`
	IsNullable  bool   `json:"is_nullable"`
	IsUpdatable bool   `json:"is_updatable"`
	Default     string `json:"default"`
}

type Key struct {
	TableName      string   `json:"table_name"`
	ConstraintName string   `json:"constraint_name"`
	ColumnNames    []string `json:"column_names"`
	IsPrimaryKey   bool     `json:"is_primary_key"`
}

type ForeignKey struct {
	TableName            string   `json:"table_name"`
	ConstraintName       string   `json:"constraint_name"`
	ColumnNames          []string `json:"column_names"`
	TargetTableName      string   `json:"target_table_name"`
	TargetConstraintName string   `json:"target_constraint_name"`
}

type Type struct {
	SQLType      string `json:"sql_type"`
	Name         string `json:"name"`
	NullableName string `json:"nullable_name"`
	Module       string `json:"module"`
}

var Types = []*Type{
	{SQLType: "boolean", Name: "bool", NullableName: "*bool"},
	{SQLType: "character", Name: "string", NullableName: "*string"},
	{SQLType: "character varying", Name: "string", NullableName: "*string"},
	{SQLType: "text", Name: "string", NullableName: "*string"},
	{SQLType: "money", Name: "string", NullableName: "*string"},
	{SQLType: "inet", Name: "string", NullableName: "*string"},
	{SQLType: "uuid", Name: "string", NullableName: "*string"},
	{SQLType: "jsonb", Name: "string", NullableName: "*string"},
	{SQLType: "json", Name: "string", NullableName: "*string"},
	{SQLType: "smallint", Name: "int16", NullableName: "*int16"},
	{SQLType: "integer", Name: "int32", NullableName: "*int32"},
	{SQLType: "bigint", Name: "int64", NullableName: "*int64"},
	{SQLType: "smallserial", Name: "uint16", NullableName: "*uint16"},
	{SQLType: "serial", Name: "uint32", NullableName: "*uint32"},
	{SQLType: "bigserial", Name: "uint64", NullableName: "*uint64"},
	{SQLType: "real", Name: "float32", NullableName: "*float32"},
	{SQLType: "numeric", Name: "float64", NullableName: "*float64"},
	{SQLType: "double precision", Name: "float64", NullableName: "*float64"},
	{SQLType: "bytea", Name: "[]byte", NullableName: "[]byte"},
	{SQLType: "date", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
	{SQLType: "timestamp with time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
	{SQLType: "time with time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
	{SQLType: "time without time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
	{SQLType: "timestamp without time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
	{SQLType: "interval", Name: "time.Duration", NullableName: "*time.Duration", Module: "time"},
	{SQLType: "char", Name: "uint8", NullableName: "*uint8"},
	{SQLType: "bit", Name: "uint8", NullableName: "*uint8"},
	{SQLType: "any", Name: "[]byte", NullableName: "[]byte"},
	{SQLType: "name", Name: "string", NullableName: "*string"},
	{SQLType: "bit varying", Name: "[]byte", NullableName: "[]byte"},
	{SQLType: "hstore", Name: "[]byte", NullableName: "[]byte"},
	{SQLType: "geometry", Name: "string", NullableName: "*string"},  // EWKBHex
	{SQLType: "geography", Name: "string", NullableName: "*string"}, // EWKBHex
}
