package prediction

type Winner struct {
	Name    string `bson:"name"`
	Comment string `bson:"comment"`
}

type Percent struct {
	Home string `bson:"home"`
	Draw string `bason:"draw"`
	Away string `bason:"away"`
}

type GoalString struct {
	Home string `bson:"home"`
	Away string `bason:"away"`
}

type Prediction struct {
	Advice     string     `bson:"advice"`
	WinOrDraw  bool       `bson:"win_or_draw"`
	Underover  string     `bson:"under_over"`
	Winner     Winner     `bson:"winner"`
	Percent    Percent    `bson:"percent"`
	GoalString GoalString `bson:"goals"`
}

type League struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Logo    string `bson:"logo"`
	Flag    string `bson:"flag"`
	Season  string `bson:"season"`
}

type TeamType struct {
	Name   string `bason:"name"`
	Logo   string `bason:"logo"`
	Winner bool   `bason:"winner"`
}

type Teams struct {
	Home TeamType `bason:"home"`
	Away TeamType `bason:"away"`
}

type PredictionResponse struct {
	Prediction Prediction `bson:"predictions"`
	League     League     `bson:"league"`
	Teams      Teams      `bson:"teams"`
	FixtureId  int64      `bson:"fixtureId"`
}
