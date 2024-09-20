package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	jsonit "github.com/json-iterator/go"
)

func multipartUpload(filename string, targetURL string, chunkSize int) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	index := 0

	ch := make(chan int)
	buf := make([]byte, chunkSize) //每次读取chunkSize大小的内容
	for {
		n, err := bfRd.Read(buf)
		if n <= 0 {
			break
		}
		index++

		bufCopied := make([]byte, 5*1048576)
		copy(bufCopied, buf)

		go func(b []byte, curIdx int) {
			fmt.Printf("upload_size: %d\n", len(b))

			resp, err := http.Post(
				targetURL+"&index="+strconv.Itoa(curIdx),
				"multipart/form-data",
				bytes.NewReader(b))
			if err != nil {
				fmt.Println(err)
			}

			body, er := ioutil.ReadAll(resp.Body)
			fmt.Printf("%+v %+v\n", string(body), er)
			resp.Body.Close()

			ch <- curIdx
		}(bufCopied[:n], index)

		//遇到任何错误立即返回，并忽略 EOF 错误信息
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	for idx := 0; idx < index; idx++ {
		select {
		case res := <-ch:
			fmt.Println(res)
		}
	}

	return nil
}

func main() {
	username := "island"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1NTY3NTc4Mjg4NjY0MTI1NDQsInVzZXJuYW1lIjoidXNlcm5hbWUiLCJleHAiOjE3NTgzNjk1NTEsImlzcyI6ImZpbGVzdG9yZSJ9.6ktw4b_gfChSVlK9NUjtvUzFPuK9mhuqy3BQTT-XA9g"
	filehash := "4780e2723a5b03d8c3f78b5ca497fac054afc8be"

	// 1. 请求初始化分块上传接口
	resp, err := http.PostForm(
		"http://localhost:8084/file/mpupload/init",
		url.Values{
			"username": {username},
			"token":    {token},
			"filehash": {filehash},
			"filesize": {"61348679"},
		})

	// 1. 请求初始化分块上传接口
	resp, err = http.PostForm(
		"http://localhost:8084/file/mpupload/init",
		url.Values{
			"username": {username},
			"token":    {token},
			"filehash": {filehash},
			"filesize": {"61348679"},
		})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 2. 解析 JSON 响应
	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			FileHash   string `json:"FileHash"`
			FileSize   int    `json:"FileSize"`
			UploadID   string `json:"UploadID"`
			ChunkSize  int    `json:"ChunkSize"`
			ChunkCount int    `json:"ChunkCount"`
		} `json:"data"`
		Location string `json:"location"`
	}

	err = jsonit.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 3. 得到 uploadID 和 chunkSize
	uploadID := result.Data.UploadID
	chunkSize := result.Data.ChunkSize
	fmt.Printf("uploadid: %s  chunksize: %d\n", uploadID, chunkSize)
	// 3. 请求分块上传接口
	filename := "/home/xusir/project/filestore/test/best.pt"
	tURL := "http://localhost:8084/file/mpupload/uppart?" +
		"username=island&token=" + token + "&uploadid=" + uploadID
	multipartUpload(filename, tURL, chunkSize)

	// 4. 请求分块完成接口
	resp, err = http.PostForm(
		"http://localhost:8084/file/mpupload/complete",
		url.Values{
			"username": {username},
			"token":    {token},
			"filehash": {filehash},
			"filesize": {"61348679"},
			"filename": {"best.pt"},
			"uploadid": {uploadID},
		})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("complete result: %s\n", string(body))
}
