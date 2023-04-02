package service

import "simpleapi/model"

type BookService interface {
	CreateBook(model.Book) (model.Book, error)
	GetBookByID(int) (model.Book, error)
	GetAllBooks() ([]*model.Book, error)
	DeleteBook(int) (string, error)
	UpdateBook(model.Book, int) (model.Book, error)
}

func (s *Service) CreateBook(book model.Book) (res model.Book, err error) {

  res, err = s.repo.CreateBook(book)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetBookByID(id int) (res model.Book, err error) {
  res, err = s.repo.GetBookByID(id)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *Service) GetAllBooks() (res []model.Book, err error) {
	res, err = s.repo.GetAllBooks()
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *Service) DeleteBook(id int) (res string, err error) {
	res, err = s.repo.DeleteBook(id)
	if err != nil {
		return "", err
	}

	return res, nil

}

func (s *Service) UpdateBook(book model.Book, id int) (res model.Book, err error) {
	res, err = s.repo.UpdateBook(book, id)
	if err != nil {
		return res, err
	}
	return res, nil

}
