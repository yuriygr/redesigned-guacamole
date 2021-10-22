package repository

import (
	"fmt"

	"github.com/yuriygr/go-posledstvie/models"
)

// GetSubsites - Получает подсайты
func (r *Repository) GetSubsites(request *models.SubsRequest) ([]*models.MiniSubsite, error) {
	subs := []*models.MiniSubsite{}

	var sql string

	switch request.Type {
	case "All":
		sql = selectSubs
	case "Recommendations":
		sql = selectSubsRecommendations
	default:
		sql = selectSubs
	}

	// Sort and limit
	//sql = fmt.Sprintf("%s order by c.state_id asc, %s %s", sql, request.Sort, request.Order)
	limit := request.Limit
	offset := request.Limit * (request.Page - 1)
	sql = fmt.Sprintf("%s limit %d offset %d", sql, limit, offset)

	nstmt, err := r.storage.PrepareNamed(sql)
	if err != nil {
		return nil, err
	}

	if err := nstmt.Select(&subs, request); err != nil {
		return nil, err
	}

	return subs, nil
}

// GetSubsite - Получает подсайт
func (r *Repository) GetSubsite(request *models.SubsiteRequest) (*models.Subsite, error) {
	subsite := models.Subsite{}

	nstmt, err := r.storage.PrepareNamed(selectSubsite)
	if err != nil {
		return &subsite, err
	}

	if err = nstmt.Get(&subsite, request); err != nil {
		return &subsite, err
	}

	return &subsite, nil
}

// ExistSubsite -
func (r *Repository) ExistSubsite(subsiteID uint32) (ok bool) {
	nstmt, err := r.storage.Preparex(existSubsite)
	if err != nil {
		return false
	}

	if err = nstmt.Get(&ok, subsiteID); err != nil {
		return false
	}

	return
}

// Subscribe -
func (r *Repository) Subscribe(request *models.SubsiteSubscribe) (err error) {
	_, err = r.storage.NamedExec(Subscribe, request)
	return
}

// Unubscribe -
func (r *Repository) Unubscribe(request *models.SubsiteUnsubscribe) (err error) {
	_, err = r.storage.NamedExec(Unsubscribe, request)
	return
}
