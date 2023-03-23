package cloud

import (
	"cloudstation/store"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var _ store.UpLoader = &TenCent{}

type TenCent struct {
	client *cos.Client
}

func NewTenCent(bucketUrl string, secretId string, secretKey string) *TenCent {
	// 在使用 API 请求时，对于私有桶您必须使用签名请求。通过永久密钥生成签名，
	// 放入 Authorization 头部中，形成签名请求；请求发送到 COS，COS 会验证签名与请求是否一致。
	u, _ := url.Parse(bucketUrl)
	client := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(secretId),
			SecretKey: os.Getenv(secretKey),
		},
	})

	return &TenCent{
		client: client,
	}
}

func (t *TenCent) UpLoad(filePath string, filename string, id string, key string) error {

	// 获取预签名 URL
	preUrl, err := t.client.Object.GetPresignedURL(context.Background(), http.MethodPut, filePath+filename, id, key, 10*time.Minute, nil)
	if err != nil {
		log.Panic("获取与预签名失败", err)
	}

	data := "test upload with preUrl"
	f := strings.NewReader(data)
	_, err = http.NewRequest(http.MethodPut, preUrl.String(), f)
	if err != nil {
		panic(err)
	}
	//req.Header.Set("Content-Type", "text/html")
	//_, err = http.DefaultClient.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	return nil
}
