package repositry

import (
	"red_board/pkg/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	client *gorm.DB
}

func NewDb() (Db, error) {
	client, err := gorm.Open("mysql", "root:root@tcp([mysql]:3306)/hoge?charset=utf8mb4&parseTime=true")
	client.AutoMigrate(&domain.Shelter{})
	return Db{client: client}, err
}
func (db *Db) Find(shell interface{}) {
	db.client.Find(shell)
}
func (db *Db) Where(key string, param ...interface{}) *Db {
	db.client = db.client.Where(key, param...)
	return db
}
func (db *Db) Close() {
	db.client.Close()
}
func (db *Db) Create(shell interface{}) {
	db.client.Create(shell)
}
func (db *Db) Delete(key string, src interface{}, param ...interface{}) bool {
	if db.client.Where(key, param).First(src).RecordNotFound() {
		return true
	}
	db.client.Delete(src)
	return false
}
func (db *Db) Having(key string, param ...interface{}) *Db {
	db.client = db.client.Having(key, param)
	return db
}

func (db *Db) Offset(key interface{}) *Db {
	db.client = db.client.Offset(key)
	return db
}

func (db *Db) Limit(key interface{}) *Db {
	db.client = db.client.Limit(key)
	return db
}
func (db *Db) Select(raw string, param ...interface{}) *Db {
	db.client = db.client.Select(raw, param...)
	return db
}
