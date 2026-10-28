package pokeapi

type IndividualPokemon struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Species        Species `json:"species"`
	Types          []Types `json:"types"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
