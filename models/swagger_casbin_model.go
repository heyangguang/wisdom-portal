package models

type SwaggerRole struct {
	RoleName    string            `json:"role_name" binding:"required"`
	Remark      string            `json:"remark"`
	RoleObjActs []SwaggerObjActId `json:"roleObjActs" binding:"required"`
}

type SwaggerObjActId struct {
	ID string `json:"id"`
}
