package common

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

type CosUtil struct {
	client *cos.Client
}

const COSURL = "https://dianshi-1251023989.cos.ap-chengdu.myqcloud.com"

func (util CosUtil) CreateCosClient() *cos.Client {
	u, _ := url.Parse(COSURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDlSyKUWjnAMUWNBFw8962796sX4d9Kv65",
			SecretKey: "veBbrx6KVVP8g2DICuSp00kUkutg2JPJ",
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	return c
}

func Delete(c *cos.Client, name string) {
	resp, err := c.Object.Delete(context.Background(), name)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}
}

func Put(c *cos.Client, name string, data string) {

	// Case1 normal put object
	f := strings.NewReader(data)

	resp, err := c.Object.Put(context.Background(), name, f, nil)

	fmt.Println(resp.Status)

	if err != nil {
		panic(err)
	}
}

func PutFile(c *cos.Client, name string, r io.Reader) {

	resp, err := c.Object.Put(context.Background(), name, r, nil)

	fmt.Println(resp.Status)

	if err != nil {
		panic(err)
	}
}
