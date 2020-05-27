package schemas

// prometheus alertManager webHook data
type AlertManagerWebHook struct {
	Alerts []AlertManagerObj `json:"alerts"`
}

// alert实例
type AlertManagerObj struct {
	Status      string                 `json:"status"`
	Labels      AlertManagerLabel      `json:"labels"`
	Annotations AlertManagerAnnotation `json:"annotations"`
	StartAt     string                 `json:"startsAt"`
	EndAt       string                 `json:"endsAt"`
}

// alert标签
type AlertManagerLabel struct {
	Instance  string `json:"instance"`
	AlertName string `json:"alertname"`
	Level     string `json:"level"`
}

// alert信息
type AlertManagerAnnotation struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
