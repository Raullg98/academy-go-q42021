package repositories

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"

	"gopkg.in/resty.v1"

	"go_project/infrastructure"
	"go_project/interfaces"
	"go_project/models"
)

type PictureRepo struct {
	client *resty.Request
	csv    interfaces.ICsvHandler
}

// NewPictureRepository - Creates a new Picture Repository
func NewPictureRepository() *PictureRepo {
	return &PictureRepo{
		client: resty.R().SetQueryParam("api_key", os.Getenv("API_KEY")),
		csv:    infrastructure.NewCsvHandler("./data/pictures.csv"),
	}
}

func worker(workerId, maxJobsPerWorker int, jobs <-chan []string, results chan<- *models.Picture) {
	fmt.Printf("Worker %d spawned\n", workerId)
	processedPictures := 0
	for job := range jobs { // you must check for readable state of the channel.
		results <- models.NewPictureFromValues(job)

		processedPictures++
		if processedPictures >= maxJobsPerWorker {
			break
		}

	}
	fmt.Printf("Worker %d finished\n", workerId)
}

func readFileWorkerPool(f *os.File, items int, criteria string, jobs chan []string) {
	csvReader := csv.NewReader(f)
	processedPictures := 0
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}

		t, err := time.Parse("2006-01-02", row[0])
		day := t.Day()

		if criteria == "Even" && day%2 == 0 {
			jobs <- row
			processedPictures++

		} else if criteria == "Odd" && day%2 != 0 {
			jobs <- row
			processedPictures++

		}

		if processedPictures >= items {
			break
		}

	}
	close(jobs)
}

// ReadPicturesWithWorkerPool - Returns a list of pictures based on criteria, items using a Worker pool limited by maxJobsPerWorker param
func (pr *PictureRepo) ReadPicturesWithWorkerPool(criteria string, items int, maxJobsPerWorker int) ([]models.Picture, error) {
	f, err := os.Open("./data/pictures.csv")
	if err != nil {
		return nil, err
	}

	jobs := make(chan []string)
	results := make(chan *models.Picture)

	var wg sync.WaitGroup

	numOfWorkers := int(math.Ceil(float64(items) / float64(maxJobsPerWorker)))

	wg.Add(numOfWorkers)
	for w := 0; w < numOfWorkers; w++ {
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, maxJobsPerWorker, jobs, results)
		}(w)
	}

	go readFileWorkerPool(f, items, criteria, jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	pictures := make([]models.Picture, 0)
	for picture := range results {
		pictures = append(pictures, *picture)
	}

	return pictures, nil
}

// FindPictureByDate - Fetches the picture from the specific date
func (pr *PictureRepo) FindPictureByDate(date string) (*models.Picture, error) {
	client := pr.client.
		SetError(&models.NasaErrorResponse{}).
		SetResult(&models.Picture{})

	resp, err := client.
		SetQueryParam("date", date).
		Get("https://api.nasa.gov/planetary/apod")

	if err != nil {
		return nil, err
	} else if resp.IsError() {
		responseError := resp.Error().(*models.NasaErrorResponse)
		return nil, responseError.GetError()
	}

	picture := resp.Result().(*models.Picture)
	return picture, nil
}

// SavePicture - Appends the picture to the csv
func (pr *PictureRepo) SavePicture(picture *models.Picture) error {

	return pr.csv.Append(picture.Values())
}
