package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Noun struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Noun      string    `gorm:"size:255;not null" json:"noun"`
	Gender    string    `gorm:"size:1;not null;" json:"gender"`
	Meaning   string    `gorm:"size:255;not null;" json:"meaning"`
	Plural    string    `gorm:"size:255;" json:"plural"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *Noun) Prepare() {
	n.ID = 0
	n.Noun = html.EscapeString(strings.TrimSpace(n.Noun))
	n.Gender = html.EscapeString(strings.TrimSpace(n.Gender))
	n.Meaning = html.EscapeString(strings.TrimSpace(n.Meaning))
	n.Plural = html.EscapeString(strings.TrimSpace(n.Plural))
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
}

func (n *Noun) Validate() error {

	if n.Noun == "" {
		return errors.New("Required Noun")
	}
	if n.Gender == "" {
		return errors.New("Required Gender")
	}
	if n.Gender == "" {
		return errors.New("Required Meaning")
	}
	return nil
}

func (n *Noun) SaveNoun(db *gorm.DB) (*Noun, error) {
	var err error
	err = db.Debug().Model(&Noun{}).Create(&n).Error
	if err != nil {
		return &Noun{}, err
	}
	// if n.ID != 0 {
	// 	err = db.Debug().Model(&Noun{}).Where("id = ?", n.AuthorID).Take(&n.Author).Error
	// 	if err != nil {
	// 		return &Noun{}, err
	// 	}
	// }
	return n, nil
}

func (n *Noun) FindAllNouns(db *gorm.DB) (*[]Noun, error) {
	var err error
	nouns := []Noun{}
	// err = db.Debug().Model(&Noun{}).Limit(100).Find(&nouns).Error
	err = db.Debug().Model(&Noun{}).Find(&nouns).Error
	if err != nil {
		return &[]Noun{}, err
	}
	// if len(nouns) > 0 {
	// 	for i, _ := range nouns {
	// 		err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
	// 		if err != nil {
	// 			return &[]Noun{}, err
	// 		}
	// 	}
	// }
	return &nouns, nil
}

func (n *Noun) FindNounByID(db *gorm.DB, pid uint64) (*Noun, error) {
	var err error
	err = db.Debug().Model(&Noun{}).Where("id = ?", pid).Take(&n).Error
	if err != nil {
		return &Noun{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Noun{}, err
	// 	}
	// }
	return n, nil
}

func (n *Noun) UpdateANoun(db *gorm.DB) (*Noun, error) {
	var err error
	err = db.Debug().Model(&Noun{}).Where("id = ?", n.ID).Updates(Noun{Noun: n.Noun, Gender: n.Gender, Meaning: n.Meaning, Plural: n.Plural, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Noun{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Noun{}, err
	// 	}
	// }
	return n, nil
}

func (n *Noun) DeleteANoun(db *gorm.DB, pid uint64) (int64, error) {

	// db = db.Debug().Model(&Noun{}).Where("id = ? and author_id = ?", pid, uid).Take(&Noun{}).Delete(&Noun{})
	db = db.Debug().Model(&Noun{}).Where("id = ?", pid).Take(&Noun{}).Delete(&Noun{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Noun not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
