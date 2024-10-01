package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

func create_main() {
	//创建记录
	dsn := "root:Long5230413@tcp(127.0.0.1:3306)/archi_faculty?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: &time.Time{}}
	result := db.Create(&user) //通过数据指针来创建

	result.Error = nil //防止编辑器报错

	//创建多项记录
	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: &time.Time{}},
		{Name: "Jackson", Age: 19, Birthday: &time.Time{}},
	}

	result = db.Create(users) // pass a slice to insert multiple row

	//用指定的字段创建记录
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

	db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

	//批量插入
	users = []*User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(users)

	for _, user := range users {
		fmt.Print(user.ID) // 1,2,3
	}

	//根据Map创建
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})

	// Do nothing on conflict
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// Update columns to default value on `id` conflict
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL

	// Use SQL expression
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
	}).Create(&users)
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count));

	// Update columns to new value on `id` conflict
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age); MySQL

	// Update all columns to new value on conflict except primary keys and those columns having default values from sql func
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&users)
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age), ...; MySQL
}
