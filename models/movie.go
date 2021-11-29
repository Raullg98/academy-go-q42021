package models

import (
	"errors"
	"strings"

	"go_project/util"
)

type Movie struct {
	ID     string
	Title  string
	Genres []string
}

// ParseRecords - Parses the movies csv file
func ParseRecords() ([]Movie, error) {
	records, err := util.ReadCsvFile("./data/movies.csv")
	if err != nil {
		return make([]Movie, 0), err
	}

	data := make([]Movie, len(records[0]))

	for _, record := range records {
		movie := Movie{
			ID:     record[0],
			Title:  record[1],
			Genres: strings.Split(record[2], "|"),
		}

		data = append(data, movie)
	}

	return data, nil
}

// GetAllMovies - Returns all movies
func GetAllMovies() ([]Movie, error) {
	movies, err := ParseRecords()
	if err != nil {
		return make([]Movie, 0), err
	}

	return movies, nil
}

// GetMovieById - Returns a single movie based on their id
func GetMovieById(id string) (*Movie, error) {
	var movies, err = ParseRecords()
	if err != nil {
		return nil, err
	}

	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}

	return nil, errors.New("Not found")
}
