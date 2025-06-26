CREATE TABLE IF NOT EXISTS posts (
    post_id CHAR(32) PRIMARY KEY,
    group_id CHAR(32),
    user_id CHAR(32) NOT NULL,
    title VARCHAR(50),
    content TEXT,
    creation_date DATETIME NOT NULL,
    visibility TEXT CHECK(visibility IN ('public', 'followers', 'private')),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);