package tables

import (
	"context"
	"encoding/json"
	"fmt"
)

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
