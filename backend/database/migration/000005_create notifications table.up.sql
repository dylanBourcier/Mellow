CREATE TABLE IF NOT EXISTS notifications (
    notification_id CHAR(32) PRIMARY KEY,
    user_id CHAR(32) NOT NULL,
    type TEXT CHECK(
        type IN ('follow', 'group_invite', 'event_created')
    ),
    seen BOOLEAN DEFAULT FALSE,
    creation_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);