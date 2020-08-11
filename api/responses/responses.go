package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Header struct {
	Status  int
	Message string
}
type ResultJson struct {
	Header interface{} `json:header`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	header := Header{
		Status:  statusCode,
		Message: "success to Fetch data",
	}
	resultJson := ResultJson{
		Header: header,
		Status: "success",
		Data:   data,
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(resultJson)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
