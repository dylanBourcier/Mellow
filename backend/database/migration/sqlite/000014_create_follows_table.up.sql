CREATE TABLE IF NOT EXISTS "follows" (
    "follower_id" VARCHAR NOT NULL,
    "followed_id" VARCHAR NOT NULL,
    "creation_date" DATETIME NOT NULL,
    PRIMARY KEY ("follower_id", "followed_id"),
    FOREIGN KEY ("follower_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("followed_id") REFERENCES "users"("user_id") ON DELETE CASCADE
);