CREATE INDEX IF NOT EXISTS "groups_member_index_0" ON "groups_member" ("group_id", "user_id");
CREATE TABLE IF NOT EXISTS "notifications" (
    "notification_id" VARCHAR PRIMARY KEY,
    "user_id" VARCHAR NOT NULL,
    "type" TEXT NOT NULL CHECK(
        "type" IN ('follow', 'group_invite', 'event_created')
    ),
    "seen" BOOLEAN NOT NULL DEFAULT false,
    "creation_date" DATETIME NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);