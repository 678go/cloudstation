package test

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/ylinyang/cloudstation/cloud"
	"github.com/ylinyang/cloudstation/store"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

var c store.UpLoader

func TestTUpLoad(t *testing.T) {
	if err := c.UpLoad("", "/test/aaa.txt", SecretId, SecretKey); err != nil {
		log.Println(err)
	}
}

func TestA(t *testing.T) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(BucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(SecretId), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(SecretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	name := "exampleobject"
	ctx := context.Background()
	f := strings.NewReader("test")

	// 获取预签名 URL
	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodPut, name, SecretId, SecretKey, time.Hour, nil)
	if err != nil {
		panic(err)
	}
	// 2. 通过预签名方式上传对象
	data := "test upload with presignedURL"
	f = strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPut, presignedURL.String(), f)
	if err != nil {
		panic(err)
	}
	// 用户可自行设置请求头部
	req.Header.Set("Content-Type", "text/html")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}

func init() {
	c = cloud.NewTenCent(BucketUrl, SecretId, SecretKey)
}
