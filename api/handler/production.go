package handler

import (
	"encoding/json"
	pbp "gateway/genproto/ishlab_chiqarish"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) NewContract(c *gin.Context) {
	req := pbp.NewContractReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ProductionService.NewContract(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewContract function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewContractUpdate(c *gin.Context) {
	req := pbp.NewContractUpdateReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ProductionService.NewContractUpdate(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewContractUpdate function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewContractDelete(c *gin.Context) {
	req := pbp.NewContractDeleteReq{
		Id: c.Param("id"),
	}

	res, err := h.ProductionService.NewContractDelete(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewContractDelete function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewContractGetName(c *gin.Context) {
	req := pbp.NewContractGetNameReq{
		Name: c.Param("name"),
	}

	res, err := h.ProductionService.NewContractGetName(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewContractGetName function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewContractGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid limit",
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid page",
		})
		return
	}

	req := pbp.NewContractGetAllReq{
		Limit: int64(limit),
		Page: int64(page),
	}

	res, err := h.ProductionService.NewContractGetAll(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewContractGetAll function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewInsideTheContract(c *gin.Context) {
	req := pbp.NewInsideTheContractReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ProductionService.NewInsideTheContract(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewInsideTheContract function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewInsideTheContractUpdate(c *gin.Context) {
	req := pbp.NewInsideTheContractUpdateReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ProductionService.NewInsideTheContractUpdate(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewInsideTheContractUpdate function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewInsideTheContractDelete(c *gin.Context) {
	req := pbp.NewInsideTheContractDeleteReq{
		Id: c.Param("id"),
	}

	res, err := h.ProductionService.NewInsideTheContractDelete(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewInsideTheContractDelete function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) NewInsideTheContractGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid limit",
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid page",
		})
		return
	}

	req := pbp.NewInsideTheContractGetAllReq{
		Limit: int32(limit),
		Page: int32(page),
	}

	res, err := h.ProductionService.NewInsideTheContractGetAll(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to NewInsideTheContractGetAll function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}