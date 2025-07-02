CREATE INDEX IF NOT EXISTS "notifications_index_0" ON "notifications" ("type", "user_id");
CREATE TABLE IF NOT EXISTS "posts" (
    "post_id" VARCHAR PRIMARY KEY,
    "group_id" VARCHAR,
    "user_id" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "content" TEXT NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "visibility" TEXT NOT NULL CHECK (
        "visibility" IN ('public', 'followers', 'private')
    ),
    "image" BLOB,
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON UPDATE NO ACTION ON DELETE
    SET NULL,
        FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE CASCADE
);