package service

import (
	"errors"
	"simpleapi/model"
	"simpleapi/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_BookService_CreateBook(t *testing.T) {

	type testCase struct {
		name           string
		input          model.Book
		expectedResult model.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name: "Success",
		input: model.Book{
			ID:          1,
			Title:       "title",
			Author:      "author",
			Description: "description",
		},
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "title",
				Author:      "author",
				Description: "description",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "title",
			Author:      "author",
			Description: "description",
		},
	})

	testTable = append(testTable, testCase{
		name:          "Invalid Error",
    input: model.Book{
			ID:          2,
			Title:       "ti",
			Author:      "au",
			Description: "de",
		},
		wantError:     true,
		expectedError: errors.New("book_author: the length must be between 5 and 50; book_title: the length must be between 5 and 50; description: the length must be between 5 and 50."),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{repo: bookRepo}
			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_GetAllBooks(t *testing.T) {

	type testCase struct {
		name           string
		expectedResult []model.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBooks().Return([]model.Book{
				{
					ID:          1,
					Title:       "title",
					Author:      "author",
					Description: "description",
				},
				{
					ID:          2,
					Title:       "title2",
					Author:      "author2",
					Description: "description2",
				},
			}, nil).Times(1)
		},
		expectedResult: []model.Book{
			{
				ID:          1,
				Title:       "title",
				Author:      "author",
				Description: "description",
			},
			{
				ID:          2,
				Title:       "title2",
				Author:      "author2",
				Description: "description2",
			},
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{repo: bookRepo}
			res, err := service.GetAllBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_GetBookByID(t *testing.T) {

	type testCase struct {
		name           string
		expectedResult model.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "title",
				Author:      "author",
				Description: "description",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "title",
			Author:      "author",
			Description: "description",
		},
	})

	testTable = append(testTable, testCase{
		name:          "Not Found",
		wantError:     true,
		expectedError: gorm.ErrRecordNotFound,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, gorm.ErrRecordNotFound).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{repo: bookRepo}
			res, err := service.GetBookByID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_UpdateBook(t *testing.T) {

	type testCase struct {
		name           string
		input          model.Book
		expectedResult model.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name: "Success",
		input: model.Book{
			ID:          1,
			Title:       "title updated",
			Author:      "author",
			Description: "description updated",
		},
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "title updated",
				Author:      "author",
				Description: "description updated",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "title updated",
			Author:      "author",
			Description: "description updated",
		},
	})

	testTable = append(testTable, testCase{
		name:          "Invalid Error",
    input: model.Book{
			ID:          2,
			Title:       "ti",
			Author:      "au",
			Description: "de",
		},
		wantError:     true,
		expectedError: errors.New("book_author: the length must be between 5 and 50; book_title: the length must be between 5 and 50; description: the length must be between 5 and 50."),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{repo: bookRepo}
			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_DeleteBook(t *testing.T) {

	type testCase struct {
		name           string
		expectedResult string
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return("Book with id: 1 deleted successfully", nil).Times(1)
		},
		expectedResult: "Book with id: 1 deleted successfully",
	})

	testTable = append(testTable, testCase{
		name:          "Not Found",
		wantError:     true,
		expectedError: gorm.ErrRecordNotFound,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return("", gorm.ErrRecordNotFound).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{repo: bookRepo}
			res, err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}


