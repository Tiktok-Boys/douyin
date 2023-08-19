DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS follow;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS video;
DROP TABLE IF EXISTS user;

-- 用户表
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '用户名称',
    avatar VARCHAR(255) COMMENT '用户头像',
    background_image VARCHAR(255) COMMENT '用户个人页顶部大图',
    signature TEXT COMMENT '个人简介',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间'
);

-- 视频表
CREATE TABLE video (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT COMMENT '视频作者',
    play_url VARCHAR(255) COMMENT '视频播放地址',
    cover_url VARCHAR(255) COMMENT '视频封面地址',
    title VARCHAR(255) COMMENT '视频标题',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE SET NULL,
    INDEX idx_videos_user_id (user_id), -- 需要频繁通过用户ID查找该用户的视频列表
    INDEX idx_videos_created_at (created_at) -- 需要频繁通过发布时间的范围查找视频
);

-- 点赞关联表（多对多关系）
CREATE TABLE likes (
    user_id INT,
    video_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, video_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (video_id) REFERENCES video(id) ON DELETE CASCADE
);

-- 评论关联表（多对多关系）
CREATE TABLE comment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    video_id INT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (video_id) REFERENCES video(id) ON DELETE CASCADE,
    INDEX idx_comments_video_id (video_id) -- 需要频繁通过视频ID查看该视频下的评论列表
);

-- 关注关联表（多对多关系）
CREATE TABLE follow (
    follower_id INT,
    followee_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followee_id),
    FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES user(id) ON DELETE CASCADE
);

-- 发消息关联表（多对多关系）
CREATE TABLE message (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sender_id INT,
    receiver_id INT,
    content TEXT,
    created_at INT,
    FOREIGN KEY (sender_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX idx_messages_sender_receiver (sender_id, receiver_id), -- 需要频繁查找两个人之间的聊天记录
    INDEX idx_messages_created_at (created_at) -- 需要根据消息生成时间排序
);
