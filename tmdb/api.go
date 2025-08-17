package tmdb

import (
	"errors"
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

type TMDB struct {
	client *tmdb.Client
}

func (t *TMDB) SearchMovie(query string) {}

func Init() (*TMDB, error) {
	key := os.Getenv("TMDB_KEY")
	if key == "" {
		return nil, errors.New("no API key found in environment ('TMDB_KEY')")
	}
	client, err := tmdb.Init(key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize TMDB API: %w", err)
	}
	return &TMDB{client}, nil
}
