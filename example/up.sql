CREATE SCHEMA wise;

CREATE TABLE wise.semester (
  id SERIAL PRIMARY KEY,
  year INTEGER NOT NULL,
  season TEXT NOT NULL,

  UNIQUE (year, season)
);

CREATE TABLE wise.course (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  credits INTEGER NOT NULL
);

CREATE TABLE wise.professor (
  id SERIAL PRIMARY KEY,
  family_name TEXT NOT NULL,
  given_name TEXT NOT NULL,
  birth_date TIMESTAMPTZ NOT NULL,
  hired_date TIMESTAMPTZ
);

CREATE TABLE wise.lecture (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  semester_id INTEGER REFERENCES wise.semester NOT NULL,
  course_id INTEGER REFERENCES wise.course NOT NULL,
  tutor_id INTEGER REFERENCES wise.professor NOT NULL,

  UNIQUE (semester_id, course_id, tutor_id)
);
