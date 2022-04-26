-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS id_seq_module;

-- Table Definition
CREATE TABLE "public"."module"
(
    "id"         int4         NOT NULL DEFAULT nextval('id_seq_module'::regclass),
    "name"       varchar(255) NOT NULL,
    "version"    varchar(255) NOT NULL,
    "status"     varchar(255) NOT NULL,
    "updated_on" timestamptz,
    PRIMARY KEY ("id")
);

-- insert data
INSERT INTO "public"."module" ("name", "version", "status", "updated_on")
VALUES ('ciCd', 'unknown', 'unknown', 'now()');

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS id_seq_module_action_audit_log;

-- Table Definition
CREATE TABLE "public"."module_action_audit_log"
(
    "id"         int4         NOT NULL DEFAULT nextval('id_seq_module_action_audit_log'::regclass),
    "module_id"  varchar(255) NOT NULL,
    "version"    varchar(255) NOT NULL,
    "action"     varchar(255) NOT NULL,
    "created_on" timestamptz  NOT NULL,
    "created_by" int4         NOT NULL,
    PRIMARY KEY ("id")
);

-- add FK
ALTER TABLE "public"."module_action_audit_log"
    ADD FOREIGN KEY ("module_id") REFERENCES "public"."module" ("id");

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS id_seq_server_action_audit_log;

-- Table Definition
CREATE TABLE "public"."server_action_audit_log"
(
    "id"         int4         NOT NULL DEFAULT nextval('id_seq_server_action_audit_log'::regclass),
    "action"     varchar(255) NOT NULL,
    "version"    varchar(255),
    "created_on" timestamptz  NOT NULL,
    "created_by" int4         NOT NULL,
    PRIMARY KEY ("id")
);