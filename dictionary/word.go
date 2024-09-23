package dictionary

type Definition struct {
  Id int
  Description string `json:"description"`
  Phonetic string `json:"phonetic"`
  Name string`json:"name"`
}
