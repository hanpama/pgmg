package dsl_test

import (
	"fmt"
	"testing"

	"github.com/hanpama/pgmg/dsl"
)

var modelTestCases = testCases{
	{
		name: "Baseline",
		test: func() string {
			return fmt.Sprintf(`SELECT "u"."id", "p"."created" ` +
				`FROM "user" AS "u" INNER JOIN "user_profile" AS "p" ON "u"."id" = "p"."user_id" ` +
				`WHERE "u"."id" = $1 ORDER BY "u"."created" LIMIT 10 OFFSET 10`,
			)
		},
		expecting: `SELECT "u"."id", "p"."created" ` +
			`FROM "user" AS "u" INNER JOIN "user_profile" AS "p" ON "u"."id" = "p"."user_id" ` +
			`WHERE "u"."id" = $1 ORDER BY "u"."created" LIMIT 10 OFFSET 10`,
	},
	{
		name: "SelectWithJoin",
		test: func() string {
			user := User.As("u")
			profile := UserProfile.As("p")

			query := dsl.Select(user.ID, profile.Created).
				From(user.InnerJoin(profile).Onf(`%s = %s`, user.ID, profile.UserID)).
				Wheref("%s = $1", user.ID).
				OrderBy(user.Created).Limit(10).Offset(10)

			return query.Statement()
		},
		expecting: `` +
			`SELECT "u"."id", "p"."created" ` +
			`FROM "public"."user" AS "u" INNER JOIN "public"."user_profile" AS "p" ON "u"."id" = "p"."user_id" ` +
			`WHERE "u"."id" = $1 ORDER BY "u"."created" LIMIT 10 OFFSET 10`,
	},
	{
		name: "InsertValues",
		test: func() string {
			user := User
			query := dsl.InsertInto(user, user.ID, user.Created).Valuesf("($1, $2)")
			return query.Statement()
		},
		expecting: `` +
			`INSERT INTO "public"."user" ("id", "created") VALUES ($1, $2)`,
	},
	{
		name: "InsertQuery",
		test: func() string {
			user := User
			query := dsl.InsertInto(user, user.ID, user.Created).Select(
				dsl.Select().Fromf(`json_populate_recordset(null::%s, $1)`, user),
			)
			return query.Statement()
		},
		expecting: `` +
			`INSERT INTO "public"."user" ("id", "created") ` +
			`SELECT * FROM json_populate_recordset(null::"public"."user", $1)`,
	},
	{
		name: "InsertQueryReturning",
		test: func() string {
			user := User
			query := dsl.InsertInto(user, user.ID, user.Created).Select(
				dsl.Select().Fromf(`json_populate_recordset(null::%s, $1)`, user),
			).Returning()
			return query.Statement()
		},
		expecting: `` +
			`INSERT INTO "public"."user" ("id", "created") ` +
			`SELECT * FROM json_populate_recordset(null::"public"."user", $1) ` +
			`RETURNING *`,
	},
	{
		name: "SelectWhereSubquery",
		test: func() string {
			user := User
			userProfile := UserProfile

			query := dsl.Select().From(user).Where(
				dsl.Select("count(*) > 0").From(userProfile).Where(user.ID.Eq("$1")),
			)
			return query.Statement()
		},
		expecting: `` +
			`SELECT * FROM "public"."user" WHERE (SELECT count(*) > 0 ` +
			`FROM "public"."user_profile" WHERE "public"."user"."id" = $1)`,
	},
	{
		name: "SelectSubquery",
		test: func() string {
			user := User
			userProfile := UserProfile
			query := dsl.Select(
				dsl.Select("count(*) > 0").From(userProfile).Where(user.ID.Eq("$1")),
			).From(user)
			return query.Statement()
		},
		expecting: `SELECT (SELECT count(*) > 0 FROM "public"."user_profile" WHERE "public"."user"."id" = $1) FROM "public"."user"`,
	},
	{
		name: "SelectSubqueryAlias",
		test: func() string {
			user := User
			userProfile := UserProfile
			query := dsl.Select(
				dsl.Select("count(*) > 0").From(userProfile).Where(user.ID.Eq("$1")).As("foobar"),
			).
				From(user).
				Wheref("foobar > $1")
			return query.Statement()
		},
		expecting: `SELECT (SELECT count(*) > 0 FROM "public"."user_profile" WHERE "public"."user"."id" = $1) AS "foobar" FROM "public"."user" ` +
			`WHERE foobar > $1`,
	},
	{
		name: "Update",
		test: func() string {
			query := dsl.
				Update(User).
				Set(
					User.Deposit.Set(User.Deposit.Mul(2)),
					User.Created.Set("$2"),
				).
				Where(User.ID.Eq("$1"))

			return query.Statement()
		},
		expecting: `` +
			`UPDATE "public"."user" SET "deposit" = "public"."user"."deposit" * 2, "created" = $2 ` +
			`WHERE "public"."user"."id" = $1`,
	},
}

func TestDSLWithModel(t *testing.T) {
	for _, tc := range modelTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.test() != tc.expecting {
				t.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}

func BenchmarkDSLWithModel(b *testing.B) {
	for _, tc := range modelTestCases {
		b.Run(tc.name, func(b *testing.B) {
			if tc.test() != tc.expecting {
				b.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}
