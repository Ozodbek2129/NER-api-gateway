package handler

import (
	"encoding/json"
	"gateway/genproto/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateGroup godoc
// @Summary      Create group
// @Description  Yangi guruh yaratish
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        body body services.CreateGroupReq true "Group ma'lumotlari"
// @Success      200 {object} services.CreateGroupRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/creategroup [post]
func (h Handler) CreateGroup(c *gin.Context) {
	var req services.CreateGroupReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.CreateGroup(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to CreateGroup function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateGroup godoc
// @Summary      Update group
// @Description  Guruhni yangilash
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        body body services.UpdateGroupReq true "Group yangilash ma'lumotlari"
// @Success      200 {object} services.UpdateGroupRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/updategroup [put]
func (h Handler) UpdateGroup(c *gin.Context) {
	var req services.UpdateGroupReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.UpdateGroup(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to UpdateGroup function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteGroup godoc
// @Summary      Delete group
// @Description  Guruhni o'chirish
// @Tags         group
// @Produce      json
// @Param        id path string true "Group ID"
// @Success      200 {object} services.DeleteGroupRes
// @Failure      500 {object} map[string]string
// @Router       /employee/deletegroup/{id} [delete]
func (h Handler) DeleteGroup(c *gin.Context) {
	req := services.DeleteGroupReq{
		Id: c.Param("id"),
	}

	res, err := h.EmployeeService.DeleteGroup(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to DeleteGroup function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllGroup godoc
// @Summary      Get all groups
// @Description  Barcha guruhlarni olish
// @Tags         group
// @Produce      json
// @Param        limit query int false "Limit" default(10)
// @Param        page  query int false "Page"  default(1)
// @Success      200 {object} services.GetAllGroupRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/getallgroup [get]
func (h Handler) GetAllGroup(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	req := services.GetAllGroupReq{
		Limit: int64(limit),
		Page:  int64(page),
	}

	res, err := h.EmployeeService.GetAllGroup(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetAllGroup function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateWorker godoc
// @Summary      Create worker
// @Description  Yangi ishchi yaratish
// @Tags         worker
// @Accept       json
// @Produce      json
// @Param        body body services.CreateWorkerReq true "Worker ma'lumotlari"
// @Success      200 {object} services.CreateWorkerRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/createworker [post]
func (h Handler) CreateWorker(c *gin.Context) {
	var req services.CreateWorkerReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.CreateWorker(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to CreateWorker function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateWorker godoc
// @Summary      Update worker
// @Description  Ishchini yangilash
// @Tags         worker
// @Accept       json
// @Produce      json
// @Param        body body services.UpdateWorkerReq true "Worker yangilash ma'lumotlari"
// @Success      200 {object} services.UpdateWorkerRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/updateworker [put]
func (h Handler) UpdateWorker(c *gin.Context) {
	var req services.UpdateWorkerReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.UpdateWorker(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to UpdateWorker function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteWorker godoc
// @Summary      Delete worker
// @Description  Ishchini o'chirish
// @Tags         worker
// @Produce      json
// @Param        id path string true "Worker ID"
// @Success      200 {object} services.DeleteWorkerRes
// @Failure      500 {object} map[string]string
// @Router       /employee/deleteworker/{id} [delete]
func (h Handler) DeleteWorker(c *gin.Context) {
	req := services.DeleteWorkerReq{
		Id: c.Param("id"),
	}

	res, err := h.EmployeeService.DeleteWorker(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to DeleteWorker function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllWorker godoc
// @Summary      Get all workers
// @Description  Barcha ishchilarni olish
// @Tags         worker
// @Produce      json
// @Param        limit    query int    false "Limit"    default(10)
// @Param        page     query int    false "Page"     default(1)
// @Param        group_id query string false "Group ID"
// @Success      200 {object} services.GetAllWorkerRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/getallworker [get]
func (h Handler) GetAllWorker(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")
	idStr := c.Query("group_id")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	req := services.GetAllWorkerReq{
		Limit:   int64(limit),
		Page:    int64(page),
		GroupId: idStr,
	}

	res, err := h.EmployeeService.GetAllWorker(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetAllWorker function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateAttendance godoc
// @Summary      Create attendance
// @Description  Davomat yaratish
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Param        body body services.Attendance true "Davomat ma'lumotlari"
// @Success      200 {object} services.CreateAttendanceRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/createattendance [post]
func (h Handler) CreateAttendance(c *gin.Context) {
	var req []*services.Attendance

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req1 := services.CreateAttendanceReq{
		Attendance: req,
	}

	res, err := h.EmployeeService.CreateAttendance(c, &req1)
	if err != nil {
		h.Log.Error("Error sending request to CreateAttendance function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateAttendance godoc
// @Summary      Update attendance
// @Description  Davomatni yangilash
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Param        body body services.AttendanceUpdate true "Davomat yangilash ma'lumotlari"
// @Success      200 {object} services.UpdateAttendanceRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/updateattendance [put]
func (h Handler) UpdateAttendance(c *gin.Context) {
	var req []*services.AttendanceUpdate

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req1 := services.UpdateAttendanceReq{
		Attendance: req,
	}

	res, err := h.EmployeeService.UpdateAttendance(c, &req1)
	if err != nil {
		h.Log.Error("Error sending request to UpdateAttendance function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteAttendance godoc
// @Summary      Delete attendance
// @Description  Davomatni o'chirish (sana oralig'i bo'yicha)
// @Tags         attendance
// @Produce      json
// @Param        todaydate  query string true "Bugungi sana (RFC3339: 2024-01-15T00:00:00Z)"
// @Param        deletedate query string true "O'chiriladigan sana (RFC3339: 2024-01-15T00:00:00Z)"
// @Success      200 {object} services.DeleteAttendanceRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/deleteattendance [delete]
func (h Handler) DeleteAttendance(c *gin.Context) {
	todayDateStr := c.Query("todaydate")
	deleteDateStr := c.Query("deletedate")

	todayDate, err := time.Parse(time.RFC3339, todayDateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid today_date format, use RFC3339 (e.g. 2024-01-15T00:00:00Z)",
		})
		return
	}

	deleteDate, err := time.Parse(time.RFC3339, deleteDateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid delete_date format, use RFC3339 (e.g. 2024-01-15T00:00:00Z)",
		})
		return
	}

	req := services.DeleteAttendanceReq{
		TodayDate:  timestamppb.New(todayDate),
		DeleteDate: timestamppb.New(deleteDate),
	}

	res, err := h.EmployeeService.DeleteAttendance(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to DeleteAttendance function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetDailyAttendance godoc
// @Summary      Get daily attendance
// @Description  Kunlik davomatni olish
// @Tags         attendance
// @Produce      json
// @Param        work_date query string true "Ish sanasi (RFC3339: 2024-01-15T00:00:00Z)"
// @Success      200 {object} services.GetDailyAttendanceRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/getdailyattendance [get]
func (h Handler) GetDailyAttendance(c *gin.Context) {
	workDateStr := c.Query("work_date")

	workdate, err := time.Parse(time.RFC3339, workDateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid work_date format, use RFC3339 (e.g. 2024-01-15T00:00:00Z)",
		})
		return
	}

	req := services.GetDailyAttendanceReq{
		WorkDate: timestamppb.New(workdate),
	}

	res, err := h.EmployeeService.GetDailyAttendance(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetDailyAttendance function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllAttendance godoc
// @Summary      Get all attendance
// @Description  Barcha davomatni sana oralig'i bo'yicha olish
// @Tags         attendance
// @Produce      json
// @Param        startdate query string true "Boshlanish sanasi (RFC3339: 2024-01-15T00:00:00Z)"
// @Param        enddate   query string true "Tugash sanasi (RFC3339: 2024-01-15T00:00:00Z)"
// @Success      200 {object} services.GetAllAttendanceRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/getallattendance [get]
func (h Handler) GetAllAttendance(c *gin.Context) {
	startDateStr := c.Query("startdate")
	endDateStr := c.Query("enddate")

	startdate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid startdate format, use RFC3339 (e.g. 2024-01-15T00:00:00Z)",
		})
		return
	}

	enddate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid enddate format, use RFC3339 (e.g. 2024-01-15T00:00:00Z)",
		})
		return
	}

	req := services.GetAllAttendanceReq{
		StartDate: timestamppb.New(startdate),
		EndDate:   timestamppb.New(enddate),
	}

	res, err := h.EmployeeService.GetAllAttendance(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetAllAttendance function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateTask godoc
// @Summary      Create task
// @Description  Yangi vazifa yaratish
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        body body services.CreateTaskReq true "Task ma'lumotlari"
// @Success      200 {object} services.CreateTaskRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/createtask [post]
func (h Handler) CreateTask(c *gin.Context) {
	var req services.CreateTaskReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.CreateTask(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to CreateTask function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateTask godoc
// @Summary      Update task
// @Description  Vazifani yangilash
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        body body services.UpdateTaskReq true "Task yangilash ma'lumotlari"
// @Success      200 {object} services.UpdateTaskRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/updatetask [put]
func (h Handler) UpdateTask(c *gin.Context) {
	var req services.UpdateTaskReq

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.EmployeeService.UpdateTask(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to UpdateTask function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTask godoc
// @Summary      Delete task
// @Description  Vazifani o'chirish
// @Tags         task
// @Produce      json
// @Param        id path string true "Task ID"
// @Success      200 {object} services.DeleteTaskRes
// @Failure      500 {object} map[string]string
// @Router       /employee/deletetask/{id} [delete]
func (h Handler) DeleteTask(c *gin.Context) {
	req := services.DeleteTaskReq{
		Id: c.Param("id"),
	}

	res, err := h.EmployeeService.DeleteTask(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to DeleteTask function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllTask godoc
// @Summary      Get all tasks
// @Description  Barcha vazifalarni olish
// @Tags         task
// @Produce      json
// @Param        limit     query int    false "Limit"    default(10)
// @Param        page      query int    false "Page"     default(1)
// @Param        worker_id query string false "Worker ID"
// @Success      200 {object} services.GetAllTaskRes
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /employee/getalltask [get]
func (h Handler) GetAllTask(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")
	idStr := c.Query("worker_id")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	req := services.GetAllTaskReq{
		Limit:    int64(limit),
		Page:     int64(page),
		WorkerId: idStr,
	}

	res, err := h.EmployeeService.GetAllTask(c, &req)
	if err != nil {
		h.Log.Error("Error sending request to GetAllTask function.", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}