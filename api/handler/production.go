package handler

import (
	"context"
	"encoding/json"
	"fmt"
	pbp "gateway/genproto/ishlab_chiqarish"
	"gateway/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// @Summary Create new contract
// @Description Create a new contract with file upload
// @Tags contract
// @Accept multipart/form-data
// @Produce json
// @Param contract_name formData string true "Contract Name"
// @Param contract_number formData string true "Contract Number"
// @Param contract_deadline formData string true "Deadline"
// @Param responsible_person formData string true "Responsible Person"
// @Param file formData file true "Contract File"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/contract/newcontract [post]
func (h Handler) NewContract(c *gin.Context) {
	var req models.CreateContract

	// 🔹 form-data bind
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 🔹 file temporary save
	filePath := "./media/" + req.File.Filename
	if err := c.SaveUploadedFile(req.File, filePath); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 🔹 MinIO client
	minioClient, err := minio.New("192.168.0.44:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	bucketName := "contracts"

	// 🔹 bucket check
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 🔹 agar yo‘q bo‘lsa yaratamiz
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// 🔥 KEYIN policy beriladi (to‘g‘ri joy)
	policy := `{
		"Version":"2012-10-17",
		"Statement":[
			{
				"Effect":"Allow",
				"Principal":"*",
				"Action":["s3:GetObject"],
				"Resource":["arn:aws:s3:::` + bucketName + `/*"]
			}
		]
	}`

	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 🔹 unique file name
	fileExt := filepath.Ext(req.File.Filename)
	newFileName := uuid.NewString() + fileExt

	// 🔹 upload to MinIO
	_, err = minioClient.FPutObject(
		context.Background(),
		bucketName,
		newFileName,
		filePath,
		minio.PutObjectOptions{
			ContentType: "application/pdf", // 🔥 tuzatildi
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 🔹 file URL
	fileURL := fmt.Sprintf("http://192.168.0.44:9000/%s/%s", bucketName, newFileName)

	// 🔹 service request
	serviceReq := pbp.NewContractReq{
		ContractName:      req.ContractName,
		ContractNumber:    req.ContractNumber,
		ContractDeadline:  req.ContractDeadline,
		ResponsiblePerson: req.ResponsiblePerson,
		ContractFileUrl:   fileURL,
	}

	res, err := h.ProductionService.NewContract(c, &serviceReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "contract created",
		"data":    res,
		"fileUrl": fileURL,
	})
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
		Page:  int64(page),
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
		Page:  int32(page),
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
