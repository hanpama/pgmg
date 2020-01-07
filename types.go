package main

import (
	"fmt"
	"log"
)

type typeMapping struct {
	PGType       string `json:"pgType"`
	Name         string `json:"name"`
	NullableName string `json:"nullableName"`
	Module       string `json:"module"`
}

var types = func() map[string]typeMapping {
	res := make(map[string]typeMapping)
	for _, item := range []typeMapping{
		{PGType: "boolean", Name: "bool", NullableName: "*bool"},
		{PGType: "character", Name: "string", NullableName: "*string"},
		{PGType: "character varying", Name: "string", NullableName: "*string"},
		{PGType: "text", Name: "string", NullableName: "*string"},
		{PGType: "money", Name: "string", NullableName: "*string"},
		{PGType: "inet", Name: "string", NullableName: "*string"},
		{PGType: "uuid", Name: "string", NullableName: "*string"},
		{PGType: "jsonb", Name: "string", NullableName: "*string"},
		{PGType: "json", Name: "string", NullableName: "*string"},
		{PGType: "smallint", Name: "int16", NullableName: "*int16"},
		{PGType: "integer", Name: "int32", NullableName: "*int32"},
		{PGType: "bigint", Name: "int64", NullableName: "*int64"},
		{PGType: "smallserial", Name: "uint16", NullableName: "*uint16"},
		{PGType: "serial", Name: "uint32", NullableName: "*uint32"},
		{PGType: "bigserial", Name: "uint64", NullableName: "*uint64"},
		{PGType: "real", Name: "float32", NullableName: "*float32"},
		{PGType: "numeric", Name: "float64", NullableName: "*float64"},
		{PGType: "double precision", Name: "float64", NullableName: "*float64"},
		{PGType: "bytea", Name: "[]byte", NullableName: "[]byte"},
		{PGType: "date", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
		{PGType: "timestamp with time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
		{PGType: "time with time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
		{PGType: "time without time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
		{PGType: "timestamp without time zone", Name: "time.Time", NullableName: "*time.Time", Module: "time"},
		{PGType: "interval", Name: "time.Duration", NullableName: "*time.Duration", Module: "time"},
		{PGType: "char", Name: "uint8", NullableName: "*uint8"},
		{PGType: "bit", Name: "uint8", NullableName: "*uint8"},
		{PGType: "any", Name: "[]byte", NullableName: "[]byte"},
		{PGType: "bit varying", Name: "[]byte", NullableName: "[]byte"},
		{PGType: "hstore", Name: "[]byte", NullableName: "[]byte"},
		{PGType: "geometry", Name: "string", NullableName: "*string"},  // EWKBHex
		{PGType: "geography", Name: "string", NullableName: "*string"}, // EWKBHex
	} {
		res[item.PGType] = item
	}

	return res
}()

func pgTypeToGoType(pgType string) typeMapping {
	l, ok := types[pgType]
	if ok {
		return l
	}
	log.Print(fmt.Errorf("Not supported postgres type: %s", pgType))
	return typeMapping{
		PGType:       pgType,
		Name:         "[]byte",
		NullableName: "[]byte",
	}
}
