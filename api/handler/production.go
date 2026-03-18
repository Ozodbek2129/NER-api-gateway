package handler

import (
	"encoding/json"
	pbp "gateway/genproto/ishlab_chiqarish"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create new contract
// @Description Create a new contract
// @Tags contract
// @Accept json
// @Produce json
// @Param body body ishlab_chiqarish.NewContractReq true "New Contract"
// @Success 200 {object} ishlab_chiqarish.NewContractRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/newcontract [post]
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

// @Summary Update contract
// @Description Update existing contract
// @Tags contract
// @Accept json
// @Produce json
// @Param body body ishlab_chiqarish.NewContractUpdateReq true "Update Contract"
// @Success 200 {object} ishlab_chiqarish.NewContractUpdateRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/contract_update [put]
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

// @Summary Delete contract
// @Description Delete contract by ID
// @Tags contract
// @Produce json
// @Param id path string true "Contract ID"
// @Success 200 {object} ishlab_chiqarish.NewContractDeleteRes
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/contract_delete/{id} [delete]
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

// @Summary Get contract by name
// @Description Get contract using name
// @Tags contract
// @Produce json
// @Param name path string true "Contract Name"
// @Success 200 {object} ishlab_chiqarish.NewContractGetNameRes
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/get_name/{name} [get]
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

// @Summary Get all contracts
// @Description Get all contracts with pagination
// @Tags contract
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} ishlab_chiqarish.NewContractGetAllRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/all_contract [get]
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

// @Summary Create inside contract
// @Description Create inside contract
// @Tags inside_contract
// @Accept json
// @Produce json
// @Param body body ishlab_chiqarish.NewInsideTheContractReq true "Inside Contract"
// @Success 200 {object} ishlab_chiqarish.NewInsideTheContractRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/inside_contract [post]
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

// @Summary Update inside contract
// @Description Update inside contract
// @Tags inside_contract
// @Accept json
// @Produce json
// @Param body body ishlab_chiqarish.NewInsideTheContractUpdateReq true "Update Inside Contract"
// @Success 200 {object} ishlab_chiqarish.NewInsideTheContractUpdateRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/insidecontract_update [put]
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

// @Summary Delete inside contract
// @Description Delete inside contract by ID
// @Tags inside_contract
// @Produce json
// @Param id path string true "Inside Contract ID"
// @Success 200 {object} ishlab_chiqarish.NewInsideTheContractDeleteRes
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/insidecontract_delete/{id} [delete]
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

// @Summary Get all inside contracts
// @Description Get all inside contracts with pagination
// @Tags inside_contract
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} ishlab_chiqarish.NewInsideTheContractGetAllRes
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/all_insidecontract [get]
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