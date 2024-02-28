package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("router", Router)
}

func ErrWriter(w http.ResponseWriter, err error) {
	var jsonBytes []byte
	jsonBytes, jsonErr := json.Marshal(map[string]interface{}{
		"err": fmt.Sprintf("%v", err),
	})
	if jsonErr != nil {
		jsonBytes = []byte(fmt.Sprintf("err: %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonBytes)
}

func Router(w http.ResponseWriter, r *http.Request) {
	dbconn := os.Getenv("DB_CONNECTION_STRING")
	env1 := os.Getenv("SERVICE_CONFIG_TEST")

	fmt.Println("DB_CONNECTION_STRING", dbconn)
	fmt.Println("SERVICE_CONFIG_TEST", env1)

	route := getFirstPath(r.URL.Path)

	switch route {
	case "hello":
		HelloRoute(w, r)
		return
	case "test":
		TestRoute(w, r)
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func getFirstPath(urlPath string) string {
	// Remove leading and trailing slashes
	urlPath = strings.Trim(urlPath, "/")

	// Split the URL path by slashes
	pathSegments := strings.Split(urlPath, "/")

	// Return the first path segment
	if len(pathSegments) > 0 {
		return pathSegments[0]
	}

	return ""
}
