CREATE TABLE IF NOT EXISTS groups (
    group_id CHAR(32) PRIMARY KEY,
    user_id CHAR(32) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    creation_date DATETIME NOT NULL,
    visibility TEXT CHECK(visibility IN ('public', 'private')),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);