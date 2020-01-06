// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package example

import (
	"context"
	"encoding/json"
	"time"
)

// Database represents PostgresQL database
type Database interface {
	QueryScan(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

// SemesterRow represents a row for table "semester"
type SemesterRow struct {
	ID     *int32 `json:"id"`
	Year   int32  `json:"year"`
	Season string `json:"season"`
}

// SemesterRowID represents column "id" of table "semester"
type SemesterRowID *int32

// SemesterRowYear represents column "year" of table "semester"
type SemesterRowYear int32

// SemesterRowSeason represents column "season" of table "semester"
type SemesterRowSeason string

// NewSemesterRow creates a new row for table "semester" with all column values
func NewSemesterRow(
	id SemesterRowID,
	year SemesterRowYear,
	season SemesterRowSeason,
) *SemesterRow {
	return &SemesterRow{
		(*int32)(id),
		(int32)(year),
		(string)(season),
	}
}

func (r *SemesterRow) receive() []interface{} {
	return []interface{}{&r.ID, &r.Year, &r.Season}
}

// LectureRow represents a row for table "lecture"
type LectureRow struct {
	ID         *int32 `json:"id"`
	Title      string `json:"title"`
	SemesterID int32  `json:"semester_id"`
	CourseID   int32  `json:"course_id"`
	TutorID    int32  `json:"tutor_id"`
}

// LectureRowID represents column "id" of table "lecture"
type LectureRowID *int32

// LectureRowTitle represents column "title" of table "lecture"
type LectureRowTitle string

// LectureRowSemesterID represents column "semester_id" of table "lecture"
type LectureRowSemesterID int32

// LectureRowCourseID represents column "course_id" of table "lecture"
type LectureRowCourseID int32

// LectureRowTutorID represents column "tutor_id" of table "lecture"
type LectureRowTutorID int32

// NewLectureRow creates a new row for table "lecture" with all column values
func NewLectureRow(
	id LectureRowID,
	title LectureRowTitle,
	semesterID LectureRowSemesterID,
	courseID LectureRowCourseID,
	tutorID LectureRowTutorID,
) *LectureRow {
	return &LectureRow{
		(*int32)(id),
		(string)(title),
		(int32)(semesterID),
		(int32)(courseID),
		(int32)(tutorID),
	}
}

func (r *LectureRow) receive() []interface{} {
	return []interface{}{&r.ID, &r.Title, &r.SemesterID, &r.CourseID, &r.TutorID}
}

// CourseRow represents a row for table "course"
type CourseRow struct {
	ID      *int32 `json:"id"`
	Title   string `json:"title"`
	Credits int32  `json:"credits"`
}

// CourseRowID represents column "id" of table "course"
type CourseRowID *int32

// CourseRowTitle represents column "title" of table "course"
type CourseRowTitle string

// CourseRowCredits represents column "credits" of table "course"
type CourseRowCredits int32

// NewCourseRow creates a new row for table "course" with all column values
func NewCourseRow(
	id CourseRowID,
	title CourseRowTitle,
	credits CourseRowCredits,
) *CourseRow {
	return &CourseRow{
		(*int32)(id),
		(string)(title),
		(int32)(credits),
	}
}

func (r *CourseRow) receive() []interface{} {
	return []interface{}{&r.ID, &r.Title, &r.Credits}
}

// ProfessorRow represents a row for table "professor"
type ProfessorRow struct {
	ID         *int32     `json:"id"`
	FamilyName string     `json:"family_name"`
	GivenName  string     `json:"given_name"`
	BirthDate  time.Time  `json:"birth_date"`
	HiredDate  *time.Time `json:"hired_date"`
}

// ProfessorRowID represents column "id" of table "professor"
type ProfessorRowID *int32

// ProfessorRowFamilyName represents column "family_name" of table "professor"
type ProfessorRowFamilyName string

// ProfessorRowGivenName represents column "given_name" of table "professor"
type ProfessorRowGivenName string

// ProfessorRowBirthDate represents column "birth_date" of table "professor"
type ProfessorRowBirthDate time.Time

// ProfessorRowHiredDate represents column "hired_date" of table "professor"
type ProfessorRowHiredDate *time.Time

// NewProfessorRow creates a new row for table "professor" with all column values
func NewProfessorRow(
	id ProfessorRowID,
	familyName ProfessorRowFamilyName,
	givenName ProfessorRowGivenName,
	birthDate ProfessorRowBirthDate,
	hiredDate ProfessorRowHiredDate,
) *ProfessorRow {
	return &ProfessorRow{
		(*int32)(id),
		(string)(familyName),
		(string)(givenName),
		(time.Time)(birthDate),
		(*time.Time)(hiredDate),
	}
}

func (r *ProfessorRow) receive() []interface{} {
	return []interface{}{&r.ID, &r.FamilyName, &r.GivenName, &r.BirthDate, &r.HiredDate}
}

// SemesterPkey represents key defined by UNIQUE constraint "semester_pkey" for table "semester"
type SemesterPkey struct {
	ID int32 `json:"id"`
}

func (r *SemesterRow) SemesterPkey() SemesterPkey {
	k := SemesterPkey{}
	if r.ID != nil {
		k.ID = *r.ID
	}
	return k
}

// GetBySemesterPkey gets matching rows for given SemesterPkey keys from table "semester"
func GetBySemesterPkey(ctx context.Context, db Database, keys ...SemesterPkey) (rows []*SemesterRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*SemesterRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &SemesterRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				id
			FROM json_populate_recordset(null::"wise"."semester", $1)
		)
		SELECT id, year, season
		FROM __key JOIN "wise"."semester" AS __table USING (id)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].SemesterPkey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveBySemesterPkey upserts the given rows for table "semester" checking uniqueness by contstraint "semester_pkey"
func SaveBySemesterPkey(ctx context.Context, db Database, rows ...*SemesterRow) ([]*SemesterRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.semester_id_seq'::regclass)) id,
			  __input.year,
			  __input.season
			FROM json_populate_recordset(null::"wise"."semester", $1) __input
		)
		INSERT INTO "wise"."semester" SELECT * FROM __values
		ON CONFLICT (id) DO UPDATE
			SET (id, year, season) = (
				SELECT id, year, season FROM __values
			)
		RETURNING id, year, season
	`, string(b))
	return rows, err
}

// DeleteBySemesterPkey deletes matching rows by SemesterPkey keys from table "semester"
func DeleteBySemesterPkey(ctx context.Context, db Database, keys ...SemesterPkey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."semester", $1))
		DELETE FROM "wise"."semester" AS __table
			USING __key
			WHERE (__key.id = __table.id)
			`, string(b))
}

// SemesterYearSeasonKey represents key defined by UNIQUE constraint "semester_year_season_key" for table "semester"
type SemesterYearSeasonKey struct {
	Year   int32  `json:"year"`
	Season string `json:"season"`
}

func (r *SemesterRow) SemesterYearSeasonKey() SemesterYearSeasonKey {
	k := SemesterYearSeasonKey{}
	k.Year = r.Year

	k.Season = r.Season

	return k
}

// GetBySemesterYearSeasonKey gets matching rows for given SemesterYearSeasonKey keys from table "semester"
func GetBySemesterYearSeasonKey(ctx context.Context, db Database, keys ...SemesterYearSeasonKey) (rows []*SemesterRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*SemesterRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &SemesterRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				year, season
			FROM json_populate_recordset(null::"wise"."semester", $1)
		)
		SELECT id, year, season
		FROM __key JOIN "wise"."semester" AS __table USING (year, season)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].SemesterYearSeasonKey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveBySemesterYearSeasonKey upserts the given rows for table "semester" checking uniqueness by contstraint "semester_year_season_key"
func SaveBySemesterYearSeasonKey(ctx context.Context, db Database, rows ...*SemesterRow) ([]*SemesterRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.semester_id_seq'::regclass)) id,
			  __input.year,
			  __input.season
			FROM json_populate_recordset(null::"wise"."semester", $1) __input
		)
		INSERT INTO "wise"."semester" SELECT * FROM __values
		ON CONFLICT (year, season) DO UPDATE
			SET (id, year, season) = (
				SELECT id, year, season FROM __values
			)
		RETURNING id, year, season
	`, string(b))
	return rows, err
}

// DeleteBySemesterYearSeasonKey deletes matching rows by SemesterYearSeasonKey keys from table "semester"
func DeleteBySemesterYearSeasonKey(ctx context.Context, db Database, keys ...SemesterYearSeasonKey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."semester", $1))
		DELETE FROM "wise"."semester" AS __table
			USING __key
			WHERE (__key.year = __table.year)
			  AND (__key.season = __table.season)
			`, string(b))
}

// LecturePkey represents key defined by UNIQUE constraint "lecture_pkey" for table "lecture"
type LecturePkey struct {
	ID int32 `json:"id"`
}

func (r *LectureRow) LecturePkey() LecturePkey {
	k := LecturePkey{}
	if r.ID != nil {
		k.ID = *r.ID
	}
	return k
}

// GetByLecturePkey gets matching rows for given LecturePkey keys from table "lecture"
func GetByLecturePkey(ctx context.Context, db Database, keys ...LecturePkey) (rows []*LectureRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*LectureRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &LectureRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				id
			FROM json_populate_recordset(null::"wise"."lecture", $1)
		)
		SELECT id, title, semester_id, course_id, tutor_id
		FROM __key JOIN "wise"."lecture" AS __table USING (id)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].LecturePkey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveByLecturePkey upserts the given rows for table "lecture" checking uniqueness by contstraint "lecture_pkey"
func SaveByLecturePkey(ctx context.Context, db Database, rows ...*LectureRow) ([]*LectureRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.lecture_id_seq'::regclass)) id,
			  __input.title,
			  __input.semester_id,
			  __input.course_id,
			  __input.tutor_id
			FROM json_populate_recordset(null::"wise"."lecture", $1) __input
		)
		INSERT INTO "wise"."lecture" SELECT * FROM __values
		ON CONFLICT (id) DO UPDATE
			SET (id, title, semester_id, course_id, tutor_id) = (
				SELECT id, title, semester_id, course_id, tutor_id FROM __values
			)
		RETURNING id, title, semester_id, course_id, tutor_id
	`, string(b))
	return rows, err
}

// DeleteByLecturePkey deletes matching rows by LecturePkey keys from table "lecture"
func DeleteByLecturePkey(ctx context.Context, db Database, keys ...LecturePkey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."lecture", $1))
		DELETE FROM "wise"."lecture" AS __table
			USING __key
			WHERE (__key.id = __table.id)
			`, string(b))
}

// LectureSemesterIDCourseIDTutorIDKey represents key defined by UNIQUE constraint "lecture_semester_id_course_id_tutor_id_key" for table "lecture"
type LectureSemesterIDCourseIDTutorIDKey struct {
	SemesterID int32 `json:"semester_id"`
	CourseID   int32 `json:"course_id"`
	TutorID    int32 `json:"tutor_id"`
}

func (r *LectureRow) LectureSemesterIDCourseIDTutorIDKey() LectureSemesterIDCourseIDTutorIDKey {
	k := LectureSemesterIDCourseIDTutorIDKey{}
	k.SemesterID = r.SemesterID

	k.CourseID = r.CourseID

	k.TutorID = r.TutorID

	return k
}

// GetByLectureSemesterIDCourseIDTutorIDKey gets matching rows for given LectureSemesterIDCourseIDTutorIDKey keys from table "lecture"
func GetByLectureSemesterIDCourseIDTutorIDKey(ctx context.Context, db Database, keys ...LectureSemesterIDCourseIDTutorIDKey) (rows []*LectureRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*LectureRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &LectureRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				semester_id, course_id, tutor_id
			FROM json_populate_recordset(null::"wise"."lecture", $1)
		)
		SELECT id, title, semester_id, course_id, tutor_id
		FROM __key JOIN "wise"."lecture" AS __table USING (semester_id, course_id, tutor_id)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].LectureSemesterIDCourseIDTutorIDKey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveByLectureSemesterIDCourseIDTutorIDKey upserts the given rows for table "lecture" checking uniqueness by contstraint "lecture_semester_id_course_id_tutor_id_key"
func SaveByLectureSemesterIDCourseIDTutorIDKey(ctx context.Context, db Database, rows ...*LectureRow) ([]*LectureRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.lecture_id_seq'::regclass)) id,
			  __input.title,
			  __input.semester_id,
			  __input.course_id,
			  __input.tutor_id
			FROM json_populate_recordset(null::"wise"."lecture", $1) __input
		)
		INSERT INTO "wise"."lecture" SELECT * FROM __values
		ON CONFLICT (semester_id, course_id, tutor_id) DO UPDATE
			SET (id, title, semester_id, course_id, tutor_id) = (
				SELECT id, title, semester_id, course_id, tutor_id FROM __values
			)
		RETURNING id, title, semester_id, course_id, tutor_id
	`, string(b))
	return rows, err
}

// DeleteByLectureSemesterIDCourseIDTutorIDKey deletes matching rows by LectureSemesterIDCourseIDTutorIDKey keys from table "lecture"
func DeleteByLectureSemesterIDCourseIDTutorIDKey(ctx context.Context, db Database, keys ...LectureSemesterIDCourseIDTutorIDKey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."lecture", $1))
		DELETE FROM "wise"."lecture" AS __table
			USING __key
			WHERE (__key.semester_id = __table.semester_id)
			  AND (__key.course_id = __table.course_id)
			  AND (__key.tutor_id = __table.tutor_id)
			`, string(b))
}

// CoursePkey represents key defined by UNIQUE constraint "course_pkey" for table "course"
type CoursePkey struct {
	ID int32 `json:"id"`
}

func (r *CourseRow) CoursePkey() CoursePkey {
	k := CoursePkey{}
	if r.ID != nil {
		k.ID = *r.ID
	}
	return k
}

// GetByCoursePkey gets matching rows for given CoursePkey keys from table "course"
func GetByCoursePkey(ctx context.Context, db Database, keys ...CoursePkey) (rows []*CourseRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*CourseRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &CourseRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				id
			FROM json_populate_recordset(null::"wise"."course", $1)
		)
		SELECT id, title, credits
		FROM __key JOIN "wise"."course" AS __table USING (id)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].CoursePkey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveByCoursePkey upserts the given rows for table "course" checking uniqueness by contstraint "course_pkey"
func SaveByCoursePkey(ctx context.Context, db Database, rows ...*CourseRow) ([]*CourseRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.course_id_seq'::regclass)) id,
			  __input.title,
			  __input.credits
			FROM json_populate_recordset(null::"wise"."course", $1) __input
		)
		INSERT INTO "wise"."course" SELECT * FROM __values
		ON CONFLICT (id) DO UPDATE
			SET (id, title, credits) = (
				SELECT id, title, credits FROM __values
			)
		RETURNING id, title, credits
	`, string(b))
	return rows, err
}

// DeleteByCoursePkey deletes matching rows by CoursePkey keys from table "course"
func DeleteByCoursePkey(ctx context.Context, db Database, keys ...CoursePkey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."course", $1))
		DELETE FROM "wise"."course" AS __table
			USING __key
			WHERE (__key.id = __table.id)
			`, string(b))
}

// ProfessorPkey represents key defined by UNIQUE constraint "professor_pkey" for table "professor"
type ProfessorPkey struct {
	ID int32 `json:"id"`
}

func (r *ProfessorRow) ProfessorPkey() ProfessorPkey {
	k := ProfessorPkey{}
	if r.ID != nil {
		k.ID = *r.ID
	}
	return k
}

// GetByProfessorPkey gets matching rows for given ProfessorPkey keys from table "professor"
func GetByProfessorPkey(ctx context.Context, db Database, keys ...ProfessorPkey) (rows []*ProfessorRow, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*ProfessorRow, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &ProfessorRow{}
		return rows[i].receive()
	}, `
		WITH __key AS (
			SELECT ROW_NUMBER() over () __keyindex,
				id
			FROM json_populate_recordset(null::"wise"."professor", $1)
		)
		SELECT id, family_name, given_name, birth_date, hired_date
		FROM __key JOIN "wise"."professor" AS __table USING (id)
		ORDER BY __keyindex
  `, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].ProfessorPkey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

// SaveByProfessorPkey upserts the given rows for table "professor" checking uniqueness by contstraint "professor_pkey"
func SaveByProfessorPkey(ctx context.Context, db Database, rows ...*ProfessorRow) ([]*ProfessorRow, error) {
	b, err := json.Marshal(rows)
	if err != nil {
		return rows, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} { return rows[i].receive() }, `
		WITH __values AS (
			SELECT
				COALESCE(__input.id, nextval('wise.professor_id_seq'::regclass)) id,
			  __input.family_name,
			  __input.given_name,
			  __input.birth_date,
			  __input.hired_date
			FROM json_populate_recordset(null::"wise"."professor", $1) __input
		)
		INSERT INTO "wise"."professor" SELECT * FROM __values
		ON CONFLICT (id) DO UPDATE
			SET (id, family_name, given_name, birth_date, hired_date) = (
				SELECT id, family_name, given_name, birth_date, hired_date FROM __values
			)
		RETURNING id, family_name, given_name, birth_date, hired_date
	`, string(b))
	return rows, err
}

// DeleteByProfessorPkey deletes matching rows by ProfessorPkey keys from table "professor"
func DeleteByProfessorPkey(ctx context.Context, db Database, keys ...ProfessorPkey) (int64, error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, `
		WITH __key AS (SELECT * FROM json_populate_recordset(null::"wise"."professor", $1))
		DELETE FROM "wise"."professor" AS __table
			USING __key
			WHERE (__key.id = __table.id)
			`, string(b))
}
