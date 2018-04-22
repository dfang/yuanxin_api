package util

import (
	"bytes"
	"context"
	"testing"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// var (
// 	accessKey = "MkFws9gjO_CScK5pXrahfBEWf9viOD_khTomtL3f"
// 	secretKey = "xVGWVTQTKFAlEEOFj6t4RRasJek5995UPlcMvv3M"
// 	bucket    = "yuanxin"
// )

func TestUpload(t *testing.T) {

	t.Log(accessKey)
	t.Log(secretKey)
	t.Log(bucket)

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
			"x:name": "github logo",
		},
	}

	data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, accessKey, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret.Key, ret.Hash)
}
