package internal

import (
	"database/sql"
	"encoding/json"
)

type Schema struct {
	Name   string   `json:"name"`
	Tables []*Table `json:"tables"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
	Keys    []Key    `json:"keys"`
	// IsInsertableInto bool      `json:"is_insertable_into"`
}

type Column struct {
	Name       string `json:"name"`
	DataType   string `json:"data_type"`
	IsNullable bool   `json:"is_nullable"`
	Default    string `json:"default"`
}

type Key struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func Introspect(db *sql.Tx, schema string) ([]Table, error) {
	var (
		tables []Table
		b      []byte
		err    error
	)
	err = db.QueryRow(tableSQL, schema).Scan(&b)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &tables)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

const tableSQL = `
WITH all_tables AS (
	SELECT table_schema, table_name FROM information_schema.tables
	WHERE table_schema = $1
),
all_cols AS (
	SELECT c.table_schema, c.table_name, c.column_name, c.is_nullable = 'YES' is_nullable, c.column_default,
		CASE WHEN c.data_type = 'USER-DEFINED' THEN c.udt_name ELSE c.data_type END as data_type
	FROM information_schema.columns c
	WHERE table_schema = $1
	ORDER BY table_name, ordinal_position
),
all_keys AS (
	SELECT * FROM information_schema.table_constraints
	WHERE table_schema = $1 AND constraint_type = 'PRIMARY KEY' OR constraint_type = 'UNIQUE'
)
SELECT json_agg(json_build_object(
	'schema', table_schema,
	'name', table_name,
	'columns', (
		SELECT json_agg(json_build_object(
			'name', c.column_name,
			'data_type', c.data_type,
			'is_nullable', c.is_nullable,
			'default', c.column_default
		))
		FROM all_cols c
		WHERE t.table_schema = c.table_schema AND t.table_name = c.table_name
	),
	'keys', (
		SELECT json_agg(json_build_object(
			'name', k.constraint_name,
			'columns', (
				SELECT json_agg(json_build_object(
					'name', c.column_name,
					'data_type', c.data_type,
					'is_nullable', c.is_nullable,
					'default', c.column_default
				))
				FROM information_schema.constraint_column_usage kc
				INNER JOIN all_cols c ON kc.column_name = c.column_name
					AND kc.table_schema = c.table_schema AND kc.table_name = c.table_name
				WHERE kc.constraint_name = k.constraint_name
			)
		))
		FROM all_keys k
		WHERE k.table_schema = t.table_schema AND k.table_name = t.table_name
	)
))
FROM all_tables t
`
