package handler

import (
	"simpleapi/helper"
	"simpleapi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	book := model.Book{}

	err := c.BindJSON(&book)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateBook(book)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)

}

func (h HttpServer) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.GetBookByID(id)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.Ok(c, res)

}

func (h HttpServer) GetAllBooks(c *gin.Context) {
	res, err := h.app.GetAllBooks()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)

}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.DeleteBook(id)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.Ok(c, res)

}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	book := model.Book{}

	err = c.BindJSON(&book)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateBook(book, id)
	if err != nil {
		if err == helper.NotFoundErr {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)

}
