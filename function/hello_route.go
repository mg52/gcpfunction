package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HelloRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		header1 := r.Header.Get("Header1")

		jsonBytes, err := json.Marshal(map[string]interface{}{
			"router":     "hello router",
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
	case "POST":
		m := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			ErrWriter(w, err)
			return
		}
		name := m["name"]
		address := m["address"]
		var age int
		age, err := strconv.Atoi(fmt.Sprintf("%v", m["age"]))
		if err != nil {
			ErrWriter(w, err)
			return
		}

		header1 := r.Header.Get("Header1")

		jsonBytes, err := json.Marshal(map[string]interface{}{
			"router":     "hello router",
			"status":     "success",
			"statusCode": 200,
			"name":       fmt.Sprintf("your name is %s", name),
			"address":    fmt.Sprintf("your address is %s", address),
			"age":        fmt.Sprintf("your age is %d", age),
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
