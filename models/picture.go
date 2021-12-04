package models

// Picture - struct for picture
type Picture struct {
	Date        string `json:"date"`
	Copyright   string `json:"copyright"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Url         string `json:"hdurl"`
	MediaType   string `json:"media_type"`
}

// Values - Return a string array containing all the values
func (p Picture) Values() []string {
	return []string{
		p.Date,
		p.Copyright,
		p.Title,
		p.Explanation,
		p.Url,
		p.MediaType,
	}
}

// Values - Returns a picture based on an array of values
func NewPictureFromValues(values []string) *Picture {
	return &Picture{
		Date:        values[0],
		Copyright:   values[1],
		Title:       values[2],
		Explanation: values[3],
		Url:         values[4],
		MediaType:   values[5],
	}
}
