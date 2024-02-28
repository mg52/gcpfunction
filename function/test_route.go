package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		header1 := r.Header.Get("Header1")

		jsonBytes, err := json.Marshal(map[string]interface{}{
			"router":     "test router",
			"status":     "success",
			"statusCode": 200,
			"id":         fmt.Sprintf("your id is %s", id),
			"data":       "data",
			"header1":    fmt.Sprintf("Header1 is %s", header1),
		})
		if err != nil {
			ErrWriter(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	default:
		ErrWriter(w, fmt.Errorf("unsupported method"))
		return
	}
}
