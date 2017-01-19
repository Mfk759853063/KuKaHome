package helper

import (
	"encoding/json"
	"net/http"
)

// 通用设置header 解决跨域
func SetGlobalHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func WriteResponse(w http.ResponseWriter, code int, customCode int, parms interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": customCode,
		"data": parms,
	})
}
