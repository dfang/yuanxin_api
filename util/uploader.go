package util

import (
	"bytes"
	"context"
	"fmt"

	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/satori/go.uuid"
)

var (
	// accessKey = os.Getenv("QINIU_ACCESS_KEY")
	// secretKey = os.Getenv("QINIU_SECRET_KEY")
	// bucket    = os.Getenv("QINIU_TEST_BUCKET")
	accessKey = "MkFws9gjO_CScK5pXrahfBEWf9viOD_khTomtL3f"
	secretKey = "xVGWVTQTKFAlEEOFj6t4RRasJek5995UPlcMvv3M"
	bucket    = "yuanxin"
)

func Upload(data []byte) string {
	return uploadFile(data)
}

func UploadFile(data []byte) (string, error) {
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			// "x:name": "github logo2",
		},
	}

	// data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, genFilename(), bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// fmt.Println(ret.Key, ret.Hash)
	return ret.Key, nil
}

func genFilename() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
