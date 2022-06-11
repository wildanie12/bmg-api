package web

// HeroAPIResponse mapper
type HeroAPIResponse struct {
	Type string `json:"type"`
	Format string `json:"format"`
	Version string `json:"version"`
	Data map[string]Hero `json:"data"`
}

// Hero main response mapper
type Hero struct {
	Version string `json:"version"`
	ID string `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	Title string `json:"title"`
	Blurb string `json:"blurb"`
	Info HeroInfo `json:"info"`
	Image HeroImage `json:"image"`
	Tags []string `json:"tags"`
	Partype string `json:"partype"`
	Stats HeroStats `json:"stats"`
}

// HeroInfo response mapper
type HeroInfo struct {
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	Magic int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

// HeroImage response mapper
type HeroImage struct {
	Full string `json:"full"`
	Sprite string `json:"sprite"`
	Group string `json:"group"`
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// HeroStats response mapper
type HeroStats struct {
	Hp float32 `json:"hp"`
	Hpperlevel float32 `json:"hpperlevel"`
	Mp float32 `json:"mp"`
	Mpperlevel float32 `json:"mpperlevel"`
	Movespeed float32 `json:"movespeed"`
	Armor float32 `json:"armor"`
	Armorperlevel float32 `json:"armorperlevel"`
	Spellblock float32 `json:"spellblock"`
	Spellblockperlevel float32 `json:"spellblockperlevel"`
	Attackrange float32 `json:"attackrange"`
	Hpregen float32 `json:"hpregen"`
	Hpregenperlevel float32 `json:"hpregenperlevel"`
	Mpregen float32 `json:"mpregen"`
	Mpregenperlevel float32 `json:"mpregenperlevel"`
	Crit float32 `json:"crit"`
	Critperlevel float32 `json:"critperlevel"`
	Attackdamage float32 `json:"attackdamage"`
	Attackdamageperlevel float32 `json:"attackdamageperlevel"`
	Attackspeedoffset float32 `json:"attackspeedoffset"`
	Attackspeedperlevel float32 `json:"attackspeedperlevel"`
}