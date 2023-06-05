package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	controller "github.com/sanmathi-g/restapi/controllers"
	"github.com/sanmathi-g/restapi/database"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	r := gin.Default()

	Convey("testing all the functions", t, func() {
		database.DatabaseConnection()
		r.GET("/movies/:id", controller.ReadMovie)
		r.GET("/movies", controller.Readmovies)
		r.POST("/movies", controller.CreateMovie)
		r.PUT("/moviess/:id", controller.UpdateMovie)
		r.DELETE("/moviess/:id", controller.DeleteMovie)

		Convey("when calling GET /movies/:id", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/movies/:id", nil)
			r.ServeHTTP(w, req)

			Convey("Then respond should have status 200 ok", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})
		Convey("when calling GET /movies", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/movies", nil)
			r.ServeHTTP(w, req)

			Convey("Then respond should have status 200 ok", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})
		Convey("when calling POST /movies", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/movies", nil)
			r.ServeHTTP(w, req)

			Convey("Then respond should have status 400 bad request ok", func() {
				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})
		})
		database.DB.Migrator().DropTable(&database.Movie{})
	})

}
