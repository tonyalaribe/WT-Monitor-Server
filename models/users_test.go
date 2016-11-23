package models

import (
	"testing"

	"github.com/tonyalaribe/monitor-server/config"

	"github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {

	convey.Convey("Should Create User", t, func() {
		c := Client{Path: config.Get().DB.TestFile}
		err := c.Open()
		convey.So(err, convey.ShouldEqual, nil)
		defer c.DB.Close()
		testUser := User{
			UserName: "admin",
		}

		err = testUser.Create(c.DB)
		convey.So(err, convey.ShouldEqual, nil)
		result, err := testUser.GetAll(c.DB)
		convey.So(err, convey.ShouldEqual, nil)

		convey.So(result, convey.ShouldNotBeEmpty)
	})

}
