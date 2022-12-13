package api

type Trackers struct {
	Title        string `json:"title"`
	Seeders      int    `json:"seeders"`
	Size         int    `json:"size"`
	Details      string `json:"details"`
	TrackerId    string `json:"trackerId"`
	DownloadLink string `json:"downloadLink"`
}

type DownloadResult struct {
	Result       bool   `json:"result"`
	ErrorMessage string `json:"errorMessage"`
}
