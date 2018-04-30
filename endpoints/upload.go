package endpoints

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dfang/yuanxin/util"
	jwt "github.com/dgrijalva/jwt-go"
)

func UploadEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user := context.Get(r, "user")

		// user := r.Context().Value("user")
		// fmt.Fprintf(w, "This is an authenticated request\n")
		// fmt.Fprintf(w, "Claim content:\n")
		// for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		// 	fmt.Fprintf(w, "%s :\t%#v\n", k, v)
		// }
		// user := r.Context().Value("user")

		claims := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
		// user.(*jwt.Token)
		fmt.Println(claims["user"])

		// for k, v := range context.GetAll(r) {
		// 	fmt.Fprintf(w, "%s :\t%#v\n", k, v)
		// }

		// fmt.Println("method:", r.Method) //获取请求的方法
		file, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
			// fmt.Println(err)
			// return
		}
		defer file.Close()

		bs, err := ioutil.ReadAll(file)
		if err != nil {
			// 统一返回服务器内部处理错误
		}

		hash, err := util.UploadFile(bs)
		if err != nil {
			// 上传失败
			// 统一返回服务器内部处理错误
		}

		// fmt.Fprintf(w, "%v", handler.Header)
		// f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// defer f.Close()
		// io.Copy(f, file)
		baseURL := "http://p7ft1yl0b.bkt.clouddn.com/"
		// w.Write([]byte(baseURL + hash))

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
			URL        string `json:"url"`
		}{
			StatusCode: 200,
			Message:    "上传成功",
			URL:        baseURL + hash,
		})
		return
	})
}
