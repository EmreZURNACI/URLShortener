package domain

type Address struct {
	UUID     string `json:"uuid" gorm:"primaryKey;type:varchar(36)"`
	URL      string `json:"url" validate:"required" gorm:"not null;type:text"`
	ShortURL string `json:"short_url" gorm:"not null;type:varchar(10);unique"`
}
