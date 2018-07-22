package db

import(
	"github.com/jinzhu/gorm"
	//SQL Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDb() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if(err != nil) {
		panic(err)
	}
	return db
}