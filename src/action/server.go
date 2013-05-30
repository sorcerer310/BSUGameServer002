package action

import (
//	"sys"
	"fmt"
	"net/http"
//	"db"
//	"github.com/gorilla/session"
)

func GetServerList(w http.ResponseWriter,r *http.Request){
	session,_:=store.Get(r,"sessionid")
	fmt.Printf("%s",session.ID)
}

