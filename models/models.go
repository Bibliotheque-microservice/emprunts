package models

import (
	"time"

	"gorm.io/gorm"
)

type Emprunt struct {
	gorm.Model
	UtilisateurID      uint       `gorm:"not null"`
	LivreID            uint       `gorm:"not null"`
	DateEmprunt        time.Time  `gorm:"not null"`
	DateRetourPrevu    time.Time  `gorm:"not null"`
	DateRetourEffectif *time.Time `gorm:"default:null"`
	DateCreation       time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
}

type Penalite struct {
	gorm.Model
	UtilisateurID uint      `gorm:"not null"`
	EmpruntID     uint      `gorm:"not null"`
	Montant       float64   `gorm:"not null"`
	Paye          bool      `gorm:"default:false"`
	DateCalcul    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DatePaiement  *time.Time
	Emprunt       Emprunt `gorm:"foreignKey:EmpruntID"`
}
