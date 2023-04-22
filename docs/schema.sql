-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-04-21T14:11:24.025Z

CREATE TYPE "user_type" AS ENUM (
  'ADMIN',
  'CLIENT'
);

CREATE TYPE "currency" AS ENUM (
  'UGX'
);

CREATE TYPE "log_type" AS ENUM (
  'REQUEST',
  'FAILED_VALIDATION',
  'POSITIVE_REQUEST',
  'NEGATIVE_REQUEST'
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "type" user_type NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "account_no" varchar(10) UNIQUE NOT NULL,
  "balance" bigint NOT NULL,
  "currency" currency NOT NULL,
  "password_changed_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "loans" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "type" log_type NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "users" ("account_no");

CREATE INDEX ON "loans" ("username");

CREATE INDEX ON "logs" ("username");

COMMENT ON COLUMN "loans"."amount" IS 'must be positive';

ALTER TABLE "loans" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "logs" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
