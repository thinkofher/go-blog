SET TIMEZONE='CET';

CREATE TABLE blog_user (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (120) NOT NULL,
    email VARCHAR (355) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

CREATE TABLE post (
    post_id SERIAL,
    author_id SERIAL REFERENCES blog_user(user_id),
    body TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL,
    PRIMARY KEY (post_id, author_id)
);
