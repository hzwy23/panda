package hret

import (
	"github.com/hzwy23/panda/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

type RetContent struct {
	Version    string      `json:"version"`
	Code       int         `json:"code"`
	Message    string      `json:"msg"`
	Details    interface{} `json:"details"`
	Data       interface{} `json:"data,omitempty"`
	Total      int64       `json:"total,omitempty"`
	Rows       interface{} `json:"rows,omitempty"`
}

// 将数据打包成json格式，返回给客户端，
// 返回数据成功，error为nil，否则为错误信息
func Json(w http.ResponseWriter, data interface{}) (error) {
	ijs, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{code:765,error_msg:"` + err.Error() + `",details:"format json type info failed."}`))
		return  err
	}
	if string(ijs) == "null" {
		w.Write([]byte("[]"))
		return nil
	}
	_, err = w.Write(ijs)
	return  err
}

func Error(w http.ResponseWriter, code int, msg string, details ...interface{}) {
	e := RetContent{
		Code:        code,
		Message:     msg,
		Details:     details,
	}
	writHttpError(w, e)
}

func Success(w http.ResponseWriter, v interface{}) {
	ok := RetContent{
		Version:   "v1.0",
		Code:      200,
		Message:   "execute successfully.",
		Data:      v,
	}
	ojs, err := json.Marshal(ok)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(http.StatusExpectationFailed) + `,error_msg:"format json type info failed.",error_details:"format json type info failed."}`))
		return
	}
	w.Write(ojs)
	return
}

func BootstrapTable(w http.ResponseWriter, total int64, v interface{}) {
	ok := RetContent{
		Version:    "v1.0",
		Code: 200,
		Message:  "execute successfully.",
		Rows:       v,
		Total:      total,
	}

	ijs, err := json.Marshal(ok)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(http.StatusExpectationFailed) + `,error_msg:"format json type info failed.",error_details:"format json type info failed."}`))
		return
	}
	w.Write(ijs)
}

func writHttpError(w http.ResponseWriter, herr RetContent) {
	herr.Version = "v1.0"
	ijs, err := json.Marshal(herr)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{code:` + strconv.Itoa(herr.Code) + `,msg:"` + herr.Message + `",details:"format json type info failed."}`))
		return
	}
	w.WriteHeader(herr.Code)
	w.Write(ijs)
}
