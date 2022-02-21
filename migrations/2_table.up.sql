CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB