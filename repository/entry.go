package repository

import (
	"fmt"

	"github.com/yuriygr/go-posledstvie/models"
)

// GetTimelineEntries - Получает записи ленты
func (r *Repository) GetTimelineEntries(request *models.EntriesRequest) ([]*models.Entry, error) {
	entries := []*models.Entry{}
	sql := selectTimelineEntries

	if request.PinnedOnTop {
		sql = fmt.Sprintf("%s order by e.is_pinned desc, %s %s", sql, request.Sort, request.Order)
	} else {
		sql = fmt.Sprintf("%s order by %s %s", sql, request.Sort, request.Order)
	}

	nstmt, err := r.storage.PrepareNamed(sql)
	if err != nil {
		return nil, err
	}

	if err := nstmt.Select(&entries, request); err != nil {
		return nil, err
	}

	return entries, nil
}

// GetSubsiteEntries - Получает записи подсайта
func (r *Repository) GetSubsiteEntries(request *models.EntriesRequest) ([]*models.Entry, error) {
	entries := []*models.Entry{}
	sql := selectSubsiteEntries

	if request.PinnedOnTop {
		sql = fmt.Sprintf("%s order by e.is_pinned desc, %s %s", sql, request.Sort, request.Order)
	} else {
		sql = fmt.Sprintf("%s order by %s %s", sql, request.Sort, request.Order)
	}

	nstmt, err := r.storage.PrepareNamed(sql)
	if err != nil {
		return nil, err
	}

	if err := nstmt.Select(&entries, request); err != nil {
		return nil, err
	}

	return entries, nil
}

// GetEntry -
func (r *Repository) GetEntry(request *models.EntryRequest) (*models.Entry, error) {
	entry := models.Entry{}

	nstmt, err := r.storage.PrepareNamed(selectEntry)
	if err != nil {
		return &entry, err
	}

	if err = nstmt.Get(&entry, request); err != nil {
		return &entry, err
	}

	return &entry, nil
}

// EntryLike -
func (r *Repository) EntryLike(request *models.EntryLikeRequest) (int16, error) {
	//_, err = r.storage.NamedExec(likeEntry, request)
	return 9999, nil
}

// EntryDislike -
func (r *Repository) EntryDislike(request *models.EntryLikeRequest) (int16, error) {
	//_, err = r.storage.NamedExec(dislikeEntry, request)
	return 9999, nil
}
