// vérification la disponibilité du livre et le statut de l'utilisateur
package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Vérifie la disponibilité du livre via le service externe
func CheckBookAvailability(bookID int) (bool, error) {
	url := fmt.Sprintf("http://service-livre/%d/disponibilité", bookID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var availability struct {
		Available bool `json:"available"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&availability); err != nil {
		return false, err
	}
	return availability.Available, nil
}
