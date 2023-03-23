package store

type UpLoader interface {
	UpLoad(filePath string, filename string, id string, key string) error
}
