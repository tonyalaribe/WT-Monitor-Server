package resources

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestUserHandlers(t *testing.T) {
	router := gin.New()
	//router.GET("/getS", Screenshots{}.Get)

	router.GET("/get", Users{}.Get)

	convey.Convey("Test get users \n", t, func() {
		req, _ := http.NewRequest("GET", "/get", nil)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		t.Log(resp.Body.String())
		convey.So(resp.Code, convey.ShouldEqual, 200)

	})

}
