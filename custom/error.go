package custom

type Error struct {
	Message    string `json:"messageError"`
	ErrorField string `json:"errorField"`
	Field      string `json:"field"`
}
