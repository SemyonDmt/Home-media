package api

type Trackers struct {
	Title         string `json:"title"`
	Seeders       int    `json:"seeders"`
	Size          int    `json:"size"`
	Details       string `json:"details"`
	BlackholeLink string `json:"blackholeLink"`
}
