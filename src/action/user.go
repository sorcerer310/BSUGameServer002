package action

import (
	"sys"
	"fmt"
	"net/http"
	"db"
	"github.com/gorilla/sessions"
)

type User struct{
	Id int 
	No string 
	Nickname string 
	Pwd string 
	Date string	 
	Question string 
	Answer string 
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func LoginAction(w http.ResponseWriter, r *http.Request){
	no := r.FormValue("no")
	pwd := r.FormValue("pwd")
	
	user := new (User)
	stmt := db.BsuStmt("select * from gamedb.user where no=? and pwd=?").QueryRow(no,pwd)
	err := stmt.Scan(&user.Id,&user.No,&user.Nickname,&user.Pwd,&user.Date,&user.Question,&user.Answer)
	if err!=nil{
		//此处不能用panic,否则有字段为nil时会中断程序
		fmt.Printf(err.Error()+"\n")
//		panic(err.Error())
	}

	session,_:=store.Get(r,"sessionid")
	session.Values["userid"] = user.Id
	session.Values["no"] = user.No
	session.Values["nickname"] = user.Nickname
	session.Save(r,w)

	fmt.Fprintf(w,sys.JsonParser(user))
	
}
