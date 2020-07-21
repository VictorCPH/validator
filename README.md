# Validator

[![Build Status](https://travis-ci.org/VictorCPH/validator.svg?branch=master)](https://travis-ci.org/VictorCPH/validator)

Validator is a http request parameters checker.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

```sh
$ go get github.com/VictorCPH/validator
```

## Import package in your project

Add following line in your `*.go` file:

```golang
import "github.com/VictorCPH/validator"
```

## Usage

### Basic

```golang
// examples/basic/main.go
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
```

```sh
# run examples/basic/main.go and visit localhost:8080
$ go run examples/basic/main.go
```

Test it with form:

```
$ curl -v -XPOST "localhost:8080" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "name=ming" \
  -d "age=20" \
  -d "passed=true" \
  -d "score=86.5" \
  -d "area=148.898877383" \
  -d "side=front" \
  -d "friends[]=Mary" -d "friends[]=Jack" \
  -d "scores[]=68.5" -d "scores[]=73.5"
```

Test it with query string:

```
$ curl -v -g -XGET "localhost:8080/?name=ming&age=20&passed=true&score=86.5&area=148.898877383&side=front&friends[]=Mary&friends[]=Jack&scores[]=68.5&scores[]=73.5"
```

Test it with json:

```
$ curl -v -XPOST "localhost:8080" \
  -H "Content-Type: application/json" \
  -d '{"name":"ming","age":20,"passed":true,"score":86.5,"area":148.898877383,"side":"front","friends":["Mary","Jack"],"scores":[68.5,73.5],"extra_info":"hello"}'
```

### Multipart file

```golang
// examples/file_bind/main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorCPH/validator"
)

type fileParam struct {
	Image []byte `form:"image" valid:"required" type:"file" max_size:"61440"`
	Name  string `form:"name" valid:"required" regexp:"^[a-zA-Z_][a-zA-Z_]*$"`
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
```

```sh
# run examples/file_bind/main.go and visit localhost:8080
$ go run examples/file_bind/main.go
```

Test it with:

```sh
$ curl -v -XPOST "localhost:8080" \
  -F "image=@testdata/Go-Logo_Aqua.jpg" \
  -F "name=ming"
```

### Base64 string

```golang
// examples/base64_bind/main.go
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
```

```sh
# run examples/base64_bind/main.go and visit localhost:8080
$ go run examples/base64_bind/main.go
```

Test it with:

```sh
$ curl -v -XPOST "localhost:8080" \
  -d "label=aGVsbG8="
```

## Support tags

``` sh
form, json, valid, default, type, values, min, max, range, regexp, max_size
```

- if use form format, you shold contain a `form` tag to give the name of the field.
- if use json format, you shold contain a `json` tag to give the name of the field.
- every param validation should contain a `valid` tag, it must be `required` or `optional`.
- `default` tag can only be used with `optional`.
- `values` tag can only be used with `int float32 float64 bool string`.
- `min, max, range` tag can only be used with `int, float32, float64`.
- `regexp` tag can only be used with `string`.
- `type` tag now only support `file` and `base64`.
- if `type:"file"`, it will read file as `[]byte`.
- if `type:"base64"`, it will read base64 string, then decode it and save as `[]byte`.
- `max_size` tag can only be used with `type:"file"`, it will check the max size of file.


Supported Types:

```sh
int, bool, float32, float64, string, slice, []byte
```

If you has more demands, report an [issue](https://github.com/VictorCPH/validator/issues/new), or open up a [pull request](https://github.com/VictorCPH/validator/pulls).

## License

The repo is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
