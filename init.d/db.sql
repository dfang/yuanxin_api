DROP DATABASE IF EXISTS news;

CREATE DATABASE news CHARACTER SET utf8;

USE news;

CREATE TABLE news_items(
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

CREATE TABLE users(
    id int PRIMARY KEY AUTO_INCREMENT,
    nickname varchar(255),  
    pwd varchar(255),
    phone varchar(11), 
    email varchar(128), 
    avatar varchar(255), 
    gender int, 
    biography text,  -- 个人简介
    created_at datetime, 
    login_date datetime,
    real_name varchar(20), 
    identity_card_num varchar(19), 
    identity_card_front varchar(256),
    identity_card_back varchar(256),
    from_code varchar(6), -- 专家从哪个邀请码注册的
    license varchar(256), -- 卖家营业执照, 必填？
    expertise varchar(256), -- 专家擅长领域, 必填？
    resume varchar(512), -- 专家简历, 必填?
    role int, -- 普通用户1，卖家2，专家3
    is_verified boolean, -- 审核通过
    remark text, -- 备注
    acc_id varchar(100), -- im用户名
    acc_token varchar(100) -- im密码
);

CREATE TABLE  authentications(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int, 
    uuid varchar(128), 
    token varchar(255)
);

CREATE TABLE captchas(
    id int PRIMARY KEY AUTO_INCREMENT,
    phone varchar(11),
    code varchar(6)
);

CREATE TABLE invitations(
    id int PRIMARY KEY AUTO_INCREMENT,
    invitation_code varchar(6), 
    has_activated boolean 
);

CREATE TABLE suggestions(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    content text
);

insert into invitations(invitation_code, has_activated) values("111111", false);
insert into invitations(invitation_code, has_activated) values("222222", true);

-- 求助
CREATE TABLE help_requests(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    title varchar(255),
    content text,
    amount int, -- 悬赏金额
    created_at datetime
);

-- 求购
CREATE TABLE buy_requests(
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    title varchar(255),
    content text,
    amount int, -- 求购数量， -1表示不限
    created_at datetime
);

-- 芯片
CREATE table chips (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    serial_number varchar(100),
    vendor varchar(255),
    amount int,
    manufacture_date datetime, 
    unit_price float,
    specification varchar(255),
    is_verified boolean, -- 审核通过
    version varchar(255),
    volume int
);

-- 收藏
CREATE table favorites (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    favorable_type varchar(20),
    favorable_id int,
    created_at datetime
);

-- 评论
CREATE table comments (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    commentable_type varchar(20),
    commentable_id int,
    content text,
    is_picked boolean, -- 求助选择为正确答案可以获得赏金
    likes int, -- 点赞数
    created_at datetime
);

-- 点赞 
CREATE table likes (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    comment_id int,
    created_at datetime
)

-- CREATE TRIGGER sum_likes AFTER INSERT ON likes(
  
-- )
