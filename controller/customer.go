package controller

import (
	"github.com/gin-gonic/gin"
	"kreditplus-test/dto"
	"kreditplus-test/helper"
	"kreditplus-test/model"
	"kreditplus-test/service"
	"net/http"
)

type CustomerController struct {
	customerService *service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) *CustomerController {
	return &CustomerController{customerService}
}

func (cu CustomerController) Login(c *gin.Context) {
	var req dto.LoginCustomerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), "", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}

	token, err := cu.customerService.Login(req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), "", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "login", c.FullPath(), req.NIK, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success login to your account", gin.H{
		"token": token,
	}))
	return
}

func (cu CustomerController) Register(c *gin.Context) {
	var req dto.RegisterCustomerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), req.NIK, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}

	err = cu.customerService.Register(req)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), req.NIK, http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "register", c.FullPath(), req.NIK, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success register your account", nil))
	return
}

func (cu CustomerController) UploadKYC(c *gin.Context) {
	id := c.MustGet("user").(model.Customer).ID

	mform, err := c.MultipartForm()
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}
	err = cu.customerService.KYC(id, *mform)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "kyc", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success upload kyc", nil))
	return
}

func (cu CustomerController) GetCurrentUser(c *gin.Context) {
	id := c.MustGet("user").(model.Customer).ID

	customer, err := cu.customerService.GetByID(id)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusInternalServerError)
		return
	}

	helper.Log(false, "get current user", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success get current user", customer))
	return
}
