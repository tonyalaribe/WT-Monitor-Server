package models

/*
func TestScreeenshots(t *testing.T) {

	convey.Convey("Should Create Screenshot", t, func() {

		c := Client{Path: config.Get().DB.TestFile}
		err := c.Open()
		convey.So(err, convey.ShouldEqual, nil)
		defer c.DB.Close()
		testUser := User{
			UserName: "admin",
		}

		err = testUser.Create(c.DB)
		convey.So(err, convey.ShouldEqual, nil)

		testData := Screenshot{
			UserID:    "admin",
			Timestamp: time.Now(),
			URL:       "//placehold.it/700x700",
		}

		//err = testData.Create(c.DB)
		//convey.So(err, convey.ShouldEqual, nil)

		convey.Convey("Should return Screenshots", func() {
			var result []Screenshot
			result, err = testData.GetTodaysScreenshots(c.DB, "admin")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(result, convey.ShouldNotBeEmpty)

			t.Log(result)
		})
		tx, err := c.DB.Begin(true)
		if err != nil {
			log.Println(err)
		}
		defer tx.Rollback()

		b := tx.Bucket([]byte(constants.USER_BUCKET))

		recur(b, 5)

	})

}

func recur(buk *bolt.Bucket, level int) {
	if buk == nil {
		return
	}
	c := buk.Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
		fmt.Printf("%s:%s\n", k, v)
		if v == nil {
			recur(buk.Bucket(k), level+1)
		}
	}
}
*/
