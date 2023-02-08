package errors

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	unknownCode defaultCoder = defaultCoder{1, http.StatusInternalServerError, "An internal server error occurred", ""}
)

type Coder interface {
	// 获取与业务错误码关联的http状态码
	HTTPStatus() int

	// 获取直接展示给用户的错误信息
	String() string

	// 获取可以直接展示给用户的错误文档
	Reference() string

	// 获取业务错误码
	Code() int
}

type defaultCoder struct {
	C    int
	HTTP int
	Ext  string
	Ref  string
}

func (coder defaultCoder) Code() int {
	return coder.Code()
}

func (coder defaultCoder) String() string {
	return coder.Ext
}

func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

func (coder defaultCoder) Reference() string {
	return coder.Ref
}

var coders = map[int]Coder{}
var codeMux = &sync.Mutex{}

// 注册用户自定义的coder对象
func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	coders[coder.Code()] = coder
}

// 注册用户自定义的coder对象，注册相同的错误码时会panic
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := coders[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	coders[coder.Code()] = coder
}

// ParseCoder将指定的error解析到withCode对象
func ParseCoder(err error) Coder {

}

// 判断err的错误链中是否包含指定code的错误对象
func IsCode(err error, code int) {

}

func init() {
	coders[unknownCode.Code()] = unknownCode
}
