CREATE INDEX IF NOT EXISTS "groups_index_0" ON "groups" ("user_id");
CREATE TABLE IF NOT EXISTS "groups_member" (
    "group_id" VARCHAR NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "role" VARCHAR,
    "join_date" DATETIME NOT NULL,
    PRIMARY KEY ("group_id", "user_id"),
    FOREIGN KEY ("group_id") REFERENCES "groups"("group_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);