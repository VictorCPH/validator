package validator

import (
	//"fmt"
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type defaultParam struct {
	Name string `form:"name" valid:"required"`
}

type optionalParam struct {
	Score float32 `form:"score" valid:"optional"`
}

type optionalWithDefaultParam struct {
	Score float32 `form:"score" valid:"optional" default:"60"`
}

type regexpTagParam struct {
	Name string `form:"name" valid:"required" regexp:"[A-Za-z0-9]+"`
}

type valuesTagParam struct {
	Side string `form:"side" valid:"required" values:"front|back"`
}

type fileMaxSizeTagParam struct {
	Image []byte `form:"image" valid:"required" type:"file" max_size:"1024"`
}

type base64Param struct {
	Label []byte `form:"label" valid:"required" type:"base64"`
}

type minTagIntParam struct {
	Age int `form:"age" valid:"required" min:"18"`
}

type minTagFloat32Param struct {
	Score float32 `form:"score" valid:"required" min:"60.0"`
}

type minTagFloat64Param struct {
	Area float64 `form:"area" valid:"required" min:"100.000000001"`
}

type maxTagIntParam struct {
	Age int `form:"age" valid:"required" max:"25"`
}

type rangeTagIntParam struct {
	Age int `form:"age" valid:"required" range:"18|25"`
}

func TestOptional(t *testing.T) {
	obj := optionalParam{}
	emptyBody := url.Values{}
	req := request("POST", "/", emptyBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, float32(0), obj.Score)

	obj2 := optionalWithDefaultParam{}
	err = Bind(req, &obj2)
	assert.NoError(t, err)
	assert.Equal(t, float32(60), obj2.Score)
}

func TestInvalidUTF8(t *testing.T) {
	obj := defaultParam{}

	badBody := url.Values{}
	badBody.Add("name", "a\xc5z")
	req := request("POST", "/", badBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "name: invalid utf8 string", err.Error())
}

func TestBlankString(t *testing.T) {
	obj := defaultParam{}

	badBody := url.Values{}
	badBody.Add("name", "")
	req := request("POST", "/", badBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "name: blank string", err.Error())
}

func TestRegexpTag(t *testing.T) {
	obj := regexpTagParam{}

	badBody := url.Values{}
	badBody.Add("name", "?&^_")
	req := request("POST", "/", badBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "name: wrong format, shold match regexp `[A-Za-z0-9]+`", err.Error())
}

func TestValuesTag(t *testing.T) {
	obj := valuesTagParam{}

	frontBody := url.Values{}
	frontBody.Add("side", "front")
	req := request("POST", "/", frontBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "front", obj.Side)

	backBody := url.Values{}
	backBody.Add("side", "back")
	req = request("POST", "/", backBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "back", obj.Side)

	badBody := url.Values{}
	badBody.Add("side", "other_value")
	req = request("POST", "/", badBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "side: other_value is not in [front back]", err.Error())
}

func TestFileMaxSize(t *testing.T) {
	obj := fileMaxSizeTagParam{}
	smallFile := map[string]string{"image": "testdata/broken.jpg"}
	req := requestMultipartForm("/", nil, smallFile)
	err := Bind(req, &obj)
	assert.NoError(t, err)

	largeFile := map[string]string{"image": "testdata/Go-Logo_Aqua.jpg"}
	req = requestMultipartForm("/", nil, largeFile)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "image: file larger than 1024 bytes", err.Error())
}

func TestBase64(t *testing.T) {
	obj := base64Param{}

	body := url.Values{}
	encoded := base64.StdEncoding.EncodeToString([]byte("hello"))
	body.Add("label", encoded)
	req := request("POST", "/", body.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello"), obj.Label)
}

func TestMinIntTag(t *testing.T) {
	obj := minTagIntParam{}

	body := url.Values{}
	body.Add("age", "20")
	req := request("POST", "/", body.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 20, obj.Age)

	badBody := url.Values{}
	badBody.Add("age", "16")
	req = request("POST", "/", badBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "age: smaller than 18", err.Error())
}

func TestMinFloat32Tag(t *testing.T) {
	obj := minTagFloat32Param{}

	body := url.Values{}
	body.Add("score", "80.0")
	req := request("POST", "/", body.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, float32(80.0), obj.Score)

	badBody := url.Values{}
	badBody.Add("score", "59.999")
	req = request("POST", "/", badBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "score: smaller than 60.0", err.Error())
}

func TestMinFloat64Tag(t *testing.T) {
	obj := minTagFloat64Param{}

	body := url.Values{}
	body.Add("area", "150.555555555")
	req := request("POST", "/", body.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, float64(150.555555555), obj.Area)

	badBody := url.Values{}
	badBody.Add("area", "100.000000000")
	req = request("POST", "/", badBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "area: smaller than 100.000000001", err.Error())
}

func TestMaxIntTag(t *testing.T) {
	obj := maxTagIntParam{}

	body := url.Values{}
	body.Add("age", "20")
	req := request("POST", "/", body.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 20, obj.Age)

	badBody := url.Values{}
	badBody.Add("age", "28")
	req = request("POST", "/", badBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "age: greater than 25", err.Error())
}

func TestRangeIntTag(t *testing.T) {
	obj := rangeTagIntParam{}

	smallerBody := url.Values{}
	smallerBody.Add("age", "16")
	req := request("POST", "/", smallerBody.Encode(), ContentTypeForm)
	err := Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "age: not in range (18, 25)", err.Error())

	body := url.Values{}
	body.Add("age", "20")
	req = request("POST", "/", body.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 20, obj.Age)

	greaterBody := url.Values{}
	greaterBody.Add("age", "28")
	req = request("POST", "/", greaterBody.Encode(), ContentTypeForm)
	err = Bind(req, &obj)
	assert.Error(t, err)
	assert.Equal(t, "age: not in range (18, 25)", err.Error())
}
