CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  "id"          UUID            PRIMARY KEY DEFAULT uuid_generate_v4(),
  "email"       VARCHAR         UNIQUE NOT NULL,
  "password"    VARCHAR         NOT NULL
);
