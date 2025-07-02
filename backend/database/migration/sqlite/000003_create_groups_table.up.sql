CREATE INDEX IF NOT EXISTS "sessions_index_0" ON "sessions" ("user_id");
CREATE TABLE IF NOT EXISTS "groups" (
    "group_id" VARCHAR NOT NULL,
    -- Foreign Key
    "user_id" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "description" TEXT,
    "creation_date" DATETIME NOT NULL,
    "visibility" TEXT NOT NULL CHECK(
        "visibility" IN ('public', 'private', 'hidden')
    ),
    PRIMARY KEY("group_id")
);