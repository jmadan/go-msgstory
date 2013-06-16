package util

type ReturnData struct {
	Success  bool   `json:"success"`
	ErrorMsg error  `json:"error_message"`
	JsonData []byte `json:"json_data"`
	Status   string `json:"status"`
}

func (return_data *ReturnData) GetSuccess() bool {
	return return_data.Success
}

func (return_data *ReturnData) GetJsonData() []byte {
	return return_data.JsonData
}

func (return_data *ReturnData) GetErrorMessage() error {
	return return_data.ErrorMsg
}

func (return_data *ReturnData) GetStatus() string {
	return return_data.Status
}
