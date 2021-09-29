package crud

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_modifyData(t *testing.T) {
	bodyData := new(RequestModify)

	bodyData.Method = "insert"
	bodyData.Data.ID = "W009"
	bodyData.Data.Name = "Daniel"
	bodyData.Data.Age = 2
	bodyData.Data.Grade = 3

	marshalBody, _ := json.Marshal(bodyData)

	mockRequest, _ := http.NewRequest("POST", "/modifyData", bytes.NewBuffer(marshalBody))

	t.Run("Insert success", func(t *testing.T) {
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

	t.Run("Update success", func(t *testing.T) {
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
