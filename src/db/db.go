package db

import (
	"database/sql"
	"database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
)
//直接返回db对象,连接池以后再研究
func DBFactory()(*sql.DB,error){
	db,e := sql.Open("mysql","root:root@/mysql?charset=utf8")
	return db,e
}

//直接返回Stmt用于代入参数 查询 执行 sql
func BsuStmt(sql string)(*sql.Stmt){
		conn,err := DBFactory()
		if err != nil{
			panic(err.Error())
		}
		//defer conn.Close()
		
		stmtOut,err := conn.Prepare(sql)
		if err!=nil{
			panic(err.Error())
		}
		//defer stmtOut.Close()
		return stmtOut
}

//一般查询
func Query(sql string,args []driver.Value)(*sql.Rows, error){
	conn,err := DBFactory()
	if err != nil{
		panic(err.Error())
	}
	defer conn.Close()
	
	stmtOut,err := conn.Prepare(sql)
	if err!=nil{
		panic(err.Error())
	}
	defer stmtOut.Close()
	
	params := []driver.Value{args[0],args[1]}

	rows,err := stmtOut.Query(params)
	//rows,err := stmtOut.Query(args[0],args[1])
	if err!=nil{
		panic(err.Error())
	}
	
	return rows,err
}


//连接池,以后整理
//var MySQLPool chan *mysql.MySQL  
//func getMySQL() *mysql.MySQL {  
//    if MySQLPool == nil {  
//        MySQLPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)  
//    }  
//    if len(MySQLPool) == 0 {  
//        go func() {  
//            for i := 0; i < MAX_POOL_SIZE/2; i++ {  
//                mysql := mysql.New()  
//                err := mysql.Connect("127.0.0.1", "root", "", "wgt", 3306)  
//                if err != nil {  
//                    panic(err.String())  
//                }     
//                putMySQL(mysql)  
//            }     
//        }()   
//    }     
//    return <-MySQLPool  
//}  
//func putMySQL(conn *mysql.MySQL) {  
//    if MySQLPool == nil {  
//        MySQLPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)  
//    }     
//    if len(MySQLPool) == MAX_POOL_SIZE {  
//        conn.Close()  
//        return  
//    }  
//    MySQLPool <- conn  
//}  
