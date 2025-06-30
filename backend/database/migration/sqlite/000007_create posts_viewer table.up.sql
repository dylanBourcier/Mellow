CREATE INDEX IF NOT EXISTS "posts_index_0" ON "posts" ("visibility", "group_id", "user_id");
CREATE TABLE IF NOT EXISTS "posts_viewer" (
    "post_id" VARCHAR NOT NULL,
    "user_id" VARCHAR NOT NULL,
    PRIMARY KEY ("post_id", "user_id"),
    FOREIGN KEY ("post_id") REFERENCES "posts"("post_id") ON UPDATE NO ACTION ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE CASCADE
);