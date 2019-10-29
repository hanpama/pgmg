package dsl_test

import "github.com/hanpama/pgmg/dsl"

type testCases []struct {
	name      string
	test      func() string
	expecting string
}

type UserTable struct {
	dsl.TableReference
	ID      dsl.ColumnReference
	Created dsl.ColumnReference
	Deposit dsl.ColumnReference
}

func newUserTable(alias string) *UserTable {
	table := dsl.TableReference{TableSchema: "public", TableName: `user`, Alias: alias}
	return &UserTable{
		table,
		dsl.ColumnReference{TableReference: table, ColumnName: "id"},
		dsl.ColumnReference{TableReference: table, ColumnName: "created"},
		dsl.ColumnReference{TableReference: table, ColumnName: "deposit"},
	}
}

var User = newUserTable("")

func (t *UserTable) As(alias string) *UserTable { return newUserTable(alias) }

type UserProfileTable struct {
	dsl.TableReference
	UserID  dsl.ColumnReference
	Created dsl.ColumnReference
}

func newUserProfileTable(alias string) *UserProfileTable {
	table := dsl.TableReference{TableSchema: "public", TableName: `user_profile`, Alias: alias}
	return &UserProfileTable{
		table,
		dsl.ColumnReference{TableReference: table, ColumnName: "user_id"},
		dsl.ColumnReference{TableReference: table, ColumnName: "created"},
	}
}

var UserProfile = newUserProfileTable("")

func (t *UserProfileTable) As(alias string) *UserProfileTable { return newUserProfileTable(alias) }
