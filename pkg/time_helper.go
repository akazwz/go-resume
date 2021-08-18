package pkg

import (
	"log"
	"time"
)

func StrToTime(str string, layout string) time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("日期转换失败")
		return time.Time{}
	}
	inLocation, err := time.ParseInLocation(layout, str, location)
	if err != nil {
		log.Println("日期转换失败")
		return time.Time{}
	}
	return inLocation
}
