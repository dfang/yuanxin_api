package endpoints

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dfang/yuanxin/util"
)

func UploadEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		fmt.Println("method:", r.Method) //获取请求的方法
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		bs, err := ioutil.ReadAll(file)

		hash := util.Upload(bs)

		fmt.Fprintf(w, "%v", handler.Header)
		w.Write([]byte(hash))
		// f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// defer f.Close()
		// io.Copy(f, file)
	})
}
