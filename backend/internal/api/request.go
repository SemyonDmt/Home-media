package api

type Request struct {
	Search string `json:"search"`
}

type Download struct {
	TrackerId    string `json:"trackerId"`
	DownloadLink string `json:"downloadLink"`
}
