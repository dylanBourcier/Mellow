CREATE TABLE IF NOT EXISTS events_response (
    event_id CHAR(32),
    user_id CHAR(32),
    status TEXT CHECK(status IN ('going', 'not_going')),
    vote TEXT,
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES events(event_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);