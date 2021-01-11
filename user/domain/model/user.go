package model

type User struct {
	//PK
	ID           string `mapstructure:"id" gorm:"user_id;primary_key;not_null;auto_increment"`
	UserName     string `mapstructure:"username" gorm:"user_name;unique_index; not_null"`
	FirstName    string `mapstructure:"firstname" gorm:"first_name;not_null"`
	LastName     string `mapstructure:"lastname" gorm:"last_name;not_null"`
	HashPassword string `mapstructure:"haspassword"`
}
