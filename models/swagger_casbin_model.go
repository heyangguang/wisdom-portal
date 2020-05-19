package models

type SwaggerRule struct {
	RuleName    string            `json:"rule_name" binding:"required"`
	Remark      string            `json:"remark"`
	RuleObjActs []SwaggerObjActId `json:"ruleObjActs" binding:"required"`
}

type SwaggerObjActId struct {
	ID string `json:"id"`
}
