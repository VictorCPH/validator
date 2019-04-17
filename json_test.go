package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type JsonParam struct {
	Name      string    `json:"name" valid:"required"`
	Age       int       `json:"age" valid:"required"`
	Passed    bool      `json:"passed" valid:"required"`
	Score     float32   `json:"score" valid:"required"`
	Area      float64   `json:"area" valid:"required"`
	Friends   []string  `json:"friends" valid:"required"`
	Scores    []float32 `form:"scores" valid:"required"`
	ExtraInfo string    `json:"extra_info" valid:"required"`
}

func TestJson(t *testing.T) {
	obj := JsonParam{}
	body := `{
		"name":"Tony",
		"age": 18,
		"passed": true,
		"score": 1.01,
		"area": 3.00000000001,
		"friends": ["Jack", "Mary"],
		"scores": [0.1, 0.2],
		"extra_info": "hello world"
	}`

	req := request("POST", "/", body)
	req.Header.Add("Content-type", ContentTypeJson)

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
