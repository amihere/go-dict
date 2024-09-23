package dictionary

type Definition struct {
  id int
  Description string `json:"description"`
  Phonetic string `json:"phonetic"`
  Name string`json:"name"`
}
