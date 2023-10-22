package api

import (
	"net/http"
	"time"

	"github.com/cdvelop/model"
)

type auth struct{}

func (auth) GetUser(r *http.Request) *model.User {

	user := model.User{
		Token:          "123",
		Id:             "123456789101112",
		Ip:             "172.0.0.1",
		Name:           "don Juanito dev test",
		Area:           "s",
		AccessLevel:    "1",
		LastConnection: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &user
}
