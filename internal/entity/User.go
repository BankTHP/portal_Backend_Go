package entity

type Users struct {
    UserId     string `gorm:"primaryKey" json:"user_id"`
    Name       string `gorm:"type:varchar(255)" json:"name"`
    Username   string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
    GivenName  string `gorm:"type:varchar(255)" json:"given_name"`
    FamilyName string `gorm:"type:varchar(255)" json:"family_name"`
    Email      string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    UniversityName string `gorm:"type:varchar(255)" json:"universityName"`
	RegisterDate string `gorm:"type:varchar(255)" json:"registerDate"`
}

func (Users) TableName() string {
    return "users"
}
