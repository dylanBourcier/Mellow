CREATE INDEX IF NOT EXISTS "users_index_0" ON "users" ("email", "username");
CREATE INDEX IF NOT EXISTS "idx_sessions_user_id" ON "sessions"("user_id");
CREATE INDEX IF NOT EXISTS "idx_sessions_last_activity" ON "sessions"("last_activity");

CREATE TABLE IF NOT EXISTS "sessions" (
    "session_id" VARCHAR NOT NULL,
    -- Foreign Key
    "user_id" VARCHAR NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "last_activity" DATETIME NOT NULL,
    PRIMARY KEY("session_id") FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
