package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID  `json:"id" db:"id" redis:"id" validate:"omitempty"`
	UserName    string     `json:"user_name" db:"user_name" redis:"user_name" validate:"required,lte=30"`
	Email       string     `json:"email,omitempty" db:"email" redis:"email" validate:"required,lte=60,email"`
	Password    string     `json:"password,omitempty" db:"password" redis:"password" validate:"omitempty,required,gte=6"`
	Role        *string    `json:"role,omitempty" db:"role" redis:"role" validate:"omitempty,lte=10"`
	About       *string    `json:"about,omitempty" db:"about" redis:"about" validate:"omitempty,lte=1024"`
	Avatar      *string    `json:"avatar,omitempty" db:"avatar" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number" validate:"omitempty,lte=20"`
	Address     *string    `json:"address,omitempty" db:"address" redis:"address" validate:"omitempty,lte=250"`
	City        *string    `json:"city,omitempty" db:"city" redis:"city" validate:"omitempty,lte=24"`
	Country     *string    `json:"country,omitempty" db:"country" redis:"country" validate:"omitempty,lte=24"`
	Gender      *string    `json:"gender,omitempty" db:"gender" redis:"gender" validate:"omitempty,lte=10"`
	Postcode    *int       `json:"postcode,omitempty" db:"postcode" redis:"postcode" validate:"omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty" db:"birthday" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
	LoginDate   time.Time  `json:"login_date" db:"login_date" redis:"login_date"`
}

type RegisterUserRequest struct {
	UserName string  `json:"user_name" form:"user_name" binding:"required"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required"`
	Role     *string `json:"role" binding:"required"`
	Gender   *string `json:"gender" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
