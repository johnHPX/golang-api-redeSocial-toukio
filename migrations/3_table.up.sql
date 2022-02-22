CREATE TABLE publication(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(500) not null,
    author_id int not null,
    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    create_at timestamp default current_timestamp
) ENGINE=INNODB;