package repository

import (
	"fmt"
	"simpleapi/helper"
	"simpleapi/model"
	"time"
)

type BookRepo interface {
	CreateBook(model.Book) (model.Book, error)
	GetBookByID(int) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
  DeleteBook(int) (string, error)
  UpdateBook(model.Book, int) (model.Book, error)
}

func (r Repo) CreateBook(book model.Book) (res model.Book, err error) {
  
	book.CreatedAt = time.Now()
  
  err = r.gorm.Create(&book).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookByID(id int) (res model.Book, err error) {
	err = r.gorm.First(&res, "id = ?", id).Error
  if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBooks() (res []model.Book, err error) {

	var books []model.Book
	
  err = r.gorm.Find(&books).Error

  if err != nil {
		return res, err
	}
  
	return books, nil
}

func (r Repo) DeleteBook(id int) (res string, err error) {

  qry := r.gorm.Where("id = ?", id).Delete(model.Book{})
  err = qry.Error

  if qry.RowsAffected < 1 {
    return res, helper.NotFoundErr
  }

  if err != nil {
		return res, err
	}

	return fmt.Sprintf("Book with id: %v deleted successfully", id), nil 

}

func (r Repo) UpdateBook(book model.Book, id int) (res model.Book, err error) {
  book.UpdatedAt = time.Now()
  qry := r.gorm.Where("id = ?", id).Updates(book)
  err = qry.Error

  if qry.RowsAffected < 1 {
    return res, helper.NotFoundErr
  }

  if err != nil {
		return res, err
	}
  
  //aneh
	err = r.gorm.First(&res, "id = ?", id).Error
  if err != nil {
		return res, err
	}

  return res, nil
}
