package introspect

import (
	"database/sql"
	"encoding/json"
)

func IntrospectSchema(db *sql.DB, schemaName string) (schema *Schema, err error) {
	var b []byte
	err = db.QueryRow(tableSQL, schemaName).Scan(&b)
	if err != nil {
		return nil, err
	}
	schema = new(Schema)
	err = json.Unmarshal(b, &schema)
	if err != nil {
		return nil, err
	}

	schema.Types = Types
	return schema, nil
}

const tableSQL = `
WITH cols AS (
	SELECT *
	FROM information_schema.columns
	WHERE table_schema = $1
	ORDER BY table_name, ordinal_position
), ref AS (
	SELECT
		tc1.table_schema,
		tc1.table_name,
		tc1.constraint_type,
		rc.constraint_name,
		tc2.table_name AS target_table_name,
		rc.unique_constraint_name AS target_constraint_name
	FROM information_schema.referential_constraints rc
	JOIN information_schema.table_constraints tc1 ON tc1.constraint_name = rc.constraint_name
	JOIN information_schema.table_constraints tc2 ON tc2.constraint_name = rc.unique_constraint_name
	WHERE rc.constraint_schema = $1
		AND rc.unique_constraint_schema = $1
)
SELECT json_build_object(
	'schema_name', $1::TEXT,
	'tables', (
		SELECT json_agg(json_build_object(
			'table_name', table_name,
			'insertable', is_insertable_into = 'YES'
		))
		FROM information_schema.tables
		WHERE table_schema = $1
	),
	'columns', (
		SELECT json_agg(json_build_object(
			'table_name', c.table_name,
			'column_name', c.column_name,
			'default', c.column_default,
			'is_nullable', c.is_nullable = 'YES',
			'is_updatable', c.is_updatable = 'YES',
			'data_type', CASE WHEN c.data_type = 'USER-DEFINED' THEN c.udt_name ELSE c.data_type END
		))
		FROM cols c
		WHERE table_schema = $1
	),
	'keys', (
		SELECT json_agg(json_build_object(
			'table_name', k.table_name,
			'constraint_name', k.constraint_name,
			'type', constraint_type,
			'is_primary_key', k.constraint_type = 'PRIMARY KEY',
			'column_names', (
				SELECT json_agg(kc.column_name)
				FROM information_schema.constraint_column_usage kc
				WHERE kc.constraint_name = k.constraint_name
			)
		))
		FROM information_schema.table_constraints k
		WHERE table_schema = $1 AND (
			constraint_type = 'PRIMARY KEY' OR constraint_type = 'UNIQUE'
		)
	),
	'foreign_keys', (
		SELECT json_agg(json_build_object(
			'table_name', ref.table_name,
			'constraint_name', ref.constraint_name,
			'type', ref.constraint_type,
			'column_names', (
				SELECT json_agg(ktc.column_name)
				FROM information_schema.key_column_usage ktc
				WHERE ktc.constraint_name = ref.constraint_name
			),
			'target_table_name', ref.target_table_name,
			'target_constraint_name', ref.target_constraint_name
		))
		FROM ref
	)
)
`
