CREATE INDEX IF NOT EXISTS "follow_requests_index_0" ON "follow_requests" ("type");
CREATE TABLE IF NOT EXISTS "messages" (
    "message_id" VARCHAR NOT NULL,
    -- Foreign Key
    "sender_id" VARCHAR NOT NULL,
    -- Foreign Key
    "receiver_id" VARCHAR NOT NULL,
    "content" TEXT NOT NULL,
    "creation_date" DATETIME NOT NULL,
    "is_read" BOOLEAN NOT NULL DEFAULT 0,
    PRIMARY KEY("message_id"),
    FOREIGN KEY ("sender_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("sender_id") REFERENCES "groups"("group_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
