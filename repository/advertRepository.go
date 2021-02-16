package repository

import (
	"gorm.io/gorm"
	"rest/models"
)

type AdvertRepository struct {
	Database *gorm.DB
}

func (ar *AdvertRepository) FindAll() ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)
	err := ar.Database.Table("advert").Find(&adverts).Error
	if err != nil {
		return nil, err
	}
	return adverts, nil
}

func NewAdvertRepository(db *gorm.DB) *AdvertRepository {
	return &AdvertRepository{db}
}
