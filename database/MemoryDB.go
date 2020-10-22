package database

type Pokemon struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}

var PokemonDb = []Pokemon{
	'1': Pokemon{ID: "1", Name: "Pikachu", Type: "Electric"},
	'2': Pokemon{ID: "2", Name: "Charmeleon", Type: "Fire"},
}
