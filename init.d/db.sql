CREATE DATABASE news CHARACTER SET utf8;

use news;

CREATE TABLE news_item(
    id int PRIMARY KEY AUTO_INCREMENT,
    title varchar(255),
    description varchar(255),
    body text,
    type varchar(255),
    link varchar(255),
    source varchar(255),
    updated_at datetime
);