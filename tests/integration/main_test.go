package tuneage_api

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	var err error
	var url string
	var client http.Client
	var res *http.Response
	var body []byte

	url = os.Getenv("apiURL")
	log.Print(url)
	assert.NotNil(t, url, "url should not be nil")

	client = http.Client{}
	assert.NotNil(t, client, "client should not be nil")

	res, err = client.Get(url)
	assert.Nil(t, err, "client.Get should not return an error")
	assert.Equal(t, res.StatusCode, 200)

	body, err = ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "ioutil.ReadAll should not return an error")
	assert.Contains(t, string(body[:]), "{ \"now\": ", "body should contain { \"now\": ")

	res.Body.Close()
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
