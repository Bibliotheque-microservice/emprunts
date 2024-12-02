//pour gérer la vérification de l'utilisateur 
package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Vérifie si l'utilisateur est actif et s'il n'a pas trop de pénalités
func CheckUserStatus(userID int) (bool, error) {
	url := fmt.Sprintf("http://service-user/userInfo/%d", userID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var userStatus struct {
		Active   bool    `json:"active"`
		Penalty float64 `json:"penalty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userStatus); err != nil {
		return false, err
	}

	// Vérifie si l'utilisateur est actif et que les pénalités sont sous le seuil
	if userStatus.Active && userStatus.Penalty <= 3.0 {
		return true, nil
	}
	return false, nil
}
