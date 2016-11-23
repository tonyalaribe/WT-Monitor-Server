package models

import (
	"testing"

	"github.com/tonyalaribe/monitor-server/config"

	"github.com/smartystreets/goconvey/convey"
)

func TestTimelog(t *testing.T) {

	convey.Convey("Should login and logout", t, func() {
		c := Client{Path: config.Get().DB.TestFile}
		err := c.Open()
		convey.So(err, convey.ShouldEqual, nil)
		defer c.DB.Close()
		testUser := UserTimeLog{
			UserID: "admin",
		}

		err = testUser.LoginNow(c.DB)
		convey.So(err, convey.ShouldEqual, nil)
		err = testUser.LogoutNow(c.DB)
		convey.So(err, convey.ShouldEqual, nil)
		result, err := testUser.GetTimeline(c.DB)
		convey.So(err, convey.ShouldEqual, nil)

		convey.So(result, convey.ShouldNotBeEmpty)
	})

}
