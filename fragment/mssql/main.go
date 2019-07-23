package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	_test = "SELECT * FROM sysobjects WHERE name = @p1"
)

func main() {

	db, err := sql.Open("sqlserver", "sqlserver://sa:password@host:port?database=database&encrypt=disable")
	if err!=nil{
		fmt.Println(err.Error())
	}
	err = db.Ping()
	if err!=nil{
		fmt.Println(err.Error())
	}

	rows,err:=db.Query(_test,"name")
	m:=""
	n:=""
	if err!=nil{
		fmt.Println(err.Error())
	}


	s,_:=rows.Columns()
	fmt.Println(s,)

	if rows.Next(){
		err=rows.Scan(&m,&n)
		if err!=nil{
			fmt.Println(err)
		}
	}

	fmt.Println(m,n)
}
