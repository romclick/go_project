package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
	Postno    string    `gorm:"type:varchar(255)"`
	PostCount int       `gorm:"default:0;comment:'文章数量'"`
}

type Post struct {
	ID            uint      `gorm:"primary_key"`
	Title         string    `gorm:"type:varchar(255)"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"update_at"`
	Commentno     string    `gorm:"type:varchar(255)"`
	CommentStatus string    `gorm:"size:20;default:'无评论';comment:'评论状态'"`
}

type Comment struct {
	ID        uint      `gorm:"primary_key"`
	Content   string    `gorm:"type:text"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Username  string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	result := tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		return fmt.Errorf("更新文章数量失败", result.Error)
	}
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	//查询档案文章的有效评论
	var commentCount int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&commentCount).Error; err != nil {
		return fmt.Errorf("统计评论数失败", err)
	}
	//根据评论数更新文章状态
	var status string
	if commentCount == 0 {
		status = "无评论"
	} else {
		status = "有评论"
	}

	//更新文章评论状态
	if err := tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_status", status).Error; err != nil {
		return fmt.Errorf("更新评论状态失败", err)
	}
	return nil
}
func main() {
	dsn := "gorm:gorm1234!@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	db.AutoMigrate(User{}, Post{}, Comment{})
	//------测试钩子函数-------
	//1.创建测试用户
	user := User{Username: "user01", Email: "user01@exmple.com"}
	db.Create(&user)
	fmt.Printf("创建用户：%+v\n", user)

	//2.创建文章
	post := Post{Title: "测试文章01", Content: "钩子函数测试", UserID: user.ID}
	db.Create(&post)
	fmt.Printf("创建文章%+v\n", post)

	//3.验证用户文章数是否更新
	var updateUser User
	db.First(&updateUser, user.ID)
	fmt.Printf("用户文章数更新后: %+v\n", updateUser)

	//4.创建评论
	comment := Comment{Content: "测试评论01", PostID: user.ID, UserID: user.ID}
	db.Create(&comment)
	var postWithComment Post
	db.First(&postWithComment, post.ID)
	fmt.Printf("添加评论后文章状态：%+v\n", postWithComment)

	//5.删除评论
	db.Delete(&comment)
	var postAfterDelete Post
	db.First(&postAfterDelete, post.ID)
	fmt.Printf("删除评论后文章状态：%+v\n", postAfterDelete)
}
