CREATE TABLE IF NOT EXISTS "group_join_requests" (
    "id" VARCHAR PRIMARY KEY,
    "group_id" VARCHAR NOT NULL,
    "requester_id" VARCHAR NOT NULL,
    "status" TEXT NOT NULL CHECK(status IN ('pending','accepted','rejected','cancelled')),
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "decided_at" DATETIME,
    "decided_by" VARCHAR,
    FOREIGN KEY("group_id") REFERENCES "groups"("group_id") ON DELETE CASCADE,
    FOREIGN KEY("requester_id") REFERENCES "users"("user_id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_gjr_group ON group_join_requests(group_id);
CREATE INDEX IF NOT EXISTS idx_gjr_requester ON group_join_requests(requester_id);
CREATE INDEX IF NOT EXISTS idx_gjr_status ON group_join_requests(status);

-- enforce only one pending per (group, requester)
CREATE UNIQUE INDEX IF NOT EXISTS idx_gjr_unique_pending 
ON group_join_requests(group_id, requester_id)
WHERE status = 'pending';
