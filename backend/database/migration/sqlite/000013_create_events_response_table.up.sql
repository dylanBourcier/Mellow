CREATE INDEX IF NOT EXISTS "events_index_0" ON "events" ("group_id", "user_id");
CREATE TABLE IF NOT EXISTS "events_response" (
    "event_id" VARCHAR NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "group_id" VARCHAR NOT NULL,
    "status" TEXT NOT NULL CHECK("status" IN ('going', 'not_going')),
    "vote" TEXT,
    PRIMARY KEY ("event_id", "user_id"),
    FOREIGN KEY ("event_id") REFERENCES "events"("event_id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS "events_response_index_0" ON "events_response" ("event_id", "user_id");