package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
)

func TestResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Data(w, req)

	approvals.VerifyJSONStruct(t, w)
}

func TestData(t *testing.T) {
	NowTimeFunc := MockNowTimeFunc
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Data(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	approvals.VerifyJSONBytes(t, data)
}

func MockNowTimeFunc() time.Time {
	return time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
}

func TestMain(m *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(m.Run())
}
