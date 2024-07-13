package tasks

import "fmt"

type Uploader struct {
}

func NewUploader() *Uploader {
	return &Uploader{}
}

func (u *Uploader) Do() {
	fmt.Println("hello!")
}
