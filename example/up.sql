DO $$ BEGIN

CREATE SCHEMA wise;

CREATE TABLE wise.semester (
  id SERIAL PRIMARY KEY,
  year INTEGER NOT NULL,
  season TEXT NOT NULL,

  UNIQUE (year, season)
);

CREATE TABLE wise.product (
  id SERIAL PRIMARY KEY,
  price NUMERIC NOT NULL CHECK(price > 0),
  stocked TIMESTAMPTZ NOT NULL,
  sold TIMESTAMPTZ
);

END$$;