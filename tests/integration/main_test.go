package tuneage_api

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	url := os.Getenv("apiURL")
	log.Print(url)
	assert.NotNil(t, url, "url should not be nil")
	client := http.Client{}
	assert.NotNil(t, client, "client should not be nil")
	err, res := client.Get(url)
	assert.NotNil(t, err, "client.Get should not return an error")
	assert.Contains(t, res, "{ \"now\": ")
}

// func _TestData(t *testing.T) {
// 	NowTimeFunc = MockNowTimeFunc

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	w := httptest.NewRecorder()

// 	Data(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()

// 	data, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	approvals.VerifyJSONBytes(t, data)
// }

// func MockNowTimeFunc() time.Time {
// 	fmt.Println("Using mock time")
// 	return time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
// }

// func TestMain(m *testing.M) {
// 	approvals.UseFolder("testdata")
// 	os.Exit(m.Run())
// }
