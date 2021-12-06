package analysis

type Pv struct {
	Moves string `json:"moves"`
	Cp    int    `json:"cp"`
}

type Analysis struct {
	Fen    string `json:"fen"`
	Knodes int    `json:"knodes"`
	Depth  int    `json:"depth"`
	Pvs    []Pv   `json:"pvs"`
}

type Request struct {
	Fen     string
	MultiPV int
	Variant string
}
