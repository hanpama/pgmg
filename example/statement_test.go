package example_test

import (
	"testing"

	"github.com/hanpama/pgmg/example/wise"
)

var statementTestCases = testCases{
	{
		name: "Select",
		test: func() string {
			q := wise.Lecture.Select()
			return q.Statement()
		},
		expecting: `` +
			`SELECT "wise"."lecture"."id", "wise"."lecture"."title", "wise"."lecture"."semester_id", ` +
			`"wise"."lecture"."course_id", "wise"."lecture"."tutor_id" ` +
			`FROM "wise"."lecture"`,
	},
	{
		name: "InsertManyJSON",
		test: func() string {
			q := wise.Lecture.InsertManyJSON("$1")
			return q.Statement()
		},
		expecting: `` +
			`INSERT INTO "wise"."lecture" ` +
			`("id", "title", "semester_id", "course_id", "tutor_id") ` +
			`SELECT ` +
			`COALESCE(__iv__."id", nextval('wise.lecture_id_seq'::regclass)), ` +
			`__iv__."title", __iv__."semester_id", __iv__."course_id", __iv__."tutor_id" ` +
			`FROM json_populate_recordset(null::"wise"."lecture", $1) AS __iv__`,
	},
	{
		name: "UpdateJSON",
		test: func() string {
			q := wise.Lecture.UpdateJSON("$1")
			return q.Statement()
		},
		expecting: `` +
			`UPDATE "wise"."lecture" ` +
			`SET "id" = COALESCE(__ch__."id", "wise"."lecture"."id"), ` +
			`"title" = COALESCE(__ch__."title", "wise"."lecture"."title"), ` +
			`"semester_id" = COALESCE(__ch__."semester_id", "wise"."lecture"."semester_id"), ` +
			`"course_id" = COALESCE(__ch__."course_id", "wise"."lecture"."course_id"), ` +
			`"tutor_id" = COALESCE(__ch__."tutor_id", "wise"."lecture"."tutor_id") ` +
			`FROM json_populate_record(null::"wise"."lecture", $1) AS __ch__`,
	},
	{
		name: "Delete",
		test: func() string {
			q := wise.Lecture.DeleteWhere(wise.Lecture.ID.Eq("$1"))
			return q.Statement()
		},
		expecting: `DELETE FROM "wise"."lecture" WHERE "wise"."lecture"."id" = $1`,
	},
}

func TestGeneratedQueryBuilder(t *testing.T) {
	for _, tc := range statementTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.test() != tc.expecting {
				t.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}

func BenchmarkGeneratedQueryBuilder(b *testing.B) {
	for _, tc := range statementTestCases {
		b.Run(tc.name, func(b *testing.B) {
			if tc.test() != tc.expecting {
				b.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}
