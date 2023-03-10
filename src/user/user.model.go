package user

type User struct {
	ID        int    `gorm:"autoIncrement:2;unique;primaryKey" json:"id"`
	FirstName string `gorm:"type: varchar(50);not null" json:"firstname"`
	LastName  string `gorm:"type: varchar(50);not null" json:"lastname"`
	Age       int    `gorm:"type: int;not null" json:"age"`
}
