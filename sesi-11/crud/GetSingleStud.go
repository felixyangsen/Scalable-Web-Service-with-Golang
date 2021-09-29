package crud

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestSingle struct {
	Param string
}

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Grade string `json:"grade"`
}

func GetSingleStud(w http.ResponseWriter, r *http.Request, connMysq *sql.DB) ResponseData {
	c, errRead := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var response ResponseData
	if errRead != nil {
		log.Println(errRead)
		response.Status = 500
		response.Data = errRead.Error()
		return response
	}
	var msg RequestSingle
	errUnmarshal := json.Unmarshal(c, &msg)
	if errUnmarshal != nil {
		log.Println(errUnmarshal)
		response.Status = 502
		response.Data = "Unmarshall failed"
		return response
	}
	// connMysq,errConn := ConnectDb("mysql")
	// defer connMysq.Close()
	// if errConn != nil {
	//  log.Println(errConn)
	//  response.Status = 502
	//  response.Data = "Connection to database failed"
	// }
	var result Student
	id := msg.Param
	err := connMysq.QueryRow("select id,name,age,grade from tb_student where id = ?", id).Scan(&result.ID, &result.Name, &result.Age, &result.Grade)
	if err != nil {
		log.Println(err)
		response.Status = 500
		response.Data = err.Error()
		return response
	}
	response.Status = 200
	return response
}
