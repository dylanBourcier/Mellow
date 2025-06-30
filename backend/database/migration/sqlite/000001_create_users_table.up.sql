CREATE TABLE IF NOT EXISTS "users" (
    "user_id" VARCHAR NOT NULL,
    "session_id" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL UNIQUE,
    "firstname" VARCHAR NOT NULL,
    "lastname" VARCHAR NOT NULL,
    "birthdate" DATETIME NOT NULL,
    "role" VARCHAR NOT NULL,
    -- optionnelle
    "avatar" BLOB,
    "creation_date" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- optionnelle
    "description" VARCHAR,
    PRIMARY KEY("user_id"),
    FOREIGN KEY ("user_id") REFERENCES "sessions"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user_id") REFERENCES "notifications"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);