package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
)

type User struct {
	Base
	Email     string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null" json:"-"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Telephone string `gorm:"size:255;not null" json:"telephone"`
	Balance   int64  `gorm:"default:0" json:"balance"`
}

func (u *User) CheckPassword(pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetBalanceFormatted By default we store the Balance as integer 60, wihout decimal,
// multiplying the real balance by 100 (2 decimals)
func (u *User) GetBalanceFormatted() string {
	b := float64(u.Balance) / 100.00
	return fmt.Sprintf("%.2f", b)
}

func (u *User) TopUpFromString(amount string) error {
	fa, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}
	ia, _ := math.Modf(fa * 100.00)
	u.Balance = int64(ia) + u.Balance

	return nil
}
