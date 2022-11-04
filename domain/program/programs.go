package program

import (
	"time"

	"github.com/VooDooStack/FitStackAPI/domain/routine"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"gorm.io/gorm"
)

type Program struct {
	ID          uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	CreatorID   uint             `json:"creator_id"`
	Creator     *user.User       `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Routine     *routine.Routine `json:"routine" gorm:"foreignKey:ID"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index"`
}

type ProgramUsecase interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) ([]*Program, error)
	Create(program *Program) error
	Update(program *Program) error
}

type ProgramRepository interface {
	GetById(uuid string) (*Program, error)
	GetByCreator(creatorId string) (*Program, error)
	Get(uuid string) ([]*Program, error)
	Create(program *Program) error
	Update(program *Program) error
}
