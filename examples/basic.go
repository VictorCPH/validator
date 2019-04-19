package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorCPH/validator"
)

type param struct {
	Name      string    `form:"name" json:"name" valid:"required" regexp:"^[a-zA-Z_][a-zA-Z_]*$"`
	Age       int       `form:"age" json:"age" valid:"required" range:"18|25"`
	Passed    bool      `form:"passed" json:"passed" valid:"required"`
	Score     float32   `form:"score" json:"score" valid:"required" min:"60.0"`
	Area      float64   `form:"area" json:"area" valid:"required" max:"200.0"`
	Side      string    `form:"side" json:"side" valid:"required" values:"front|back"`
	Friends   []string  `form:"friends" json:"friends" valid:"required" regexp:"^[a-zA-Z_][a-zA-Z_]*$"`
	Scores    []float32 `form:"scores" json:"scores" valid:"required" range:"60|100"`
	ExtraInfo string    `form:"extra_info" json:"extra_info" valid:"optional" default:"hello"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	obj := param{}
	err := validator.Bind(r, &obj)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
	} else {
		json.NewEncoder(w).Encode(obj)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
