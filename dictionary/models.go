// Package dictionary all models and operations needed to persist data
package dictionary

import (
	"github.com/meilisearch/meilisearch-go"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
  "os"
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

func SetupMeili() {
  client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey(os.Getenv("MEILI_KEY")))

	// An index is where the documents are stored.
	index := client.Index("movies")

	// If the index 'movies' does not exist, Meilisearch creates it when you first add the documents.
	documents := []map[string]interface{}{
		{"id": 1, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 2, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 3, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 4, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 5, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 6, "title": "Philadelphia", "genres": []string{"Drama"}},
	}
	task, err := index.AddDocuments(documents)
  failIfErr(err)

	fmt.Println(task.TaskUID)
}

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
