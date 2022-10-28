package util

import (
	"io"
	"net/http"
	"os"
)

var DownLoad = newDownLoad()

type downLoad struct {
}

func newDownLoad() *downLoad {
	return &downLoad{}
}

// DownloadFile 会将url下载到本地文件，它会在下载时写入，而不是将整个文件加载到内存中。
func (ud *downLoad) DownloadFile(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
