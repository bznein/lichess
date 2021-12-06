package account

type PerfStats struct {
	Games  int  `json:"games"`
	Rating int  `json:"rating"`
	Rd     int  `json:"rd"`
	Prog   int  `json:"prog"`
	Prov   bool `json:"prov"`
}

type Perf struct {
	Chess960       PerfStats `json:"chess960"`
	Atomic         PerfStats `json:"atomic"`
	RacingKings    PerfStats `json:"racingKings"`
	UltraBullet    PerfStats `json:"ultraBullet"`
	Blitz          PerfStats `json:"blitz"`
	KingOfTheHill  PerfStats `json:"kingOfTheHill"`
	Bullet         PerfStats `json:"bullet"`
	Correspondence PerfStats `json:"correspondence"`
	Horde          PerfStats `json:"horde"`
	Puzzle         PerfStats `json:"puzzle"`
	Classical      PerfStats `json:"classical"`
	Rapid          PerfStats `json:"rapid"`
	Storm          PerfStats `json:"storm"`
}

type Profile struct {
	Country    string `json:"country"`
	Location   string `json:"location"`
	Bio        string `json:"bio"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FideRating int    `json:"fideRating"`
	UscfRating int    `json:"uscfRating"`
	EcfRating  int    `json:"ecfRating"`
	Links      string `json:"links"`
}

type PlayTime struct {
	Total int `json:"total"`
	Tv    int `json:"tv"`
}

type Count struct {
	All      int `json:"all"`
	Rated    int `json:"rated"`
	Ai       int `json:"ai"`
	Draw     int `json:"draw"`
	DrawH    int `json:"drawH"`
	Loss     int `json:"loss"`
	LossH    int `json:"lossH"`
	Win      int `json:"win"`
	WinH     int `json:"winH"`
	Bookmark int `json:"bookmark"`
	Playing  int `json:"playing"`
	Import   int `json:"import"`
	Me       int `json:"me"`
}

type Email struct {
	Email string `json:"email"`
}

type Account struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	Online         bool     `json:"online"`
	Perfs          Perf     `json:"perfs"`
	CreatedAt      int      `json:"createdAt"`
	Disabled       bool     `json:"disabled"`
	TosViolation   bool     `json:"tosViolation"`
	Profile        Profile  `json:"profile"`
	SeenAt         int      `json:"seenAt"`
	Patron         bool     `json:"patron"`
	Verified       bool     `json:"verified"`
	PlayTime       PlayTime `json:"playTime"`
	Title          string   `json:"title"`
	Url            string   `json:"url"`
	Playing        string   `json:"playing"`
	CompletionRate int      `json:"completionRate"`
	Streaming      bool     `json:"streaming"`
	Followable     bool     `json:"followable"`
	Following      bool     `json:"following"`
	Blocking       bool     `json:"blocking"`
	FollowsYou     bool     `json:"followsYou"`
}

type Prefs struct {
	Dark          bool   `json:"dark"`
	Trans         bool   `json:"trans"`
	BgImg         string `json:"bgImg"`
	Is3d          bool   `json:"is3D"`
	Theme         string `json:"theme"`
	PieceSet      string `json:"pieceSet"`
	Theme3d       string `json:"theme3D"`
	PieceSet3d    string `json:"pieceSet3D"`
	SoundSet      string `json:"soundSet"`
	Blindfold     string `json:"blindfold"`
	AutoQueen     int    `json:"autoQueen"`
	AutoThreefold int    `json:"autoThreefold"`
	Takeback      int    `json:"takeback"`
	Moretime      int    `json:"moretime"`
	ClockTenths   int    `json:"clockTenths"`
	ClockBar      bool   `json:"clockBar"`
	ClockSound    bool   `json:"clockSound"`
	Premove       bool   `json:"premove"`
	Animation     int    `json:"animation"`
	Captured      bool   `json:"captured"`
	Follow        bool   `json:"follow"`
	Highlight     bool   `json:"highlight"`
	Destination   bool   `json:"destination"`
	Coords        int    `json:"coords"`
	Replay        int    `json:"replay"`
	Challenge     int    `json:"challenge"`
	Message       int    `json:"message"`
	CoordColor    int    `json:"coordColor"`
	SubmitMove    int    `json:"submitMove"`
	ConfirmResign int    `json:"confirmResign"`
	InsightShare  int    `json:"insightShare"`
	KeybpoardMove int    `json:"keybpoardMove"`
	Zen           int    `json:"zen"`
	MoveEvent     int    `json:"moveEvent"`
	RookCastle    int    `json:"rookCastle"`
}

type Preferences struct {
	Prefs    Prefs  `json:"prefs"`
	Language string `json:"language"`
}

type Kid struct {
	Kid bool `json:"kid"`
}
