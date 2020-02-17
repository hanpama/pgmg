package tables

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type Keys []interface{}

// Unique returns new Keys each key of which is unique
func (ks Keys) Unique() (uks Keys) {
	seen := make(map[interface{}]struct{}, len(ks))
	uks = make(Keys, 0, len(uks))
	for _, k := range ks {
		if _, ok := seen[k]; !ok && k != nil {
			seen[k] = struct{}{}
			uks = append(uks, k)
		}
	}
	return uks
}

// MarshalJSON implements Marshaler
func (ks Keys) MarshalJSON() ([]byte, error) {
	var buff bytes.Buffer
	var err error
	buff.WriteRune('[')
	enc := json.NewEncoder(&buff)
	for i, k := range ks {
		if i > 0 {
			buff.WriteRune(',')
		}
		if err = enc.Encode(k); err != nil {
			return nil, err
		}
	}
	buff.WriteRune(']')
	return buff.Bytes(), nil
}

func execWithJSONArgs(ctx context.Context, db SQLHandle, sql string, args ...interface{}) (numRows int64, err error) {
	bArgs := make([]interface{}, len(args))
	for i, arg := range args {
		if bArgs[i], err = json.Marshal(arg); err != nil {
			return 0, err
		}
	}
	return db.ExecAndCount(ctx, sql, bArgs...)
}

func queryWithJSONArgs(ctx context.Context, db SQLHandle, receive func(int) []interface{}, sql string, args ...interface{}) (numRows int, err error) {
	bArgs := make([]interface{}, len(args))
	for i, arg := range args {
		if bArgs[i], err = json.Marshal(arg); err != nil {
			return 0, err
		}
	}
	return db.QueryAndReceive(ctx, receive, sql, bArgs...)
}

func formatError(methodName string, err error) error {
	return fmt.Errorf("%w(%s, %s)", ErrPGMG, methodName, err.Error())
}

var ErrPGMG = fmt.Errorf("errPGMG")

type SQLHandle interface {
	QueryAndReceive(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	ExecAndCount(ctx context.Context, sql string, args ...interface{}) (int64, error)
}
