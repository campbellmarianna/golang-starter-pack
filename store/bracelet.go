package store

import (
	"github.com/jinzhu/gorm"
	"golang-starter-pack/model"
)

type BraceletStore struct {
	db *gorm.DB
}

func NewBraceletStore(db *gorm.DB) *BraceletStore {
	return &BraceletStore{
		db: db,
	}
}

// Bracelet Funtions
func (bs *BraceletStore) CreateBracelet(a *model.Bracelet) error {
	tx := bs.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		return err
	}
}

func (bs *BraceletStore) GetBySlug(s string) (*model.Bracelet, error) {
	var m model.Bracelet
	err := bs.db.Where(&model.Bracelet{Slug: s}).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (bs *BraceletStore) DeleteBracelet(a *model.Bracelet) error {
	return bs.db.Delete(a).Error
}

// Bead Functions
func (bs *BraceletStore) AddBracelet(a *model.Bracelet, c *model.Bead) error {
	err := bs.db.Model(a).Association("Beads").Append(c).Error
	if err != nil {
		return err
	}
	return bs.db.Where(c.ID).First(c).Error
}

func (bs *BraceletStore) GetBeadsBySlug(slug string) ([]model.Bead, error) {
	var m model.Bracelet
	if err := bs.db.Where(&model.Bracelet{Slug: slug}).Preload("Beads").First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return m.Beads, nil
}

func (as *BraceletStore) GetBeadByID(id uint) (*model.Bead, error) {
	var m model.Bead
	if err := as.db.Where(id).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}


func (bs *BraceletStore) DeleteBead(b *model.Bead) error {
	return bs.db.Delete(b).Error
}