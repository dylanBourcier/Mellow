CREATE TABLE IF NOT EXISTS messages (
    message_id CHAR(32) PRIMARY KEY,
    sender_id CHAR(32) NOT NULL,
    receiver_id CHAR(32) NOT NULL,
    content TEXT,
    creation_date DATETIME NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES users(user_id)
);