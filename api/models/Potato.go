package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Potato struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Potato) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Potato) SavePotato(db *gorm.DB) (*Potato, error) {
	var err error
	err = db.Debug().Model(&Potato{}).Create(&p).Error
	if err != nil {
		return &Potato{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Potato{}, err
		}
	}
	return p, nil
}

func (p *Potato) FindAllPotatoes(db *gorm.DB) (*[]Potato, error) {
	var err error
	potatoes := []Potato{}
	err = db.Debug().Model(&Potato{}).Limit(100).Find(&potatoes).Error
	if err != nil {
		return &[]Potato{}, err
	}

	return &potatoes, nil
}

func (p *Potato) FindPotatoByID(db *gorm.DB, pid uint64) (*Potato, error) {
	var err error
	err = db.Debug().Model(&Potato{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Potato{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Potato{}, err
		}
	}
	return p, nil
}

func (p *Potato) UpdateAPotato(db *gorm.DB) (*Potato, error) {

	var err error

	err = db.Debug().Model(&Potato{}).Where("id = ?", p.ID).Updates(Potato{Name: p.Name, Description: p.Description, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Potato{}, err
	}
	if p.ID != 0 {
		if err != nil {
			return &Potato{}, err
		}
	}
	return p, nil
}

func (p *Potato) DeleteAPotato(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Potato{}).Where("id = ?", pid).Take(&Potato{}).Delete(&Potato{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Potato not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
