package jackett

type Response struct {
	Results []struct {
		FirstSeen            string        `json:"FirstSeen"`
		Tracker              string        `json:"Tracker"`
		TrackerID            string        `json:"TrackerId"`
		TrackerType          string        `json:"TrackerType"`
		CategoryDesc         string        `json:"CategoryDesc"`
		BlackholeLink        string        `json:"BlackholeLink"`
		Title                string        `json:"Title"`
		GUID                 string        `json:"Guid"`
		Link                 string        `json:"Link"`
		Details              string        `json:"Details"`
		PublishDate          TimeRFC3339   `json:"PublishDate"`
		Category             []int         `json:"Category"`
		Size                 int           `json:"Size"`
		Files                interface{}   `json:"Files"`
		Grabs                int           `json:"Grabs"`
		Description          interface{}   `json:"Description"`
		RageID               interface{}   `json:"RageID"`
		TVDBID               interface{}   `json:"TVDBId"`
		Imdb                 interface{}   `json:"Imdb"`
		TMDb                 interface{}   `json:"TMDb"`
		TVMazeID             interface{}   `json:"TVMazeId"`
		TraktID              interface{}   `json:"TraktId"`
		DoubanID             interface{}   `json:"DoubanId"`
		Genres               []interface{} `json:"Genres"`
		Year                 interface{}   `json:"Year"`
		Author               interface{}   `json:"Author"`
		BookTitle            interface{}   `json:"BookTitle"`
		Publisher            interface{}   `json:"Publisher"`
		Artist               interface{}   `json:"Artist"`
		Album                interface{}   `json:"Album"`
		Label                interface{}   `json:"Label"`
		Track                interface{}   `json:"Track"`
		Seeders              int           `json:"Seeders"`
		Peers                int           `json:"Peers"`
		Poster               interface{}   `json:"Poster"`
		InfoHash             interface{}   `json:"InfoHash"`
		MagnetURI            interface{}   `json:"MagnetUri"`
		MinimumRatio         float64       `json:"MinimumRatio"`
		MinimumSeedTime      int           `json:"MinimumSeedTime"`
		DownloadVolumeFactor float64       `json:"DownloadVolumeFactor"`
		UploadVolumeFactor   float64       `json:"UploadVolumeFactor"`
		Gain                 float64       `json:"Gain"`
	} `json:"Results"`
	Indexers []struct {
		ID      string      `json:"ID"`
		Name    string      `json:"Name"`
		Status  int         `json:"Status"`
		Results int         `json:"Results"`
		Error   interface{} `json:"Error"`
	} `json:"Indexers"`
}

type Download struct {
	Result string `json:"result"`
}
