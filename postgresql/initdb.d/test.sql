CREATE DATABASE test_db;

\c test_db

CREATE TABLE "users" (
  "id" serial,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY(id)
);

CREATE TABLE "todos" (
  "id" serial,
  "text" text NOT NULL,
  "status" text NOT NULL,
  "user_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "finished_at" timestamptz,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("user_id") REFERENCES users ("id")
);

CREATE TABLE "labels" (
  "id" serial,
  "label" text NOT NULL,
  "todo_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  FOREIGN KEY ("todo_id") REFERENCES todos ("id")
);

insert into users (name) values ('test');
insert into todos (text, status, user_id) values 
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1),
  ('test', 'status', 1);
insert into labels (label, todo_id) values ('label1', 1);
