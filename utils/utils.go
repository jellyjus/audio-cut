package utils

import (
	"fmt"
	"io"
	"mime/multipart"
)

type Writer struct {
	*multipart.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		multipart.NewWriter(w),
	}
}

func (w *Writer) CreateFormAudioFile(fieldname string, file multipart.FileHeader) error {
	file.Header.Set("title", "test title")
	file.Header.Set("artist", "test artist")
	file.Header.Set("Content-Type", "multipart/form-data")
	fmt.Println(file.Header)
	part, err := w.CreatePart(file.Header)
	if err != nil {
		return err
	}

	fileSrc, _ := file.Open()
	defer fileSrc.Close()

	_, err = io.Copy(part, fileSrc)
	return err
}
