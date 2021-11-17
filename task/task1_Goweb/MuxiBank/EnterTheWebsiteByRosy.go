package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {
	//开始访问

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code", nil)
	req.Header.Add("Code", "120")
	res, _ := client.Do(req)
	fmt.Println(res.Header)
	passport := res.Header["Passport"][0]
	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(content))
	//解析密文

	imfor := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	im, _ := base64.StdEncoding.DecodeString(imfor)
	fmt.Println(string(im))
	secret_key := "MuxiStudio203304"
	error_code := "for {go func(){time.Sleep(1*time.Hour)}()}"
	//加密error_code

	plaintext := []byte(error_code)
	key := []byte(secret_key)
	ivT := make([]byte, aes.BlockSize+len(plaintext))
	iv := ivT[:aes.BlockSize]
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	error := base64.StdEncoding.EncodeToString(crypted)
	//获得加密的error_code，并破门而入！

	type Body struct {
		Content string
	}
	body := Body{
		Content: error,
	}
	by, _ := json.Marshal((body))
	payload := strings.NewReader(string(by))
	req, _ = http.NewRequest("PUT", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", payload)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	//第二扇门破破破！

	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	//图片转码

	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	f, _ := os.Open("./t.txt")
	var b = make([]byte, 136344)
	f.Read(b)
	imageBytes, _ := base64.StdEncoding.DecodeString(string(b))
	ioutil.WriteFile("./test.jpg", imageBytes, 0777)
	//输入图片访问

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodywriter, _ := bodyWriter.CreateFormFile("file", "test.jpg")
	file, _ := os.Open("./test.jpg")
	io.Copy(bodywriter, file)
	bodyWriter.Close() //此处特别注意
	req, _ = http.NewRequest("POST", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", bodyBuf)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	res, _ = client.Do(req)
	//fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	//进入最后一道门

	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	//破开最后一道门

	bodyBuf = &bytes.Buffer{}
	bodyWriter = multipart.NewWriter(bodyBuf)
	bodywriter, _ = bodyWriter.CreateFormFile("file", "permute.go")
	file, _ = os.Open("./permute.go")
	io.Copy(bodywriter, file)
	bodyWriter.Close()
	req, _ = http.NewRequest("POST", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", bodyBuf)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	res, _ = client.Do(req)
	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contents))
	//我就知道你能成功！Backend组织欢迎你！

}
