package domain

type Store struct {
	Content        string   `json:"Content"`
	CoverPhotoPath string   `json:"CoverPhotoPath"`
	District       string   `json:"District"`
	IsActive       bool     `json:"IsActive"`
	MapsLink       string   `json:"MapsLink"`
	PlacesName     string   `json:"PlacesName"`
	Province       string   `json:"Province"`
	Rate           int32    `json:"Rate"`
	Services       []string `json:"Services"`
	Title          string   `json:"Title"`
}
