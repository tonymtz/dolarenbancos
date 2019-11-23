package banxico

type Response struct {
	Bmx *Bmx `json:"bmx"`
}

type Bmx struct {
	Series []Serie `json:"series"`
}

type Serie struct {
	Datos []Dato `json:"datos"`
}

type Dato struct {
	Dato string `json:"dato"`
}
