package biz

import (
	"encoding/json"
	"log"

	"../service"
)

func GetDevice() (string, error) {
	device, err := service.QueryDevice()
	if err == nil {
		return "", nil
	} else {
		dataRet, err := json.Marshal(map[string]interface{}{"device": device})
		if err == nil {
			return string(dataRet), nil
		} else {
			log.Println(err)
			return "", nil
		}

	}

}
