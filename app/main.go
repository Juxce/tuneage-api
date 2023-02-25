package tuneage_api

import (
	"fmt"
	"net/http"
	"time"
)

var NowTimeFunc func() time.Time = time.Now

func Data(w http.ResponseWriter, r *http.Request) {
	// if NowTimeFunc == nil {
	// 	NowTimeFunc = RealNowTimeFunc
	// }

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{ "now": %d }`, NowTimeFunc().UnixNano()/1000000)
}
