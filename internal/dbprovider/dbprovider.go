package dbprovider

import (
	"encoding/json"
	"fmt"
	"log"

	model "github.com/deanonqq/microservice-user-balance/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MessageUser struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Email   string `json:"email"`
	Balance int64  `json:"balance"`
	// action string `json:"action"`
}

type PostMessageUser struct {
	Id            int64  `json:"id"`
	Sum           int64  `json:"sum"`
	Action        string `json:"action"`
	DestinationId int64  `json:"destination"`
}

type Manager interface {
	AddUser(usr *MessageUser) error
	GetBalance(id uint64) []byte
	AddBalance(id uint64, sum uint64) []byte
	SubBalance(id uint64, sum uint64) []byte
	// FindUser(usr *MessageUser) (*MessageUser, error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
}

func (mgr *manager) AddUser(usr *MessageUser) (err error) {
	var new_user model.User = model.User{Name: usr.Name, Age: usr.Age, Email: usr.Email, Balance: usr.Balance}
	fmt.Println(new_user)
	mgr.db.Create(&new_user)
	return
}

func (mgr *manager) GetBalance(id uint64) []byte {
	var usr model.User
	fmt.Println(id)
	mgr.db.First(&usr, "Id = ?", id)

	js, err := json.Marshal(MessageUser{int64(usr.Id), usr.Name, int64(usr.Age), usr.Email, usr.Balance})
	if err != nil {
		log.Fatal("Failed to find")
	}
	return js
}
func (mgr *manager) AddBalance(id uint64, sum uint64) []byte {
	var usr model.User
	fmt.Println(id)
	mgr.db.Model(&usr).Where("Id = ?", id).Update("Balance", gorm.Expr("Balance + ?", sum))

	mgr.db.First(&usr, "Id = ?", id)
	js, err := json.Marshal(MessageUser{int64(usr.Id), usr.Name, int64(usr.Age), usr.Email, usr.Balance})
	if err != nil {
		log.Fatal("Failed to find")
	}
	return js
}

func (mgr *manager) SubBalance(id uint64, sum uint64) []byte {
	var usr model.User
	mgr.db.First(&usr, "Id = ?", id)

	if usr.Balance-int64(sum) > 0 {
		mgr.db.Model(&usr).Where("Id = ?", id).Update("Balance", gorm.Expr("Balance - ?", sum))
		js, err := json.Marshal(MessageUser{int64(usr.Id), usr.Name, int64(usr.Age), usr.Email, usr.Balance})
		if err != nil {
			log.Fatal("Failed to find")
		}
		return js

	} else {
		js, err := json.Marshal("Balance is less than 0")
		if err != nil {
			log.Fatal("Failed to find")
		}
		return js
	}
}

// func (mgr *manager, key string) FindUser(usr *MessageUser) {

// 	return (MessageUser)
// }
