CREATE INDEX IF NOT EXISTS "comments_index_0" ON "comments" ("user_id", "post_id");
CREATE TABLE IF NOT EXISTS "reports" (
    "post_id" VARCHAR NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "group_id" VARCHAR NOT NULL,
    "content" VARCHAR NOT NULL,
    "type" TEXT NOT NULL CHECK ("type" IN ('spam', 'abuse', 'other')),
    PRIMARY KEY ("post_id", "user_id", "group_id"),
    FOREIGN KEY ("post_id") REFERENCES "posts"("post_id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON DELETE CASCADE
);