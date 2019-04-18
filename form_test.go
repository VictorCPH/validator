package validator

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FormParam struct {
	Name      string    `form:"name" valid:"required"`
	Age       int       `form:"age" valid:"required"`
	Passed    bool      `form:"passed" valid:"required"`
	Score     float32   `form:"score" valid:"required"`
	Area      float64   `form:"area" valid:"required"`
	Friends   []string  `form:"friends" valid:"required"`
	Scores    []float32 `form:"scores" valid:"required"`
	ExtraInfo string    `form:"extra_info" valid:"required"`
}

func TestBindForm(t *testing.T) {
	testBindForm(t, "POST")
	testBindForm(t, "GET")
}

func testBindForm(t *testing.T, method string) {
	obj := FormParam{}

	body := url.Values{}
	body.Add("name", "Tony")
	body.Add("age", "18")
	body.Add("passed", "true")
	body.Add("score", "1.01")
	body.Add("area", "3.00000000001")
	body.Add("friends", "Jack")
	body.Add("friends", "Mary")
	body.Add("scores", "0.1")
	body.Add("scores", "0.2")
	body.Add("extra_info", "hello world")

	var req *http.Request
	if method == "GET" {
		req = request(method, "/?"+body.Encode(), "", "")
	} else {
		req = request(method, "/", body.Encode(), ContentTypeForm)
	}

	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "Tony", obj.Name)
	assert.Equal(t, 18, obj.Age)
	assert.Equal(t, true, obj.Passed)
	assert.Equal(t, float32(1.01), obj.Score)
	assert.Equal(t, float64(3.00000000001), obj.Area)
	assert.Equal(t, "Jack", obj.Friends[0])
	assert.Equal(t, "Mary", obj.Friends[1])
	assert.Equal(t, float32(0.1), obj.Scores[0])
	assert.Equal(t, float32(0.2), obj.Scores[1])
	assert.Equal(t, "hello world", obj.ExtraInfo)
}
