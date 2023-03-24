package cloud

import (
	"context"
	"errors"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/ylinyang/cloudstation/store"
	"log"
	"net/http"
	"net/url"
	"os"
)

var _ store.UpLoader = &TenCent{}

type TenCent struct {
	client   *cos.Client
	listener cos.ProgressListener
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
		client:   client,
		listener: &selfListener{},
	}
}

func (t *TenCent) UpLoad(filePath string, filename string, id string, key string) error {
	if filename == "" && filePath == "" {
		return errors.New("filename或者filepath为空")
	}
	// 上传进度条
	opt := &cos.ObjectPutOptions{ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
		ContentType: "text/html",
		Listener:    t.listener},
	}

	// 获取预签名 URL
	//preUrl, err := t.client.Object.GetPresignedURL(context.Background(), http.MethodPut, filePath+filename, id, key, 10*time.Second, nil)
	//if err != nil {
	//	log.Panic("获取与预签名失败", err)
	//}

	// 单个文件
	//content, err := os.ReadFile("./" + filename)
	//if err != nil {
	//	log.Panicf("读取%s文件失败", filename)
	//}
	// 文件夹 todo
	// 多个上传 todo
	//f := strings.NewReader(string(content))

	// 公有读写
	_, err := t.client.Object.PutFromFile(context.Background(), filePath, filename, opt)
	if err != nil {
		log.Println(err)
	}

	//req, err := http.NewRequest(http.MethodPut, preUrl.String(), f)
	//if err != nil {
	//	panic(err)
	//}
	//
	//req.Header.Set("Content-Type", "text/html")
	//_, err = http.DefaultClient.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("上传盘：上传成功")
	return nil
}
