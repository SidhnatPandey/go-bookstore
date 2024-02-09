package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Book struct {
	ID          int64  `gorm:"column:id" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func (b *Book) CreateBook() *Book {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	DB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	DB.Where("ID = ?", ID).Delete(book)
	return book
}
