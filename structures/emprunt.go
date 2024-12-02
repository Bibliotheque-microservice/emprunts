package structures

type EmpruntReturned struct {
	EmpruntID int  `json:"empruntId"`
	Returned  bool `json:"returned"`
}
