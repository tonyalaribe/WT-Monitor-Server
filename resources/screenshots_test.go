package resources

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func TestScreenshotHandlers(t *testing.T) {
	router := gin.New()
	//router.GET("/getS", Screenshots{}.Get)

	router.POST("/createS", Screenshots{}.Post)
	router.GET("/getS", Screenshots{}.Get)

	convey.Convey("Test adding new screenshots \n", t, func() {
		extraParams := map[string]string{
			"id": "admin",
		}

		req, err := newfileUploadRequest("/createS", extraParams, "image", "/Users/macbook/Desktop/cc.png")
		convey.So(err, convey.ShouldBeNil)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, 200)

		convey.Convey("Test return screenshots \n", func() {
			req, _ := http.NewRequest("GET", "/getS?id=admin", nil)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			//t.Log(resp.Body.String())
			convey.So(resp.Code, convey.ShouldEqual, 200)
		})

	})

}
