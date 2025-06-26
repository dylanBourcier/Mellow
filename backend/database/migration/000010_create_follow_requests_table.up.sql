CREATE TABLE IF NOT EXISTS follow_requests (
    sender_id CHAR(32),
    receiver_id CHAR(32),
    group_id CHAR(32),
    status BOOLEAN,
    creation_date DATETIME NOT NULL,
    type TEXT CHECK(type IN ('user', 'group')),
    PRIMARY KEY (sender_id, receiver_id),
    FOREIGN KEY (sender_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES users(user_id)
);