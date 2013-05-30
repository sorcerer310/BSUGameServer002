package main 

import (
	"fmt"
	"net/http"
	"io"
	"action"
)

func main() {
	fmt.Printf("BSUGameServer001 is running\n")
	http.HandleFunc("/",defaultHandler)
	http.HandleFunc("/login",action.LoginAction)
	http.HandleFunc("/getserverlist",action.GetServerList)
	err := http.ListenAndServe(":80",nil)
	if err!=nil{
		fmt.Printf(err.Error())
	}
}

func defaultHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello bsugame!")
}