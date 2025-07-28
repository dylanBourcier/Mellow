CREATE INDEX IF NOT EXISTS "reports_index_0" ON "reports" ("post_id", "user_id");
CREATE TABLE IF NOT EXISTS "follow_requests" (
    "request_id" VARCHAR PRIMARY KEY,
    "sender_id" VARCHAR NOT NULL,
    "receiver_id" VARCHAR,
    "group_id" VARCHAR,
    "status" BOOLEAN NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "type" TEXT NOT NULL CHECK ("type" IN ('user', 'group')),
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON DELETE CASCADE,
    FOREIGN KEY ("sender_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("receiver_id") REFERENCES "users"("user_id") ON DELETE CASCADE
);