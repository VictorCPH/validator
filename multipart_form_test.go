package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MultipartFormParam struct {
	Image []byte `form:"image" valid:"required" type:"file"`
	Name  string `form:"name" valid:"required"`
}

func TestMultipartForm(t *testing.T) {
	obj := MultipartFormParam{}

	params := map[string]string{"name": "Tony"}
	files := map[string]string{"image": "testdata/Go-Logo_Aqua.jpg"}

	req := requestMultipartForm("/", params, files)

	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "Tony", obj.Name)
	assert.True(t, len(obj.Image) > 0)
}
