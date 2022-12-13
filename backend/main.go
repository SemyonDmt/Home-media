package main

import (
	"github.com/SemyonDmt/home-media/internal/api"
	"github.com/SemyonDmt/home-media/internal/jackett"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"sort"
	"strings"
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
	e.POST("/api/download", downloadTrackersHandler(j))
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
				const p = "&path="
				response[i] = api.Trackers{
					Title:        item.Title,
					Seeders:      item.Seeders,
					Size:         item.Size,
					Details:      item.Details,
					TrackerId:    item.TrackerID,
					DownloadLink: item.BlackholeLink[strings.Index(item.BlackholeLink, p)+1:],
				}
			}
			sort.Slice(response, func(i, j int) bool {
				return response[i].Seeders > response[j].Seeders
			})
			return c.JSON(http.StatusOK, response)
		}
	}
}

func downloadTrackersHandler(j *jackett.Client) func(c echo.Context) error {
	return func(c echo.Context) error {
		r := new(api.Download)
		if err := c.Bind(r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := j.Download(r.TrackerId, r.DownloadLink); err != nil {
			return c.JSON(http.StatusOK, api.DownloadResult{Result: false, ErrorMessage: err.Error()})
		} else {
			return c.JSON(http.StatusOK, api.DownloadResult{Result: true})
		}
	}
}
