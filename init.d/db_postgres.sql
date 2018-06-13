CREATE TABLE news_items(
    id serial PRIMARY KEY,
    title varchar(255),
    description varchar(255),
    body text,
    type varchar(255),
    link varchar(255),
    image varchar(255),
    source varchar(255),
    updated_at timestamp
);

CREATE TABLE users(
    id serial PRIMARY KEY,
    nickname varchar(255),  
    pwd varchar(255),
    phone varchar(11), 
    email varchar(128), 
    avatar varchar(255), 
    gender int, 
    biography text,  -- 个人简介
    created_at timestamp, 
    login_date timestamp,
    real_name varchar(20), 
    identity_card_num varchar(19), 
    identity_card_front varchar(256),
    identity_card_back varchar(256),
    from_code varchar(6), -- 专家从哪个邀请码注册的
    license varchar(256), -- 卖家营业执照, 必填？
    expertise varchar(256), -- 专家擅长领域, 必填？
    resume varchar(512), -- 专家简历, 必填?
    role int, -- 普通用户1，卖家2，专家3
    is_verified boolean -- 审核通过
);

CREATE TABLE  authentications(
    id serial PRIMARY KEY,
    user_id int, 
    uuid varchar(128), 
    token varchar(255)
);

CREATE TABLE captchas(
    id serial PRIMARY KEY,
    phone varchar(11),
    code varchar(6)
);

CREATE TABLE invitations(
    id serial PRIMARY KEY,
    invitation_code varchar(6), 
    has_activated boolean 
);

CREATE TABLE suggestions(
    id serial PRIMARY KEY,
    user_id int,
    content text
);

INSERT INTO invitations (invitation_code, has_activated) VALUES ('111111', false) RETURNING id;
INSERT INTO invitations (invitation_code, has_activated) VALUES ('222222', true) RETURNING id;

-- 求助
CREATE TABLE help_requests(
    id serial PRIMARY KEY,
    user_id int,
    title varchar(255),
    content text,
    amount int, -- 悬赏金额
    created_at timestamp
);

-- 求购
CREATE TABLE buy_requests(
    id serial PRIMARY KEY,
    user_id int,
    title varchar(255),
    content text,
    amount int, -- 求购数量， -1表示不限
    created_at timestamp
);

-- 芯片
CREATE table chips (
    id serial PRIMARY KEY,
    user_id int REFERENCES users,
    serial_number varchar(100),
    vendor varchar(255),
    amount int,
    manufacture_date timestamp, 
    unit_price float,
    specification varchar(255),
    is_verified boolean -- 审核通过
);

-- 收藏
CREATE table favorites (
    id serial PRIMARY KEY,
    user_id int,
    favorable_type varchar(20),
    favorable_id int,
    created_at timestamp
);

-- 评论
CREATE table comments (
    id serial PRIMARY KEY,
    user_id int,
    commentable_type varchar(20),
    commentable_id int,
    content text,
    is_picked boolean, -- 求助选择为正确答案可以获得赏金
    likes int, -- 点赞数
    created_at timestamp
);

-- 点赞 
CREATE table likes (
    id serial PRIMARY KEY,
    user_id int,
    comment_id int,
    created_at timestamp
);

-- CREATE TRIGGER sum_likes AFTER INSERT ON likes(
  
-- )
