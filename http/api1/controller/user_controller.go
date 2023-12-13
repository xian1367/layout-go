package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xian137/layout-go/database/model/scope"
	"github.com/xian137/layout-go/database/model/user"
	"github.com/xian137/layout-go/http/api1/request"
	"github.com/xian137/layout-go/pkg/gin/response"
	"github.com/xian137/layout-go/pkg/validator"
)

type UserController struct {
	BaseAPIController
}

// Index 列表
// @Summary 列表
// @Description 列表
// @Tags api1.user
// @Accept */*
// @Produce json
// @Param payload query scope.Query true "payload"
// @Success 200 {object} response.Paging{data=[]user.User,pager=scope.Paging}
// @Failure 500 {object} response.Failure
// @Router /user [get]
func (ctrl *UserController) Index(c *gin.Context) {
	var pager scope.Paging
	var data []user.User
	data, pager = user.Paginate(c)
	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// Show 详情
// @Summary 详情
// @Description 详情
// @Tags api1.user
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} user.User
// @Failure 500 {object} response.Failure
// @Router /user/{id} [get]
func (ctrl *UserController) Show(c *gin.Context) {
	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, userModel)
}

// Store 新增
// @Summary 新增
// @Description 新增
// @Tags api1.user
// @Accept json
// @Produce json
// @Param payload body request.UserStoreRequest true "payload"
// @Success 201 {object} user.User
// @Failure 500 {object} response.Failure
// @Router /user [post]
func (ctrl *UserController) Store(c *gin.Context) {
	req := request.UserStoreRequest{}
	if ok := validator.Validate(c, &req); !ok {
		return
	}

	userModel := user.User{}
	userModel.Create()
	if userModel.ID > 0 {
		response.Created(c, userModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// Update 更新
// @Summary 更新
// @Description 更新
// @Tags api1.user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body request.UserUpdateRequest true "payload"
// @Success 200 {object} user.User
// @Failure 500 {object} response.Failure
// @Router /user/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
	req := request.UserUpdateRequest{}
	if ok := validator.Validate(c, &req); !ok {
		return
	}

	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// Delete 删除
// @Summary 删除
// @Description 删除
// @Tags api1.user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 ""
// @Failure 500 {object} response.Failure
// @Router /user/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := userModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
