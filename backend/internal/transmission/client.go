package transmission

import (
	"context"
	"fmt"
	"github.com/hekmon/transmissionrpc/v2"
)

type Folders struct {
	Title string
	Name  string
}

type Client struct {
	transmissionUrl string
	user            string
	password        string
	folders         []Folders
}

func New(url, user, password string) *Client {
	return &Client{
		transmissionUrl: url,
		user:            user,
		password:        password,
		folders: []Folders{
			{Title: "Кино", Name: "movies"},
			{Title: "Сериалы", Name: "series"},
			{Title: "Детское кино", Name: "kids_movies"},
			{Title: "Детские сериалы", Name: "kids_series"},
			{Title: "Работа", Name: "work"},
			{Title: "Игры", Name: "games"},
		},
	}
}

func (c *Client) FoldersForDownload() []string {
	f := make([]string, len(c.folders))
	for i := range c.folders {
		f[i] = c.folders[i].Title
	}
	return f
}

func (c *Client) DownloadToDir(folderTitle, metaInfo string) error {
	transmission, err := transmissionrpc.New(c.transmissionUrl, c.user, c.password, nil)
	if err != nil {
		return err
	}

	toDir, err := c.convertToFolderName(folderTitle)
	if err != nil {
		return err
	}

	_, err = transmission.TorrentAdd(context.TODO(), transmissionrpc.TorrentAddPayload{
		DownloadDir: &toDir,
		MetaInfo:    &metaInfo,
	})

	return err
}

func (c *Client) convertToFolderName(folderTitle string) (string, error) {
	var toDir string
	for i := range c.folders {
		if c.folders[i].Title == folderTitle {
			toDir = "/downloads/" + c.folders[i].Name
			break
		}
	}

	if toDir == "" {
		return "", fmt.Errorf("download folder name '%s' not found", folderTitle)
	}
	return toDir, nil
}
