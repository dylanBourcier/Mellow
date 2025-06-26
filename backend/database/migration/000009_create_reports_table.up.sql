CREATE TABLE IF NOT EXISTS reports (
    post_id CHAR(32),
    user_id CHAR(32),
    content VARCHAR(250),
    type TEXT CHECK(type IN ('spam', 'abuse', 'other')),
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);