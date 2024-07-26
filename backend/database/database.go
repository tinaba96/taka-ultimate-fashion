package database

import (
	"fmt"

	"github.com/tinaba96/taka-ultimate-fashion/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type DBConfig struct {
// 	User string
// 	Password string
// 	Host string
// 	Port int
// 	Table string
// }

// func getDBConfig() DBConfig {
//     port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
//     return DBConfig{
//         User: os.Getenv("MYSQL_USER"),
//         Password: os.Getenv("MYSQL_PASSWORD"),
//         Host: os.Getenv("MYSQL_HOST"),
//         Port: port,
// 		Table: os.Getenv("MYSQL_DB"),
//     }
// }

var db *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", c.User, c.Password, c.Host, c.Port, c.Table)
	// dsn := "user:password@tcp(db:3306)/taka_database?charset=utf8mb4&parseTime=True&loc=Local"
	_, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errdb != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&Product{})



    // if isReset {
        // db.DropTableIfExists(models)
    // }
    // d.AutoMigrate(models...)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	c := config.GetDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", c.User, c.Password, c.Host, c.Port, c.Table)
	db, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errdb != nil {
		panic("failed to connect database")
	}
    return db
}

// Close closes database
// func Close() {
//     db.Close()
// }
