package lib

import "net/http"
import "bytes"
import "io"

const HOST = "https://services.tokenview.com/vipapi"

func Post(path string, data []byte) (body io.ReadCloser, err error) {

	res, err := http.Post(HOST+path, "text/html", bytes.NewBuffer(data))

	if err != nil {
		return
	}

	return res.Body, nil
}

func Get(path string) (body io.ReadCloser, err error) {
	res, err := http.Get(HOST + path)
	if err != nil {
		return
	}

	return res.Body, nil
}
