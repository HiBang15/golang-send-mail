CREATE TABLE IF NOT EXISTS "mail_templates" (
    "id" SERIAL PRIMARY KEY,
    "subject" VARCHAR NOT NULL,
    "content" VARCHAR NOT NULL,
    "created_at" timestamptz DEFAULT (now()),
    "deleted_at" timestamptz,
    "updated_at" timestamptz DEFAULT (now())
);