package model

import (
  "time"

  "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
  ID          int    `json:"book_id" gorm:"primaryKey"`
  Title       string `json:"book_title" gorm:"not null;unique;type:varchar(50)"`
	Author      string `json:"book_author" gorm:"not null;type:varchar(50)"`
	Description string `json:"description" gorm:"not null;type:varchar(50)"`
  CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:false"`
  UpdatedAt time.Time `json:"updated_at"`
}

func (a Book) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Title, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Author, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.Description, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of five digits
	)
}
