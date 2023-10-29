package response

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

type JsonTime time.Time

func (j JsonTime) Marsha1JSON() ([]byte, error) {
	zap.S().Info("jsontime", j)
	stmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-01"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	Nickname string `json:"name"`
	//Birthday JsonTime `json:"birthday"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
}
