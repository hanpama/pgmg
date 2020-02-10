DO $$ BEGIN

CREATE SCHEMA wise;

CREATE TABLE wise.pop (
  name TEXT,
  year INT,
  description TEXT NOT NULL,

  PRIMARY KEY (name, year)
);

CREATE TABLE wise.product (
  id SERIAL PRIMARY KEY,
  price MONEY NOT NULL,
  name VARCHAR(120) NOT NULL,
  alias VARCHAR(32) NOT NULL,
  stocked TIMESTAMPTZ NOT NULL,
  sold TIMESTAMPTZ,

  UNIQUE(alias)
);

CREATE TABLE wise.package (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  available BOOLEAN DEFAULT true
);

CREATE TABLE wise.package_product (
  package_id UUID REFERENCES wise.package ON DELETE CASCADE NOT NULL,
  product_id INTEGER REFERENCES wise.product NOT NULL,

  UNIQUE(package_id, product_id)
);

CREATE TABLE wise.campaign (
  id UUID PRIMARY KEY,
  pop_name TEXT,
  pop_year INT,

  FOREIGN KEY (pop_name, pop_year) REFERENCES wise.pop (name, year)
);

CREATE VIEW wise.package_agg AS
  SELECT id, name, available, (SELECT count(*) FROM wise.package_product WHERE package_id = t1.id)
  FROM wise.package t1;

END$$;