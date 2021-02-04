package main

import (
	"database/sql"
	"fmt"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// User struct
type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string  `gorm:"type:varchar(100);unique_index"`
	Serial   *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Address  string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe int     `gorm:"-"`               // 忽略本字段
}

// Animal struct
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func main() {
	db, err := gorm.Open("mysql", "root:admin123@/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(false) // 默认关闭
	db.SingularTable(true)
	if err != nil {
		panic("DB connect failed...")
	}
	defer db.Close()

	fmt.Println("DB connected...")
	time := time.Now()
	seri := uuid.New().String()
	user := User{Name: "Tom", Age: sql.NullInt64{Int64: int64(20)}, Birthday: &time, Email: "tom@gmail.com", Serial: &seri}
	exist := db.HasTable(user)
	if exist {
		// db.DropTable(&user)
		user = User{}
	} else {
		db.AutoMigrate(&user)
		db.Create(&user)
	}

	err = db.Where("Email = ?", "tom@gmail.com").Find(&user).Error
	if err != nil {
		fmt.Println("Find Error: ", err)
	}

	fmt.Println("Find by Email: ", user.ID, user)
	err = db.Delete(&user).Error
	if err != nil {
		fmt.Println(err)
	}

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}
