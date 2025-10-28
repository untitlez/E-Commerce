package config

import (
	"fmt"
	"time"
)

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Time Zone Error", err.Error())
	}

	time.Local = ict
}
