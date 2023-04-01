package repository

import (
	"fmt"
	"simpleapi/helper"
	"simpleapi/model"
	"simpleapi/repository/query"
)

type BookRepo interface {
	CreateBook(model.Book) (model.Book, error)
	GetBookByID(int) (model.Book, error)
	GetAllBooks() ([]*model.Book, error)
  DeleteBook(int) (string, error)
  UpdateBook(model.Book, int) (string, error)
}

func (r Repo) CreateBook(book model.Book) (res model.Book, err error) {

	err = r.db.QueryRow(
		query.CreateBook,
		book.Title,
		book.Author,
		book.Description,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookByID(id int) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.GetBookByID,
		id,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBooks() (res []*model.Book, err error) {

	var books []*model.Book
	rows, err := r.db.Query(query.GetAllBooks)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			return res, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (r Repo) DeleteBook(id int) (res string, err error) {

	qry, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return res, err
	}

	count, err := qry.RowsAffected()
	if err != nil {
		return res, err
	}

  if count < 1 {
    return res, helper.NotFoundErr 
  }

	return fmt.Sprintf("Deleted data amount: %v", count), nil 

}

func (r Repo) UpdateBook(book model.Book, id int) (res string, err error) {
  
	qry, err := r.db.Exec(query.UpdateBook, id, book.Title, book.Author, book.Description)
	if err != nil {
		return res, err
	}

	count, err := qry.RowsAffected()
	if err != nil {
		return res, err
	}

  if count < 1 {
    return res, helper.NotFoundErr
  }

	return fmt.Sprintf("Updated data amount: %v", count), err

}
