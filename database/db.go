package database

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Server   string = "localhost" // 服务器地址，通常是localhost或IP地址
	Port     string = "1433"      // 端口号，对于SQL Server通常是1433
	User     string = "user_yqb"  // 数据库用户名
	Password string = "123456"    // 数据库密码
	Database string = "my_oneday"
)

func ConnectSql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s;",
		Server,
		Port, // 确保Port已经被定义为一个字符串或者可以转换为字符串的类型
		Database,
		User,
		Password,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		// 处理错误，不要使用panic
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
