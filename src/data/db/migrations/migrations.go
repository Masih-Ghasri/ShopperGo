package migrations

import (
	"github.com/Masih-Ghasri/GolangBackend/config"
	constant "github.com/Masih-Ghasri/GolangBackend/constants"
	"github.com/Masih-Ghasri/GolangBackend/data/db"
	"github.com/Masih-Ghasri/GolangBackend/data/models"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.Getconfig())

func Up1() {
	database := db.GetDb()
	if database == nil {
		logger.Fatal(logging.Postgres, logging.Migration, "Database client is nil", nil)
	}

	CreateTables(database)
	createDefaultInformation(database)

}

func CreateTables(database *gorm.DB) {
	tables := []interface{}{}

	User := models.User{}
	Role := models.Role{}
	UserRole := models.UserRole{}
	Category := models.Category{}
	Comment := models.Comment{}
	Follow := models.Follow{}
	Like := models.Like{}
	Notification := models.Notification{}
	Post := models.Post{}

	//Base table
	tables = addNewTable(database, Role, tables)
	tables = addNewTable(database, User, tables)

	//another tables
	tables = addNewTable(database, UserRole, tables)
	tables = addNewTable(database, Category, tables)
	tables = addNewTable(database, Post, tables)
	tables = addNewTable(database, Comment, tables)
	tables = addNewTable(database, Follow, tables)
	tables = addNewTable(database, Like, tables)
	tables = addNewTable(database, Notification, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, "Can Not Create DataBase Tables", nil)
		return
	}
	logger.Info(logging.Postgres, logging.Migration, "Database migration complete", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constant.DefaultUserName, Name: "Masih", LastName: "Ghasi", Email: "masih@ghasi.com", PhoneNumber: "09383864535"}
	pass := "123456789"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name=?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username=?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserID: u.Id, RoleID: roleId}
		database.Create(&ur)
	}
}

func Down1() {}
