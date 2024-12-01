package models

import (
	"time"

	"gorm.io/gorm"
)

type Emprunt struct {
	IDEmprunt          uint       `gorm:"primaryKey;column:id_emprunt"`
	UtilisateurID      uint       `gorm:"not null"`
	LivreID            uint       `gorm:"not null"`
	DateEmprunt        time.Time  `gorm:"not null"`
	DateRetourPrevu    time.Time  `gorm:"not null"`
	DateRetourEffectif *time.Time `gorm:"default:null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

type Penalite struct {
	IDPenalite    uint      `gorm:"primaryKey;column:id_penalite"`
	UtilisateurID uint      `gorm:"not null"`
	EmpruntID     uint      `gorm:"not null;column:emprunt_id;constraint:OnDelete:CASCADE"`
	Montant       float64   `gorm:"not null"`
	Paye          bool      `gorm:"default:false"`
	DateCalcul    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DatePaiement  *time.Time
	Emprunt       Emprunt `gorm:"foreignKey:EmpruntID;references:IDEmprunt"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
