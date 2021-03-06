package httptestdemo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesthelloHandler(t testing.T) {
	//定义两个测试用例
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"base case", `{"name": "liwenzhou"}`, "hello liwenzhou"},
		{"bad case", "", "we need a name"},
	}
	r := SetupRouter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//mock 一个HTTP请求
			req := httptest.NewRequest(
				"POST",
				"/hello",
				strings.NewReader(tt.param), //请求参数
			)

			//mock 一个响应记录器
			w := httptest.NewRecorder()
			//让server段处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			//校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)
			//解析并检验响应内容是否符合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
