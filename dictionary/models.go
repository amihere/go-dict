package dictionary

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type DefinitionModel struct {
  Id          string `json:"id"`
	Description string `json:"description"`
	Phonetic    string `json:"phonetic"`
	Name        string `json:"name"`
}

type Definition struct {
	id          int
	name        string
	description string
	phonetic    string
	refs        []int
	aliases     []int
}

func AddDefinition(name string) bool {
	return false
}

func GetDefinition(name string) (DefinitionModel, error) {
	rows, err := DB.Query("SELECT definition FROM words WHERE definition->>name LIKE 'ready'")

	def := DefinitionModel{}
	if err != nil {
		return def, err
	}

	rows.Scan(&def)
	return def, nil
}

func getConn() error {
	db, err := sql.Open("sqlite3", "./fante_dict.db")
	checkError(err)

	DB = db

  return nil
}

func SetupDatabase(init bool) {
	getConn()

	if !init {
		return
	}

	DB.Exec(`create table words (definition jsonb)`)

	stmt, err := DB.Prepare("insert into words(definition) values(?)")
	checkError(err)
	defer stmt.Close()

  def := DefinitionModel{Id : "796e9bd273244c4e5edabaad5bfc7b4", Name: "ready", Description: "Mentally disposed; willing m\u025Bk\u037B", Phonetic: "re_a_dyia"}
	marshalled, _ := json.Marshal(def)

	_, err = stmt.Exec(marshalled)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
