package models

type SwaggerUserGroup struct {
	GroupName string          `binding:"required" json:"group_name"`
	Remark    string          `json:"remark"`
	Users     []SwaggerUserId `binding:"required" json:"users"`
}

type SwaggerUserId struct {
	ID string `json:"id"`
}
