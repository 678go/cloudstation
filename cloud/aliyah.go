package cloud

import "fmt"

type Aliyah struct{}

func (a *Aliyah) UpLoad(filePath string, filename string, id string, key string) error {
	fmt.Println("上传了个寂寞")
	return nil
}
