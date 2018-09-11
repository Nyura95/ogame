package ogame

type MoonInfos struct {
	ID       int
	Diameter int
}

type AllianceInfos struct {
	ID     int
	Name   string
	Rank   int
	Member int
}

// PlanetInfos ...
type PlanetInfos struct {
	ID              int
	Activity        int
	Name            string
	Img             string
	Coordinate      Coordinate
	Administrator   bool
	Inactive        bool
	Vacation        bool
	StrongPlayer    bool
	HonorableTarget bool
	Debris          struct {
		Metal           int
		Crystal         int
		RecyclersNeeded int
	}
	Moon   *MoonInfos
	Player struct {
		ID   int
		Name string
		Rank int
	}
	Alliance *AllianceInfos
}