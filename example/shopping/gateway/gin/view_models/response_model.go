package viewmodels

//ResultInfo 请求应答json结构体
type ResultInfo struct {
	//状态码
	Code int `json:"code"`
	//结果
	Data interface{} `json:"result"`

}