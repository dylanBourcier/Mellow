CREATE INDEX IF NOT EXISTS "users_index_0" ON "users" ("email", "username");
CREATE TABLE IF NOT EXISTS "sessions" (
    "session_id" VARCHAR NOT NULL,
    -- Foreign Key
    "user_id" VARCHAR NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "last_activity" DATETIME NOT NULL,
    PRIMARY KEY("session_id") FOREIGN KEY ("user_id") REFERENCES "sessions"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);