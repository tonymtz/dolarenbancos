package banks

import "github.com/tonymtz/dolarenbancos/models"

var Banxico = models.Bank{
	Id:   0,
	Name: "Banxico",
}

var Inbursa = models.Bank{
	Id:   1,
	Name: "Inbursa",
}

var Banamex = models.Bank{
	Id:   2,
	Name: "Banamex",
}

var Santander = models.Bank{
	Id:   3,
	Name: "Santander",
}

var Bbva = models.Bank{
	Id:   4,
	Name: "BBVA",
}

var Banorte = models.Bank{
	Id:   5,
	Name: "Banorte",
}

var Banks = []models.Bank{
	Banxico, Inbursa, Banamex, Santander, Bbva, Banorte,
}
