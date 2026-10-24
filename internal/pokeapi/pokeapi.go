package pokeapi

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

func GetLocationAreaUrl() *string {
	url := baseUrl + "/location-area/"
	return &url
}
