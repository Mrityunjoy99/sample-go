package url

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"

	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/repository"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
)

const (
	Host = "https://abc.com"
)

type Service interface {
	RegisterUrl(originalUrl string, ttlInSec int) (*RegisterUrlRespDto, genericerror.GenericError)
	Redirect(shortUrl string) (string, genericerror.GenericError)
}

type service struct {
	urlRepo repository.URLRepository
}

func NewService(urlRepo repository.URLRepository) Service {
	return &service{
		urlRepo: urlRepo,
	}
}

func (s *service) RegisterUrl(originalUrl string, ttlInSec int) (*RegisterUrlRespDto, genericerror.GenericError) {
	shortenedUrl, gerr := s.getShortenedUrl(originalUrl)
	if gerr != nil {
		return nil, gerr
	}

	urlEntity := entity.Url{
		OriginalURL: originalUrl,
		ShortURL:    shortenedUrl,
		ExpiredAt:   time.Now().Add((time.Second * time.Duration(ttlInSec))),
	}

	createdUrl, gerr := s.urlRepo.RegisterUrl(urlEntity)
	if gerr != nil {
		return nil, gerr
	}

	return &RegisterUrlRespDto{
		OriginalURL: originalUrl,
		ShortURL:    createdUrl.ShortURL,
		ExpiredAt:   createdUrl.ExpiredAt,
	}, nil
}

func (s *service) getShortenedUrl(originalUrl string) (string, genericerror.GenericError) {
	hasher := md5.New()
	_, err := io.WriteString(hasher, originalUrl)
	if err != nil {
		return "", genericerror.NewInternalErrByErr(err)
	}

	hash := hex.EncodeToString(hasher.Sum(nil))
	// Take first 6 characters of hash
	shortHash := hash[:6]

	return Host + "/" + shortHash, nil
}

func (s *service) Redirect(shortUrl string) (string, genericerror.GenericError) {
	return s.urlRepo.GetOriginalUrlByShortUrl(shortUrl)
}
