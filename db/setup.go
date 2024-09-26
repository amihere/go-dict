package db

func initDb()  {
  const tableInit string = `
    CREATE TABLE IF NOT EXISTS Definition(
      id INTEGER NOT NULL PRIMARY KEY,
      time DATETIME NOT NULL,
      description TEXT
    );`


}
