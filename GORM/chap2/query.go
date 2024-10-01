package main

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func query_main() {
	dsn := "root:Long5230413@tcp(127.0.0.1:3306)/archi_faculty?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: &time.Time{}}

	//检索单个对象

	// 获取第一条记录（主键升序）
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	res := db.First(&user)

	// 检查 ErrRecordNotFound 错误
	errors.Is(res.Error, gorm.ErrRecordNotFound)

	//========================
	// works because model is specified using `db.Model()`
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	db.Table("users").First(&result)

	// works with Take
	db.Table("users").Take(&result)

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type Language struct {
		Code string
		Name string
	}
	db.First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1

	//根据主键检索
	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	var users []User
	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	//如果主键是字符串
	db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	//检索全部对象
	// Get all records
	res = db.Find(&users)
	// SELECT * FROM users;

	//==============================================
	//String 条件
	// Get first matched record
	db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// Get all matched records
	db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	lastWeek := time.Now().Day() + 1
	today := time.Now().Day()
	db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	//============================================
	//Struct Map条件
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Slice of primary keys
	db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	//==============================================
	//指定结构体查询字段
	db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

	//====================================
	//内联条件
	// Get by primary key if it were a non-integer type
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;

	//=============================================
	//NOT条件
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	//==============================================
	//Or条件
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	//===========================================================
	//排序
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

	//========================================================
	//Limit & Offset
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	var users1 User
	var users2 User
	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// Cancel offset condition with -1
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

	//=========================================================
	//Group By & Having
	db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"
}
