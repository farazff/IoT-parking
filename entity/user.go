package entity

type User interface {
	ID() int
	FirstName() string
	LastName() string
	CarTag() string
	Phone() string
	Password() string
}

type UserUpdater struct {
	ID          int    `json:"ID" db:"id"`
	FirstName   string `json:"first_name" db:"first_name" validate:"required"`
	LastName    string `json:"last_name" db:"last_name" validate:"required"`
	CarTag      string `json:"car_tag" db:"car_tag" validate:"required"`
	OldPassword string `json:"old_password" db:"password"`
	NewPassword string `json:"new_password" db:"first_name"`
}
