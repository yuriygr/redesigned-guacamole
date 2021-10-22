package repository

import "github.com/yuriygr/go-posledstvie/models"

// GetUser - Получает подсайт
func (r *Repository) GetUser(request *models.SubsiteRequest) (*models.User, error) {
	subsite := models.User{}

	nstmt, err := r.storage.PrepareNamed(selectUser)
	if err != nil {
		return &subsite, err
	}

	if err = nstmt.Get(&subsite, request); err != nil {
		return &subsite, err
	}

	return &subsite, nil
}
