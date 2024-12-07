package controller

import (
	"github.com/gin-gonic/gin"
	"kreditplus-test/dto"
	"kreditplus-test/helper"
	"kreditplus-test/model"
	"kreditplus-test/service"
	"net/http"
	"strconv"
)

type TransaksiController struct {
	transaksiService *service.TransaksiService
}

func NewTransaksiController(transaksiService *service.TransaksiService) *TransaksiController {
	return &TransaksiController{transaksiService}
}

func (t TransaksiController) CreateTransaction(c *gin.Context) {
	var req dto.TransaksiRequestDTO
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), "", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}

	id := c.MustGet("user").(model.Customer).ID

	transaksi, err := t.transaksiService.CreateTransaksi(id, req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), "", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "create transaction", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success create transaction", transaksi))
	return
}

func (t TransaksiController) GetAllTransactionCurrentUser(c *gin.Context) {
	id := c.MustGet("user").(model.Customer).ID

	page := c.Query("page")
	limit := c.Query("limit")
	if page == "" || limit == "" {
		helper.Log(true, "page and limit required", c.FullPath(), "", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, "page or limit is required", nil))
		return
	}

	//Convert To Int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, "page or limit is required", nil))
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, "page or limit is required", nil))
		return
	}

	getAllTransaction, err := t.transaksiService.GetAllTransaction(id, pageInt, limitInt)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "get all transaction", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success get all transaction", getAllTransaction))
	return
}
