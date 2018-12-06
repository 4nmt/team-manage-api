-- Adminer 4.6.3 PostgreSQL dump


DROP TABLE IF EXISTS "projects";
DROP SEQUENCE IF EXISTS project_id_seq;
CREATE SEQUENCE project_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."projects" (
    "id" integer DEFAULT nextval('project_id_seq') NOT NULL,
    "name" text,
    "description" text,
    "created_at" integer,
    "updated_at" integer,
    CONSTRAINT "projects_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('user_id_seq') NOT NULL,
    "name" text,
    "email" text,
    "created_at" integer,
    "updated_at" integer,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

DROP TABLE IF EXISTS "user_project";
DROP SEQUENCE IF EXISTS user_project_id_seq;
CREATE SEQUENCE user_project_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."user_project" (
    "id" integer DEFAULT nextval('user_project_id_seq') NOT NULL,
    "user_id" integer NOT NULL,
    "project_id" integer NOT NULL,
    "created_at" integer,
    "updated_at" integer,
    CONSTRAINT "user_project_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "user_project_to_book_id_fkey" FOREIGN KEY (project_id) REFERENCES projects(id) NOT DEFERRABLE,
    CONSTRAINT "user_project_to_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE
) WITH (oids = false);


-- 2018-12-06 11:16:01.214315+00