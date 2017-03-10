package data

type Storage interface {
	GetUploadedImages() ([]UploadedFragment, error)
}
