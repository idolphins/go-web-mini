package vo

// 用户登录结构体
type RegisterAndLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 创建用户结构体
type CreateUserRequest struct {
	Username     string `form:"username" json:"username" validate:"required"`
	Password     string `form:"password" json:"password"`
	Mobile       string `form:"mobile" json:"mobile" validate:"required"`
	Avatar       string `form:"avatar" json:"avatar"`
	Nickname     string `form:"nickname" json:"nickname"`
	Introduction string `form:"introduction" json:"introduction"`
	Status       uint   `form:"status" json:"status"`
	RoleIds      []uint `form:"roleIds" json:"roleIds" validate:"required"`
}

// 获取用户列表结构体
type UserListRequest struct {
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Nickname string `json:"nickname" form:"nickname"`
	Status   uint   `json:"status" form:"status"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 批量删除用户结构体
type DeleteUserRequest struct {
	UserIds []uint `json:"userIds" form:"userIds"`
}

// 更新密码结构体
type ChangePwdRequest struct {
	OldPassword string `json:"oldPassword" form:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required"`
}
