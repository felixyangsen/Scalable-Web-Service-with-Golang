package crud

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetSingleStud(t *testing.T) {
// 	// bodyData := new(dt.RequestModify)
// 	param := map[string]interface{}{
// 		"param": "E001",
// 	}
// 	marshalBody, _ := json.Marshal(param)
// 	mockRequest, _ := http.NewRequest("POST", "/getSingleStud", bytes.NewBuffer(marshalBody))
// 	t.Run("Get single student", func(t *testing.T) {
// 		db, mock, err := sqlmock.New()
// 		if err != nil {
// 			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 		}
// 		defer db.Close()
// 		rows := sqlmock.NewRows([]string{"id", "name", "age", "grade"}).
// 			AddRow("E001", "daniel", "4", "3")
// 		mock.ExpectQuery("select id,name,age,grade from tb_student where id = ?").WithArgs(param["param"]).WillReturnRows(rows)
// 		expectedResult := ResponseData{Status: 200}
// 		actualResp := GetSingleStud(httptest.NewRecorder(), mockRequest, db)
// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 		}
// 		assert.Equal(t, expectedResult, actualResp)
// 	})
// }
