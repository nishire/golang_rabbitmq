package consume

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "root"
	// dbName := "todo"
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func ConnectToDB() {
// 	dsn := "root:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	fmt.Println("::::::::::::::::::::::::::", Db)
// }

// package consume

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	"github.com/nishire/golang_rabbitmq/consumer/model"
// )

// var DB *gorm.DB

// // DBConfig represents db configuration
// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

// func BuildDBConfig() *DBConfig {
// 	dbConfig := DBConfig{
// 		Host:     "localhost",
// 		Port:     3306,
// 		User:     "root",
// 		Password: "1234",
// 		DBName:   "first_go",
// 	}
// 	return &dbConfig
// }

// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }

// func CreateHotel(hotel *model.Hotel) error {
// 	if createError := DB.Create(hotel).Error; createError != nil {
// 		return createError
// 	}
// 	return nil
// }
