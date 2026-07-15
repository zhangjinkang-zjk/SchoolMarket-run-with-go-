package userhandler

import (
	userservice "SchoolMarket-run-with-go-/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	serv userservice.UserService
}

func NewUserHandler(serv userservice.UserService) *UserHandler {
	return &UserHandler{serv: serv}
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Psw  string `json:"psw"`
}

type DeleteUserRequest struct {
	Id uint `json:"id"`
}

type UpdateUserRequest struct {
	Id      uint    `json:"id"`
	NewName *string `json:"newname"`
	NewPsw  *string `json:"newpsw"`
}

type GetAimUserRequest struct {
	Id uint `json:"id"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	user, err := h.serv.CreateUser(userservice.CreateUserRequest{
		Name: req.Name,
		Psw:  req.Psw,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := h.serv.Delete(userservice.DeleteUserRequest{
		Id: req.Id,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"status": "done",
		},
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	user, err := h.serv.UpdateUser(userservice.UpdateUserRequest{
		Id:      req.Id,
		NewName: req.NewName,
		NewPsw:  req.NewPsw,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	var req GetAimUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	user, err := h.serv.GetUserById(userservice.GetAimUserRequest{
		Id: req.Id,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.serv.GetAllUser()
	data := make([]gin.H, 0, len(users))
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	for _, user := range users {
		data = append(data, gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
