package models

type Song struct {
	ID          uint64 `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name;index;size:130"`
	Group       string `gorm:"column:group;index;size:130"`
	ReleaseDate string `gorm:"column:release_date;size:10"`
	Text        string `gorm:"column:text;type:text"`
	Link        string `gorm:"column:link;size:150"`
}
