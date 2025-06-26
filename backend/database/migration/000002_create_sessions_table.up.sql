CREATE TABLE IF NOT EXISTS sessions (
    session_id CHAR(32) PRIMARY KEY,
    user_id CHAR(32) NOT NULL,
    creation_date DATETIME NOT NULL,
    last_activity DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);