package controller

import (
	"github.com/gin-gonic/gin"
	"kreditplus-test/helper"
	"kreditplus-test/model"
	"kreditplus-test/service"
	"net/http"
	"strconv"
)

type LimitController struct {
	limitService *service.LimitService
}

func NewLimitController(limitService *service.LimitService) *LimitController {
	return &LimitController{limitService}
}

func (l LimitController) GetAllLimitByUserID(c *gin.Context) {
	id := c.MustGet("user").(model.Customer).ID

	page := c.Query("page")
	limit := c.Query("limit")
	if page == "" || limit == "" {
		helper.Log(true, "page and limit required", c.FullPath(), id, http.StatusBadRequest)
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

	limitByUserID, err := l.limitService.GetAllLimitByUserID(id, pageInt, limitInt)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "get all limit", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success get limits", limitByUserID))
	return
}

func (l LimitController) GetLimitByTenor(c *gin.Context) {
	id := c.MustGet("user").(model.Customer).ID
	tenor := c.Param("tenor")

	//Convert To Int
	tenorInt, err := strconv.Atoi(tenor)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, "page or limit is required", nil))
		return
	}
	limitByTenor, err := l.limitService.GetLimitByTenor(id, tenorInt)
	if err != nil {
		helper.Log(true, err.Error(), c.FullPath(), id, http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), nil))
		return
	}

	helper.Log(false, "get tenor", c.FullPath(), id, http.StatusOK)
	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success get limit by tenor", limitByTenor))
	return
}
