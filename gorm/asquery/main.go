package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type User struct {
	ID       uint      `gorm:"primary_key"`
	Username string    `gorm:"type:varchar(255)"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Postno   string    `gorm:"type:varchar(255)"`
}

type Post struct {
	ID        uint      `gorm:"primary_key"`
	Content   string    `gorm:"type:text"`
	Userid    uint      `gorm:"type:varchar(255)"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	Commentno string    `gorm:"type:varchar(255)"`
}

type Comment struct {
	ID       uint      `gorm:"primary_key"`
	Content  string    `gorm:"type:text"`
	Postid   uint      `gorm:"type:varchar(255)"`
	Userid   uint      `gorm:"type:varchar(255)"`
	Username string    `gorm:"type:varchar(255)"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func initDB() (*gorm.DB, error) {
	dsn := "gorm:gorm1234!@tcp(127.0.0.1:3306)/gorm?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 表名默认复数（users/posts/comments），与模型对应
		},
	})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取连接失败", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return db, nil
}

func getUserPwithC(db *gorm.DB, userid uint) (*User, error) {
	var user User

	result := db.WithContext(context.Background()).
		Preload("Posts").
		Preload("Posts.Comments").
		Preload("Posts.Comments.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "create_at")
		}).
		First(&user, "id = ?", userid)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在", userid)
		}
		return nil, fmt.Errorf("获取文章失败", result.Error)
	}
	return &user, nil
}

func getUserPwithMostC(db *gorm.DB) ([]Post, error) {
	var posts []Post
	subQuery := db.WithContext(context.Background()).
		Model(&Post{}).
		Select("post_id, count(*) as comment_count").
		Group("post_id").
		Having("count(*) = (SELECT count(*) FROM comments GROUP BY post_id ORDER BY comment_count DESC LIMIT 10)").
		SubQuery()

	result := db.WithContext(context.Background()).
		Model(&Post{}).
		Select("post, sub.comment_count").
		Joins("join ? as sub on post.id = sub.post_id", subQuery).
		Preload("user", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "create_at")
		}).
		Find(&posts)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return posts, fmt.Errorf("暂无文章数据")
		}
		return posts, fmt.Errorf("查询最多评论文章失败", result.Error)
	}
	return posts, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal("初始化失败", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	userid := uint(0)
	user, err := getUserPwithC(db, userid)
	if err != nil {
		log.Printf("执行1失败", err)
	} else {
		log.Printf("\n执行1结果：%s（ID：%d）的文章及评论", user.Username, userid)
		for _, post := range user.Posts {
			fmt.Printf("文章内容", post.Context)
			fmt.Printf("评论数%s"，len(post.Comments))
			for _, comment := range post.Comments {
				fmt.Printf("评论者",comment.User.Username,comment.Content,)
			}
			fmt.Printf("\n"+"----")
		}
	}

	topPosts, err := getUserPwithMostC(db)
	if err != nil {
		log.Printf("执行2失败"，err)
	}else {
		log.Printf("评论最多的文章数量：",len(topPosts))
		for _, topPost := range topPosts {
			fmt.Printf("文章ID：%d，标题：%s",post.ID,post.title)
		}
	}
}
