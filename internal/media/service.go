package media

import (
	"context"
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type Service interface {
	UploadMedia(ctx context.Context, iFile multipart.File, iHandler *multipart.FileHeader) (string, error)
	GetPublicID(ctx context.Context, url string) (string, error)
	DeleteMedia(ctx context.Context, url string) error
	UploadByUrl(ctx context.Context, url string) (string, error)
}

type service struct {
}

var imgType = []string{"png", "jpg", "jpeg"}

var (
	cloudname        = ""
	cloudapikey      = ""
	cloudinarysecret = ""
)

func NewService() Service {
	return &service{}
}

func setKeys() {
	cloudname = os.Getenv("CLOUD_NAME")
	cloudapikey = os.Getenv("CLOUD_API_KEY")
	cloudinarysecret = os.Getenv("CLOUDINARY_SECRET")
}

func (s *service) UploadMedia(ctx context.Context, iFile multipart.File, iHandler *multipart.FileHeader) (string, error) {
	setKeys()
	cld, err := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
	if err != nil {
		return "", err
	}
	//https://res.cloudinary.com/dangvuvyq/image/upload/c_fill,g_auto,h_199,w_800/v1662473607/fmb5w8ya5dgkezanzgx9.png
	options := uploader.UploadParams{AllowedFormats: imgType}
	resp, err := cld.Upload.Upload(ctx, iFile, options)
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

func (s *service) UploadByUrl(ctx context.Context, url string) (string, error) {
	setKeys()
	cld, err := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
	if err != nil {
		return "", err
	}
	options := uploader.UploadParams{AllowedFormats: imgType}
	resp, err := cld.Upload.Upload(ctx, url, options)
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

func (s *service) DeleteMedia(ctx context.Context, url string) error {
	setKeys()
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
	result := ""
	auxUrl := ""
	first := strings.Split(url, "/")
	if len(first) > 1 {
		auxUrl = first[len(first)-1]
	}
	second := strings.Split(auxUrl, ".")
	if len(second) > 0 {
		result = second[0]
	}
	return result, nil
}
