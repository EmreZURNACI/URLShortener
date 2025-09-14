package infra

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/EmreZURNACI/url-shortener/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func Connection() (*Handler, error) {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&tls=false",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, ErrConnectionFailed
	}

	con, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, ErrOrmConnectionFailed
	}

	return &Handler{
		db: con,
	}, nil
}

func (h *Handler) CreateTable() error {

	if err := h.db.AutoMigrate(&domain.Address{}); err != nil {
		return ErrMigrateFailed
	}
	return nil

}
func (h *Handler) GetURL(ctx context.Context, address domain.Address) (*domain.Address, error) {

	var link domain.Address

	err := h.db.WithContext(ctx).Model(&domain.Address{}).Where("url = ?", address.URL).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrQueryFailed
	}
	return &link, nil
}
func (h *Handler) GetShortURL(ctx context.Context, address domain.Address) (*domain.Address, error) {

	var link domain.Address

	err := h.db.WithContext(ctx).Model(&domain.Address{}).Where("short_url = ?", address.ShortURL).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrQueryFailed
	}
	return &link, nil
}
func (h *Handler) CreateURL(ctx context.Context, address domain.Address) (*string, error) {

	tx := h.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, ErrTransactionFailed
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&domain.Address{}).Create(&address).Error; err != nil {
		return nil, ErrQueryFailed
	}

	if err := tx.Commit().Error; err != nil {
		return nil, ErrCommitFailed
	}

	if err := h.db.WithContext(ctx).Where("url = ?", address.URL).First(&address).Error; err != nil {
		return nil, ErrQueryFailed

	}

	return &address.ShortURL, nil

}
