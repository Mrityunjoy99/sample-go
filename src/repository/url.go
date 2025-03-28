package repository

import (
	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"gorm.io/gorm"
)

type URLRepository interface {
	RegisterUrl(url entity.Url) (*entity.Url, genericerror.GenericError)
	GetOriginalUrlByShortUrl(shortUrl string) (string, genericerror.GenericError)
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) URLRepository {
	return &urlRepository{
		db: db,
	}
}

func (r *urlRepository) RegisterUrl(url entity.Url) (*entity.Url, genericerror.GenericError) {
	err := r.db.Create(&url).Error
	if err != nil {
		return nil, genericerror.NewInternalErrByErr(err)
	}

	return &url, nil
}

func (r *urlRepository) GetOriginalUrlByShortUrl(shortUrl string) (string, genericerror.GenericError) {
	var url entity.Url

	result := r.db.Model(&url).Where("short_url=?", shortUrl).First(&url)
	if result.RowsAffected == 0 {
		return "", genericerror.NewGenericError(constant.ErrorCodeResourceNotFound, "url does not exists", nil, nil)
	}

	if result.Error != nil {
		return "", genericerror.NewInternalErrByErr(result.Error)
	}

	return url.OriginalURL, nil
}
