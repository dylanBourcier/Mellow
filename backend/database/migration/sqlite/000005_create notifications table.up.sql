CREATE INDEX IF NOT EXISTS "groups_member_index_0" ON "groups_member" ("group_id", "user_id");
CREATE TABLE IF NOT EXISTS "notifications" (
    "notification_id" VARCHAR PRIMARY KEY,
    "user_id" VARCHAR NOT NULL,
    "sender_id" VARCHAR NOT NULL,
    "request_id" VARCHAR,
    "group_id" VARCHAR,
    "type" TEXT NOT NULL,
    "seen" BOOLEAN NOT NULL DEFAULT false,
    "creation_date" DATETIME NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
    FOREIGN KEY("request_id") REFERENCES "follow_requests"("request_id") ON UPDATE NO ACTION ON DELETE SET NULL
);