package user

type Perfs struct {
	Rating   int `json:"rating"`
	Progress int `json:"progress"`
}

type BulletPerfs struct {
	Bullet Perfs `json:"bullet"`
}
type BlitzPerfs struct {
	Blitz Perfs `json:"blitz"`
}
type RapidPerfs struct {
	Rapid Perfs `json:"rapid"`
}
type ClassicalPerfs struct {
	Classical Perfs `json:"classical"`
}
type UltraBulletPerfs struct {
	UltraBullet Perfs `json:"ultraBullet"`
}
type CrazyhousePerfs struct {
	Crazyhouse Perfs `json:"crazyhouse"`
}
type Chess960Perfs struct {
	Chess960 Perfs `json:"chess960"`
}
type KingOfTheHillPerfs struct {
	KingOfTheHill Perfs `json:"kingOfTheHill"`
}
type ThreeCheckPerfs struct {
	ThreeCheck Perfs `json:"threeCheck"`
}
type AntichessPerfs struct {
	Antichess Perfs `json:"antichess"`
}
type AtomicPerfs struct {
	Atomic Perfs `json:"atomic"`
}
type HordePerfs struct {
	Horde Perfs `json:"horde"`
}
type RacingKingsPerfs struct {
	RacingKings Perfs `json:"racingKings"`
}

type TopBulletPlayer struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Perfs    BulletPerfs `json:"perfs"`
	Title    string      `json:"title"`
	Patron   bool        `json:"patron"`
	Online   bool        `json:"online"`
}

type TopBlitzPlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Perfs    BlitzPerfs `json:"perfs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}

type TopRapidPlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Perfs    RapidPerfs `json:"perfs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}

type TopClassicalPlayer struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Perfs    ClassicalPerfs `json:"perfs"`
	Title    string         `json:"title"`
	Patron   bool           `json:"patron"`
	Online   bool           `json:"online"`
}

type TopUltraBulletPlayer struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Perfs    UltraBulletPerfs `json:"perfs"`
	Title    string           `json:"title"`
	Patron   bool             `json:"patron"`
	Online   bool             `json:"online"`
}

type TopCrazyHousePlayer struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Perfs    CrazyhousePerfs `json:"perfs"`
	Title    string          `json:"title"`
	Patron   bool            `json:"patron"`
	Online   bool            `json:"online"`
}

type TopChess960Player struct {
	ID       string        `json:"id"`
	Username string        `json:"username"`
	Perfs    Chess960Perfs `json:"perfs"`
	Title    string        `json:"title"`
	Patron   bool          `json:"patron"`
	Online   bool          `json:"online"`
}

type TopKingOfTheHillPlayer struct {
	ID       string             `json:"id"`
	Username string             `json:"username"`
	Perfs    KingOfTheHillPerfs `json:"perfs"`
	Title    string             `json:"title"`
	Patron   bool               `json:"patron"`
	Online   bool               `json:"online"`
}

type TopThreeCheckPlayer struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Perfs    ThreeCheckPerfs `json:"perfs"`
	Title    string          `json:"title"`
	Patron   bool            `json:"patron"`
	Online   bool            `json:"online"`
}

type TopAntichessPlayer struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Perfs    AntichessPerfs `json:"perfs"`
	Title    string         `json:"title"`
	Patron   bool           `json:"patron"`
	Online   bool           `json:"online"`
}

type TopAtomicPlayer struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Perfs    AtomicPerfs `json:"perfs"`
	Title    string      `json:"title"`
	Patron   bool        `json:"patron"`
	Online   bool        `json:"online"`
}
type TopHordePlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Perfs    HordePerfs `json:"perfs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}
type TopRacingKingsPlayer struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Perfs    RacingKingsPerfs `json:"perfs"`
	Title    string           `json:"title"`
	Patron   bool             `json:"patron"`
	Online   bool             `json:"online"`
}

type Top10 struct {
	Bullet        []TopBulletPlayer        `json:"bullet"`
	Blitz         []TopBlitzPlayer         `json:"blitz"`
	Rapid         []TopRapidPlayer         `json:"rapid"`
	Classical     []TopClassicalPlayer     `json:"classical"`
	UltraBullet   []TopUltraBulletPlayer   `json:"ultraBullet"`
	Chess960      []TopChess960Player      `json:"chess960"`
	Crazyhouse    []TopCrazyHousePlayer    `json:"crazyhouse"`
	Antichess     []TopAntichessPlayer     `json:"antichess"`
	Atomic        []TopAtomicPlayer        `json:"atomic"`
	Horde         []TopHordePlayer         `json:"horde"`
	KingOfTheHill []TopKingOfTheHillPlayer `json:"kingOfTheHill"`
	RacingKings   []TopRacingKingsPlayer   `json:"racingKings"`
	ThreeCheck    []TopThreeCheckPlayer    `json:"threeCheck"`
}
