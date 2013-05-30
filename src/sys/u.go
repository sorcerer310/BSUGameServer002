package sys

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func CheckError(w http.ResponseWriter,err error) {
	if err != nil{
		//fmt.Printf("Fatal error: %s\n",err.Error())
		fmt.Fprintf(w,"Error:"+err.Error())
		panic(err.Error())
	}
}
//解析一般对象为json字符串
func JsonParser(v interface{}) string{
	b,err := json.Marshal(v)
	if err != nil{
		panic(err.Error())
	}
	return string(b)
}
