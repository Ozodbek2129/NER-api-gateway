package handler

import (
	"context"
	"encoding/json"
	"fmt"
	pbu "gateway/genproto/user"
	"gateway/models"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (h Handler) GetUSerByEmail(c *gin.Context) {
	req := pbu.GetUSerByEmailReq{
		Email: c.Param("email"),
	}

	res, err := h.UserService.GetUSerByEmail(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetUserByEmail function.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) UpdatePassword(c *gin.Context) {
	req := pbu.UpdatePasswordReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.UserService.UpdatePassword(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to password update function", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) DeleteUser(c *gin.Context) {
	req := pbu.UserId{
		Id: c.Param("id"),
	}

	res, err := h.UserService.DeleteUser(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to Delete user function", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) UpdateRole(c *gin.Context) {
	req := pbu.UpdateRoleReq{}

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.UserService.UpdateRole(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to UpdateRole function", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h Handler) ProfileImage(c *gin.Context) {
	var file models.File
	email := c.Param("email")

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileUrl := filepath.Join("./media", file.File.Filename)

	err = c.SaveUploadedFile(&file.File, fileUrl)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileExt := filepath.Ext(file.File.Filename)

	newFile := uuid.NewString() + fileExt

	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bucketName := "photos"

	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	_, err = minioClient.FPutObject(context.Background(), bucketName, newFile, fileUrl, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	policy := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "AWS": ["*"]
                },
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::%s/*"]
            }
        ]
    }`, bucketName)

	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	objUrl, err := minioClient.PresignedGetObject(context.Background(), bucketName, newFile, time.Hour*24, nil)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	req := pbu.ImageReq{Image: objUrl.String(), Email: email}

	_, err = h.UserService.ProfileImage(c, &req)
	if err != nil {
		h.Log.Error("Email buyicha malumotlarni olishda xatolik", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	maduUrl := fmt.Sprintf("http://localhost:9000/photos/%s", newFile)

	c.JSON(201, gin.H{
		"obj":     objUrl.String(),
		"madeUrl": maduUrl,
	})
}

func (h Handler) GetAllUsers(c *gin.Context) {
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

	req := pbu.GetAllUsersReq{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.UserService.GetallUsers(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetallUsers function", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h *Handler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization header is required"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is empty"})
		return
	}

	// Token muddati — configdan olish tavsiya etiladi
	const tokenExpiry = 15 * time.Minute

	if err := h.Redis.BlacklistToken(tokenStr, tokenExpiry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "logout failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}
