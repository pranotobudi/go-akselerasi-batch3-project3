package users

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"gorm.io/gorm"
)

func NewUserSeeder(db *gorm.DB) {
	CreateDBTable(db) //create table first, then seeder
	DBSeedUser(db)
}
func DBSeedUser(db *gorm.DB) error {
	UserDataSeed(db)
	// AdminDataSeed(db)
	// ReaderDataSeed(db)
	// CategoryDataSeed(db)
	// NewsDataSeed(db)
	// NewsReaderDataSeed(db)
	// NewsCommentDataSeed(db)
	return nil
}

func UserDataSeed(db *gorm.DB) {
	statement := "INSERT INTO users (role, email, password, username, profile_pic, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, "admin", "admin1@gmail.com", common.GeneratePassword("admin1"), "admin1", "http://profile_pic_url_admin1.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "user", "user1@gmail.com", common.GeneratePassword("user1"), "user1", "http://profile_pic_url_user1.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "guest", "guest1@gmail.com", common.GeneratePassword("guest1"), "guest1", "http://profile_pic_url_guest1.jpg", faker.Timestamp(), faker.Timestamp())
}

func CreateDBTable(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Event{}, &Transaction{}, &Registration{})
	db.AutoMigrate(UserRegistration{}, User{})

	// Create Fresh UserRegistration Table
	if (db.Migrator().HasTable(&UserRegistration{})) {
		fmt.Println("UserRegistration table exist")
		db.Migrator().DropTable(&UserRegistration{})
	}
	db.Migrator().CreateTable(&UserRegistration{})

	// Create Fresh User Table
	if (db.Migrator().HasTable(&User{})) {
		fmt.Println("User table exist")
		db.Migrator().DropTable(&User{})
	}
	db.Migrator().CreateTable(&User{})

}
