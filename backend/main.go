package main

import (
	"github.com/SemyonDmt/home-media/internal/api"
	"github.com/SemyonDmt/home-media/internal/jackett"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"sort"
)

func main() {
	hostAddress := os.Getenv("HOST_ADDRESS")
	url := os.Getenv("JACKETT_URL")
	key := os.Getenv("JACKETT_APIKEY")
	j := jackett.New(url, key)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "uri=${uri}, method=${method}, status=${status}, error=${error}\n",
	}))
	e.Static("/", "./frontend/dist")
	e.File("/", "./frontend/dist/index.html")
	e.POST("/api/items", searchTrackersHandler(j))
	e.Logger.Fatal(e.Start(hostAddress))
}

func searchTrackersHandler(j *jackett.Client) func(c echo.Context) error {
	return func(c echo.Context) error {
		r := new(api.Request)
		if err := c.Bind(r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if items, err := j.Search(r.Search); err != nil {
			return err
		} else {
			response := make([]api.Trackers, len(items.Results))
			for i, item := range items.Results {
				response[i] = api.Trackers{
					Title:         item.Title,
					Seeders:       item.Seeders,
					Size:          item.Size,
					Details:       item.Details,
					BlackholeLink: item.BlackholeLink,
				}
			}
			sort.Slice(response, func(i, j int) bool {
				return response[i].Seeders > response[j].Seeders
			})
			return c.JSON(http.StatusOK, response)
		}
	}
}
