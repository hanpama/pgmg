package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

// IntrospectQuery inspects query by preparing it and try to create a temporary view
func IntrospectQuery(tx *sql.Tx, name string, query string) (qi Query, err error) {
	qi.Name = name
	qi.SQL = query

	// Prepare and inspect
	qi.ParamTypes, err = prepareAndInspect(tx, name, query)
	if err != nil {
		return qi, err
	}

	// Create temporary view and inspect
	for i := range qi.ParamTypes {
		query = strings.ReplaceAll(
			query, fmt.Sprintf("$%d", len(qi.ParamTypes)-i), "NULL",
		)
	}
	qi.Returning, _ = createTemporaryViewAndInspect(tx, name, query) // ignoring error

	return qi, nil
}

func prepareAndInspect(tx *sql.Tx, name, query string) (paramTypes []string, err error) {
	_, err = tx.Exec(`PREPARE "` + name + `" AS ` + query)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(`
		SELECT UNNEST(parameter_types) FROM pg_catalog.pg_prepared_statements WHERE name = $1
	`, name)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dat string
		err = rows.Scan(&dat)
		if err != nil {
			return nil, err
		}
		paramTypes = append(paramTypes, dat)
	}
	return paramTypes, nil
}

func createTemporaryViewAndInspect(tx *sql.Tx, name, query string) ([]Column, error) {
	var columns []Column

	_, err := tx.Exec(`CREATE TEMPORARY VIEW "` + name + `" AS ` + query)
	if err != nil {
		return nil, err
	}

	var b []byte
	err = tx.QueryRow(`
		SELECT json_agg(json_build_object(
			'name', c.column_name,
			'is_nullable', c.is_nullable = 'YES',
			'default', c.column_default,
			'data_type', CASE WHEN c.data_type = 'USER-DEFINED' THEN c.udt_name ELSE c.data_type END
		)) FROM information_schema.columns c WHERE table_name = $1`, name).Scan(&b)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &columns); err != nil {
		return nil, err
	}
	return columns, nil
}
