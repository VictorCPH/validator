package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorCPH/validator"
)

type base64Param struct {
	Label []byte `form:"label" valid:"required" type:"base64"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	obj := base64Param{}
	err := validator.Bind(r, &obj)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"label": string(obj.Label)})
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
