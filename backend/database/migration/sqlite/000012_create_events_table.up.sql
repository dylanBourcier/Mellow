CREATE INDEX IF NOT EXISTS "messages_index_0" ON "messages" ("receiver_id", "sender_id");
CREATE TABLE IF NOT EXISTS "events" (
    "event_id" VARCHAR NOT NULL,
    -- Foreign Key
    "user_id" VARCHAR NOT NULL,
    -- Foreign Key
    "group_id" VARCHAR NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "event_date" DATETIME NOT NULL,
    "title" VARCHAR NOT NULL,
    PRIMARY KEY("event_id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON DELETE CASCADE,
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);