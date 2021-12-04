package interfaces

type ICsvHandler interface {
	Read() ([][]string, error)
	Append(row []string) error
}
