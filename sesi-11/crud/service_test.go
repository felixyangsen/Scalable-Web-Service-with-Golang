package crud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_modifyData_insert(t *testing.T) {
	t.Run("Insert success", func(t *testing.T) {
		bodyData := new(RequestModify)

		bodyData.Method = "insert"
		bodyData.Data.ID = "W009"
		bodyData.Data.Name = "Daniel"
		bodyData.Data.Age = 2
		bodyData.Data.Grade = 3

		marshalBody, _ := json.Marshal(bodyData)

		mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(marshalBody))

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()

		mock.ExpectExec("insert into tb_student").WithArgs(bodyData.Data.ID, bodyData.Data.Name, bodyData.Data.Age, bodyData.Data.Grade).WillReturnResult(sqlmock.NewResult(1, 1))

		expectedResult := ResponseData{Status: 200, Data: "sukses insert"}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, expectedResult, actualResp)
	})
}

func Test_modifyData_update(t *testing.T) {
	t.Run("Update success", func(t *testing.T) {
		bodyData := new(RequestModify)

		bodyData.Method = "update"
		bodyData.Data.ID = "W009"
		bodyData.Data.Name = "Daniel"
		bodyData.Data.Age = 2
		bodyData.Data.Grade = 3

		marshalBody, _ := json.Marshal(bodyData)

		mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(marshalBody))

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()

		mock.ExpectExec("update tb_student").WillReturnResult(sqlmock.NewResult(1, 1))

		expectedResult := ResponseData{Status: 200, Data: "sukses update"}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, expectedResult, actualResp)
	})
}

func Test_modifyData_delete(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		bodyData := new(RequestModify)

		bodyData.Method = "delete"
		bodyData.Data.ID = "W009"
		bodyData.Data.Name = "Daniel"
		bodyData.Data.Age = 2
		bodyData.Data.Grade = 3

		marshalBody, _ := json.Marshal(bodyData)

		mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(marshalBody))

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		defer db.Close()

		mock.ExpectExec("delete from tb_student").WillReturnResult(sqlmock.NewResult(1, 1))

		expectedResult := ResponseData{Status: 200, Data: "sukses delete"}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, expectedResult, actualResp)
	})
}

func Test_modifyData_wrong_method(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		bodyData := new(RequestModify)

		bodyData.Method = ""
		bodyData.Data.ID = "W009"
		bodyData.Data.Name = "Daniel"
		bodyData.Data.Age = 2
		bodyData.Data.Grade = 3

		marshalBody, _ := json.Marshal(bodyData)

		mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(marshalBody))

		expectedResult := ResponseData{Status: 500}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, nil)

		assert.Equal(t, expectedResult, actualResp)
	})
}

func Test_modifyData_unmarshal_error(t *testing.T) {
	t.Run("read failed", func(t *testing.T) {
		mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(nil))

		expectedResult := ResponseData{Status: 502, Data: "Unmarshall failed"}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, nil)

		fmt.Println(expectedResult)
		fmt.Println(actualResp)

		assert.Equal(t, expectedResult, actualResp)
	})
}

type r struct{}

func (r) Read(p []byte) (n int, err error) {
	return 0, errors.New(`failed to read`)
}

func Test_modifyData_readall_error(t *testing.T) {
	t.Run("read failed", func(t *testing.T) {
		mockRequest, _ := http.NewRequest("POST", "/modifyData", r{})

		_, readErr := ioutil.ReadAll(r{})

		expectedResult := ResponseData{Status: 500, Data: readErr.Error()}
		actualResp := ModifyData(httptest.NewRecorder(), mockRequest, nil)

		assert.Equal(t, expectedResult, actualResp)
	})
}
