package repository

import (
	"fmt"

	"github.com/yuriygr/go-posledstvie/models"
)

// GetFastSearch -
func (r *Repository) GetFastSearch(request *models.SearchRequest) ([]*models.FastSearch, error) {
	subsities := []*models.MiniSubsite{}
	sql := selectFastSearch

	if request.Query != "" {
		sql = fmt.Sprintf("%s and (s.name like '%%%s%%')", sql, request.Query)
	}

	// Sort and limit
	//sql = fmt.Sprintf("%s order by c.state_id asc, %s %s", sql, request.Sort, request.Order)
	limit := request.Limit
	offset := request.Limit * (request.Page - 1)
	sql = fmt.Sprintf("%s limit %d offset %d", sql, limit, offset)

	if err := r.storage.Select(&subsities, sql); err != nil {
		return nil, err
	}

	result := []*models.FastSearch{}

	for _, subsite := range subsities {
		result = append(result, &models.FastSearch{
			Value: subsite.Name,
			Type:  "subsite",
			UUID:  subsite.Avatar,
			URL:   "",
		})
	}

	return result, nil
}
