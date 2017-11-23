package hret

import (
	"github.com/hzwy23/panda/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

// 返回信息数据结构
type RetContent struct {
	// 版本号
	Version    string      `json:"version"`
	// 状态吗
	Code       int         `json:"code"`
	// 返回消息
	Message    string      `json:"msg"`
	// 详细信息
	Details    interface{} `json:"details"`
	// 返回给客户端数据信息
	Rows       interface{} `json:"rows,omitempty"`
	// 信息记录数，用于分页查询，表示总共记录数
	Total      int64       `json:"total,omitempty"`
}

func NewRetContent() *RetContent{
	return &RetContent{
		Version:   "v1.0",
		Code:      200,
		Message:   "execute successfully.",
		Rows: "[]",
	}
}

func (rc *RetContent)SetVersion(str string)*RetContent{
	rc.Version = str
	return rc
}

func (rc *RetContent)SetCode(code int)*RetContent {
	rc.Code = code
	return rc
}

func (rc *RetContent)SetMessage(msg string) *RetContent {
	rc.Message = msg
	return rc
}

func (rc *RetContent)SetDetails(d interface{}) *RetContent {
	rc.Details = d
	return rc
}

func (rc *RetContent)SetRows(rows interface{}) *RetContent {
	rc.Rows = rows
	return rc
}

func (rc *RetContent)SetTotal(total int64) *RetContent{
	rc.Total = total
	return rc
}


func Write(w http.ResponseWriter, ret *RetContent) error {
	ojs, err := json.Marshal(ret)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{"code":"` + strconv.Itoa(http.StatusExpectationFailed) + `","msg":"`+err.Error()+`","details":"format json type info failed."}`))
		return err
	}
	w.Write(ojs)
	return nil
}


// 将数据打包成json格式，返回给客户端，
// 返回数据成功，error为nil，否则为错误信息
func Json(w http.ResponseWriter, data interface{}) error {
	ijs, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{"code":"428","msg":"` + err.Error() + `",details:"format json type info failed."}`))
		return  err
	}
	if string(ijs) == "null" {
		w.Write([]byte("[]"))
		return nil
	}
	_, err = w.Write(ijs)
	return  err
}


// 请求失败，返回JSON格式数据信息，code为返回状态码，msg是返回的信息，details是相信的信息
func Error(w http.ResponseWriter, code int, msg string, details ...interface{}) error{
	e := &RetContent{
		Code:        code,
		Message:     msg,
		Details:     details,
	}
	return writHttpError(w, e)
}

// 请求成功，返回JSON格式数据信息。
func Success(w http.ResponseWriter, v interface{}) error{
	ok := RetContent{
		Version:   "v1.0",
		Code:      200,
		Message:   "execute successfully.",
		Rows: v,
	}
	ojs, err := json.Marshal(ok)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{"code":"` + strconv.Itoa(http.StatusExpectationFailed) + `","msg":"format json type info failed.","details":"format json type info failed."}`))
		return err
	}
	w.Write(ojs)
	return nil
}

// 返回bootstrap-table后台分页时数据格式，total是总共数据条数，v是当前分页数据，
// 这个函数只适用于bootstrap-table前段插件采用后台数据分页时使用。
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
		w.Write([]byte(`{"code":"` + strconv.Itoa(http.StatusExpectationFailed) + `","msg":"format json type info failed.","details":"format json type info failed."}`))
		return
	}
	w.Write(ijs)
}

func writHttpError(w http.ResponseWriter, herr *RetContent) error{
	herr.Version = "v1.0"
	ijs, err := json.Marshal(herr)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{"code":"` + strconv.Itoa(herr.Code) + `","msg":"` + herr.Message + `","details":"format json type info failed."}`))
		return err
	}
	w.WriteHeader(herr.Code)
	_,err = w.Write(ijs)
	return err
}
