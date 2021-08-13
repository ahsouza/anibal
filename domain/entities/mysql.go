package entities

type (
	User struct {
		ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
		FirstName string `gorm:"type:varchar(255)" json:"first_name"`
		LastName  string `gorm:"type:varchar(255)" json:"last_name"`
		Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
		Avatar    string `gorm:"not null;default:"`
		Password  string `gorm:"->;<-;not null" json:"-"`
		Token     string `gorm:"-" binding:"-" json:"token,omitempty"`
	}

	Book struct {
		ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
		Title       string `gorm:"type:varchar(255)" json:"title"`
		Description string `gorm:"type:text" json:"description"`
		Author      string `gorm:"type:varchar(255)" json:"author"`
		Photo       string `gorm:"not null;default:"`
	}

	Movie struct {
		ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
		Title       string `gorm:"type:varchar(255)" json:"title"`
		Description string `gorm:"type:text" json:"description"`
		Director    string `gorm:"type:varchar(255)" json:"director"`
		Imdb        string `gorm:"type:varchar(255)" json:"imdb"`
		Photo       string `gorm:"not null;default:"`
	}
)
