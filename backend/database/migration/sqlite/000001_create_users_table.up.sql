CREATE TABLE IF NOT EXISTS "users" (
    "user_id" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL UNIQUE,
    "firstname" VARCHAR NOT NULL,
    "lastname" VARCHAR NOT NULL,
    "birthdate" DATETIME NOT NULL,
    "role" VARCHAR NOT NULL,
    -- optionnelle
    "image_url" VARCHAR,
    "creation_date" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- optionnelle
    "description" VARCHAR,
    "privacy" TEXT CHECK("privacy" IN ('public', 'private')) NOT NULL DEFAULT 'public',
    PRIMARY KEY("user_id")
);