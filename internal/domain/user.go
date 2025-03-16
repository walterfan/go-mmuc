package domain

import (
	"github.com/google/uuid"
	"time"

)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Enabled     int       `gorm:"type:tinyint;default:1" json:"enabled"`
	CreatedBy   string    `json:"created_by"`
	CreatedTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_time"`
	ModifyBy    string    `json:"modify_by"`
	ModifyTime  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"modify_time"`
}
