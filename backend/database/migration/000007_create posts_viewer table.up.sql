CREATE TABLE IF NOT EXISTS posts_viewer (
    post_id CHAR(32),
    user_id CHAR(32),
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);