package entities

type Movie struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Director    string `gorm:"type:varchar(255)" json:"director"`
	Imdb        string `gorm:"type:varchar(255)" json:"imdb"`
	Photo       string `gorm:"not null;default:"`
}
