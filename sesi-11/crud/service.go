package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseData struct {
	Status int
	Data   string
	Error  error
}

type RequestModify struct {
	Method string
	Data   struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}
}

func ModifyData(w http.ResponseWriter, r *http.Request, connMysq *sql.DB) ResponseData {
	//1
	fmt.Println("request", r)
	fmt.Println("body", r.Body)

	c, errRead := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var response ResponseData
	if errRead != nil {
		log.Println(errRead)
		response.Status = 500
		response.Data = errRead.Error()
		return response
	}

	//2
	var msg RequestModify
	errUnmarshal := json.Unmarshal(c, &msg)
	if errUnmarshal != nil {
		log.Println(errUnmarshal)
		response.Status = 502
		response.Data = "Unmarshall failed"
		return response
	}

	//3
	// connMysq,errConn := ConnectDb("mysql")
	// if errConn != nil {
	//  log.Println(errConn)
	//  response.Status = 502
	//  response.Data = "Connection to database failed"
	//  return response
	// }

	var errExec error
	if msg.Method == "insert" {
		_, errExec = connMysq.Exec("insert into tb_student value(?,?,?,?)", msg.Data.ID, msg.Data.Name, msg.Data.Age, msg.Data.Grade)
	} else if msg.Method == "update" {
		_, errExec = connMysq.Exec("update tb_student set name= ?, age = ?, grade = ? where id = ?", msg.Data.Name, msg.Data.Age, msg.Data.Grade, msg.Data.ID)
	} else if msg.Method == "delete" {
		_, errExec = connMysq.Exec("delete from tb_student where id = ?", msg.Data.ID)
	} else {
		log.Println("Error, method info not available")
		response.Status = 500
		// response.Data = "error when running modify data"
		response.Error = errExec
		return response
	}
	if errExec != nil {
		log.Println(errExec)
		response.Status = 502
		// response.Data= "error when running modify statement"
		response.Error = errExec
	}

	response.Status = 200
	response.Data = "sukses " + msg.Method
	return response
}
