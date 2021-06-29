package core

type DiskStatus struct {
	Total             float64 `json:"total"`
	Used              float64 `json:"used"`
	Available         float64 `json:"available"`
	Used_Percent      float64 `json:"used-percent"`
	Available_Percent float64 `json:"available-percent"`
}
