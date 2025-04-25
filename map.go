package main

type HeroeslandMap []struct {
	M  string `json:"m"`
	A  string `json:"a"`
	ID int    `json:"id"`
	P  struct {
		Twalk int `json:"twalk"`
		M     struct {
			Mid   int  `json:"mid"`
			Pid   int  `json:"pid"`
			Mw    int  `json:"mw"`
			Mh    int  `json:"mh"`
			X     int  `json:"x"`
			Y     int  `json:"y"`
			W     int  `json:"w"`
			H     int  `json:"h"`
			R     int  `json:"r"`
			Piece bool `json:"piece"`
			Fow   bool `json:"fow"`
		} `json:"m"`
		D []struct {
			ID  int    `json:"id"`
			Tid int    `json:"tid"`
			Fid int    `json:"fid"`
			X   int    `json:"x"`
			Y   int    `json:"y"`
			T   int    `json:"t"`
			Dx  int    `json:"dx"`
			Dy  int    `json:"dy"`
			Ft  int    `json:"ft"`
			Et  int    `json:"et"`
			Dir int    `json:"dir"`
			L   int    `json:"l"`
			N   string `json:"n"`
			M   string `json:"m"`
			Tw  int    `json:"tw"`
		} `json:"d"`
		P struct {
			Cells map[string]int `json:"cells"`
			Clans []struct {
				ID int    `json:"id"`
				N  string `json:"n"`
			} `json:"clans"`
		} `json:"p"`
	} `json:"p"`
}
