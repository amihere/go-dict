package dictionary

type DefinitionModel struct {
	id          int
	Description string `json:"description"`
	Phonetic    string `json:"phonetic"`
	Name        string `json:"name"`
}
