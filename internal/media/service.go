package media

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type Service interface {
	UploadMedia(ctx context.Context, iFile multipart.File, iHandler *multipart.FileHeader) (string, error)
	GetPublicID(ctx context.Context, url string) (string, error)
	DeleteMedia(ctx context.Context, url string) error
}

type service struct {
}

var imgType = []string{"png", "jpg", "jpeg"}

// TODO : mover a .env
const (
	cloudname        = "dangvuvyq"
	cloudapikey      = "754218821349648"
	cloudinarysecret = "rjiWDoS5G0yNdiY4NZkEXtvit8k"
)

func NewService() Service {
	return &service{}
}

func (s *service) UploadMedia(ctx context.Context, iFile multipart.File, iHandler *multipart.FileHeader) (string, error) {
	cld, err := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
	if err != nil {
		return "", err
	}
	options := uploader.UploadParams{AllowedFormats: imgType}
	resp, err := cld.Upload.Upload(ctx, iFile, options)
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

func (s *service) DeleteMedia(ctx context.Context, url string) error {
	cld, err := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
	if err != nil {
		return err
	}
	publicId, err := s.GetPublicID(ctx, url)
	if err != nil {
		return errors.New("no se pudo obtener la publicId de la url ")
	}
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicId})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetPublicID(ctx context.Context, url string) (string, error) {
	firstUrl := strings.Split(url, ".")
	auxUrl := firstUrl[len(firstUrl)-2]
	secondUrl := strings.Split(auxUrl, "/")
	result := secondUrl[len(secondUrl)-1]
	return result, nil
}
