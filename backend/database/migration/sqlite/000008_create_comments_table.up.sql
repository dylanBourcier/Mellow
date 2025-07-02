CREATE INDEX IF NOT EXISTS "posts_viewer_index_0" ON "posts_viewer" ("post_id", "user_id");
CREATE TABLE IF NOT EXISTS "comments" (
    "comment_id" VARCHAR PRIMARY KEY,
    "user_id" VARCHAR NOT NULL,
    "post_id" VARCHAR NOT NULL,
    "content" TEXT NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "image_url" VARCHAR,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("post_id") REFERENCES "posts"("post_id") ON DELETE CASCADE
);