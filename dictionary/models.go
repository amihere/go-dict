// Package dictionary all models and operations needed to persist data
package dictionary

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
)

// DB is a SQL pool variable
var DB *sql.DB

type DefinitionModel struct {
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Phonetic    string   `json:"phonetic"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
}

type Definition struct {
	id          int
	name        string
	description string
	phonetic    string
	refs        []int
	aliases     []int
}

func (t *DefinitionModel) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), t)
}

func (t *DefinitionModel) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}

func AddDefinition(name string) bool {
	return false
}

// GetDefinition finds a word using name for fuzzy find
func GetDefinition(name string) (DefinitionModel, error) {
	var def DefinitionModel

	name += "%"
	fmt.Println(name)
	err := DB.QueryRow("SELECT definition FROM words WHERE definition->>'name' LIKE ?", name).Scan(&def)

	if err != nil {
		fmt.Println(err)
	}
	return def, err
}

func getConn() error {
	db, err := sql.Open("sqlite3", "./fante_dict.db")
	failIfErr(err)

	DB = db

	return nil
}

// SetupDatabase setup scripts for database connection and/or init
func SetupDatabase(init bool) {
	getConn()

	if !init {
		return
	}

	DB.Exec(`create table words (definition jsonb)`)

	stmt, err := DB.Prepare("insert into words(definition) values(?)")
	failIfErr(err)
	defer stmt.Close()

	def := DefinitionModel{Id: "796e9bd273244c4e5edabaad5bfc7b4", Name: "ready", Description: "Mentally disposed; willing m\u025Bk\u037B", Phonetic: "re_a_dyia"}
	marshalled, _ := json.Marshal(def)
	_, err = stmt.Exec(string(marshalled))

	failIfErr(err)
}

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
