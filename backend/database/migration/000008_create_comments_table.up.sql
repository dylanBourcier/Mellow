CREATE TABLE IF NOT EXISTS comments (
    comment_id CHAR(32) PRIMARY KEY,
    user_id CHAR(32) NOT NULL,
    post_id CHAR(32) NOT NULL,
    content TEXT,
    creation_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
);