CREATE TABLE IF NOT EXISTS events (
    event_id CHAR(32) PRIMARY KEY,
    user_id CHAR(32) NOT NULL,
    group_id CHAR(32) NOT NULL,
    creation_date DATETIME NOT NULL,
    event_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);