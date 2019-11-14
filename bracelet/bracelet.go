package bracelet

import (
	"golang-starter-pack/model"
)

// Interface

type Store interface {
	CreateBracelet(*model.Bracelet) error
	GetBySlug(string) (*model.Bracelet, error)
	DeleteBracelet(*model.Bracelet) error

	AddBead(*model.Bracelet, *model.Bead) error
	GetBeadsBySlug(string) ([]model.Bead, error)
	DeleteBead(*model.Bead) error
}
