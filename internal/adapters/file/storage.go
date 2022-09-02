package file

import (
	"bytes"
	"context"
	"github.com/rs/zerolog"
	"log"
	"os"
)

type Database struct {
	file   *os.File
	logger *zerolog.Logger
}

func New(ctx context.Context, l *zerolog.Logger, fileName string) (*Database, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	db := &Database{
		file:   f,
		logger: l,
	}

	return db, nil
}

func (db *Database) Read() []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(db.file)

	return buf.Bytes()
}

func (db *Database) Append(data []byte) (err error) {
	_, err = db.file.Write(data)

	return
}

func (db *Database) Truncate() (err error) {
	fName := db.file.Name()
	db.file.Close()

	db.file, err = os.Create(fName)
	if err != nil {
		return err
	}

	return
}

func (db *Database) Close() {
	err := db.file.Close()
	if err != nil {
		db.logger.Error().Err(err)
	}
}
