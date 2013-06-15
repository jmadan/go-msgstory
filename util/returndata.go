package util

type ReturnData struct {
	Success  bool   `json:"success"`
	JsonData string `json:"json_data"`
}

func (return_data *ReturnData) GetSuccess() bool {
	return return_data.Success
}

func (return_data *ReturnData) GetJsonData() string {
	return return_data.JsonData
}
