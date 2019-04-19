package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorCPH/validator"
)

type fileParam struct {
	Image []byte `form:"image" valid:"required" type:"file" max_size:"61440"`
	Name  string `form:"name" json:"name" valid:"required" regexp:"^[a-zA-Z_][a-zA-Z_]*$"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	obj := fileParam{}
	err := validator.Bind(r, &obj)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]int{"image_size": len(obj.Image)})
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
