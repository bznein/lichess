package tablebase

type TablebaseMove struct {
	Uci     string `json:"uci"`
	San     string `json:"san"`
	Zeroing bool   `json:"zeroing"`
	MoveResult
}

type Tablebase struct {
	Moves []TablebaseMove `json:"moves"`
	MoveResult
}

type MoveResult struct {
	Category             string `json:"category"`
	DTZ                  int    `json:"dtz"`
	PreciseDTZ           int    `json:"precise_dtz"`
	DTM                  int    `json:"dtm"`
	Checkmate            bool   `json:"checkmate"`
	Stalemate            bool   `json:"stalemate"`
	VariantWin           bool   `json:"variant_win"`
	VariantLoss          bool   `json:"variant_loss"`
	InsufficientMaterial bool   `json:"insufficient_material"`
}
