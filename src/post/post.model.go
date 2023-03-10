package post

type Post struct {
	ID          int    `gorm:"autoIncrement:1;unique;primaryKey" json:"id"`
	Title       string `gorm:"type: varchar(200);not null;unique" json:"title"`
	Description string `gorm:"type: text;not null" json:"description"`
}
