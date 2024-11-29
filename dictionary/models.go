// Package dictionary all models and operations needed to persist data
package dictionary

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"log"
	"os"
	"time"
)

// DB is a SQL pool variable
var DB *sql.DB

// DefinitionModel - model object for word in dictionary
type DefinitionModel struct {
	ID          string   `json:"id"`
	Updated     int      `json:"updated"`
	Description string   `json:"description"`
	Phonetic    string   `json:"phonetic"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
}

type definitionModelIndex struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type definition struct {
	id          int
	name        string
	description string
	phonetic    string
	refs        []int
	aliases     []int
}

func getIndexName() string {
	return "definition-model"
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
	fmt.Println(name)
	// find the name using meili

	id := "796e9bd273244c4e5edabaad5bfc7b4"
	return getDefinition(id)
}

func getDefinition(id string) (DefinitionModel, error) {
	var def DefinitionModel

	err := DB.QueryRow("SELECT definition FROM words WHERE definition->>'id' = ?", id).Scan(&def)

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

	def := DefinitionModel{ID: "796e9bd273244c4e5edabaad5bfc7b4", Updated: int(time.Now().UTC().UnixMilli()), Name: "ready", Description: "Mentally disposed; willing m\u025Bk\u037B", Phonetic: "re_a_dyia"}
	marshalled, _ := json.Marshal(def)
	_, err = stmt.Exec(string(marshalled))

	failIfErr(err)
}

// SetupMeili should init meili driver pool
func SetupMeili() {
	meiliKey := os.Getenv("MEILI_KEY")
	search := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey(meiliKey))

	// An index is where the documents are stored.
	index := search.Index(getIndexName())

	// If the index 'movies' does not exist, Meilisearch creates it when you first add the documents.
	documents := []definitionModelIndex{
		{ID: "796e9bd273244c4e5edabaad5bfc7b4", Name: "ready", Tags: []string{"Fantasy", "Action"}},
		{ID: "6", Name: "Philadelphia", Tags: []string{"Drama"}},
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
