CREATE INDEX IF NOT EXISTS "sessions_index_0" ON "sessions" ("user_id");
CREATE TABLE IF NOT EXISTS "groups" (
    "group_id" VARCHAR NOT NULL,
    -- Foreign Key
    "user_id" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "description" TEXT,
    "creation_date" DATETIME NOT NULL,
    PRIMARY KEY("group_id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE CASCADE

);