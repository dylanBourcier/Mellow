CREATE TABLE IF NOT EXISTS groups_member (
    group_id CHAR(32),
    user_id CHAR(32),
    role VARCHAR(20),
    join_date DATETIME,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);