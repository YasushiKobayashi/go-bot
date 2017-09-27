package handler

import "log"

type (
	Res struct {
		Res string `json:"res"`
		Err string `json:"err"`
	}
)

func HandleErrorRes(param string, err error) (res Res) {
	log.Printf("data : %v", err)
	res = Res{
		Res: param,
		Err: err.Error(),
	}
	return res
}
