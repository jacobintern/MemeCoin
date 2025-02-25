package repository

import (
	"database/sql"
	"time"

	"github.com/jacobintern/meme_coin/service/adapter_model.go"
)

type memeCoin struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement"`
	Name            string         `gorm:"size:255;unique;not null"`
	Description     sql.NullString `gorm:"type:text"`
	CreatedAt       time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	PopularityScore uint           `gorm:"not null;default:0"`
}

func (*memeCoin) TableName() string {
	return "meme_coins"
}

func newMemeCoin() memeCoin {
	return memeCoin{}
}

func (m memeCoin) toEntity() (e adapter_model.MemeCoinInfo) {
	e.ID = m.ID
	e.Name = m.Name
	if m.Description.Valid {
		e.Description = &m.Description.String
	}
	e.CreatedAt = m.CreatedAt
	e.PopularityScore = m.PopularityScore
	return e
}

func (m memeCoin) initModel(e adapter_model.MemeCoinInfo) memeCoin {
	m.Name = e.Name
	if e.Description != nil {
		m.Description.Valid = true
		m.Description.String = *e.Description
	}
	return m
}

func (m memeCoin) toModel(e adapter_model.MemeCoinInfo) memeCoin {
	m.ID = e.ID
	m.Name = e.Name
	if e.Description != nil {
		m.Description.Valid = true
		m.Description.String = *e.Description
	}
	m.CreatedAt = e.CreatedAt
	m.PopularityScore = e.PopularityScore
	return m
}
