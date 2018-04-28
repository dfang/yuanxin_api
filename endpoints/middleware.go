package endpoints

import (
	"net/http"

	"github.com/dfang/yuanxin/util"
)

// recovery middleware, compitable with negroni
// panic in http handler, recover here responds error to client
// recover here instead of every handler
// 一言不合就panic, 然后这里集中处理
// 注意app.go 里的 	n.Use(negroni.HandlerFunc(endpoints.Recovery))
// common 定义各种error 类型
func Recovery(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func() {
		err := recover()
		if err != nil {
			switch v := err.(type) {
			case *MissParameterError:
				util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
			case *ParseError:
				util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
			case *RecordNotFound:
				util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
			case string:
				util.RespondWithError(w, http.StatusBadRequest, 400, v)
			case error:
				util.RespondWithError(w, http.StatusBadRequest, 400, v.Error())
			default:
				util.RespondWithError(w, http.StatusBadRequest, 400, "Internal Error")
			}
			return
		}
	}()

	next(w, req)
}
