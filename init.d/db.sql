CREATE DATABASE news CHARACTER SET utf8;

use news;

CREATE TABLE news_item(
    id int PRIMARY KEY AUTO_INCREMENT,
    title varchar(255),
    description varchar(255),
    body text,
    type varchar(255),
    link varchar(255),
    image varchar(255),
    source varchar(255),
    updated_at datetime
);


-- 所有
-- 产业报道               
-- 厂商动态               
-- 数码相机/摄像机        
-- 智能家电               
-- 智能手机               
-- 电脑
CREATE TABLE user(
    id int PRIMARY KEY AUTO_INCREMENT,
    nickname varchar(255),  
    pwd varchar(56),  
    phone varchar(11), 
    email varchar(128), 
    avantar varchar(255), 
    sex int, 
    created_at datetime, 
    login_date datetime 
);

CREATE TABLE  authentication(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int, 
    uuid varchar(128), 
    token varchar(255)
);

CREATE TABLE user_role(
    id int PRIMARY KEY AUTO_INCREMENT,
    role_id int,
    user_id int  
);

CREATE TABLE role(
    id int PRIMARY KEY AUTO_INCREMENT,
    real_name varchar(20), 
    identity_card_num varchar(19), 
    identity_card_front varchar(256),
    identity_card_end varchar(256), 
    license varchar(256),
    expertise varchar(256) 
);

CREATE TABLE captcha(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int , 
    captcha_code varchar(6), 
);

CREATE TABLE invitation(
    id int PRIMARY KEY AUTO_INCREMENT,
    invitation_code varchar(6), 
    has_activated boolean 
);

