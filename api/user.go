package api

import (
	"github.com/iswangwenbin/golion/service"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	user := service.GetOneUser()
	w.Write([]byte(user.ToJson()))
}
