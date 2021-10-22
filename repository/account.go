package repository

import (
	"fmt"

	"github.com/yuriygr/go-posledstvie/models"
	"github.com/yuriygr/go-posledstvie/utils"
)

// GetAccountByEmail - Получает пользователя по электронной почте
func (r *Repository) GetAccountByEmail(email string) (*models.Account, error) {
	account := models.Account{}
	sql := selectAccount

	sql = fmt.Sprintf("%s and a.email = '%s'", sql, email)

	if err := r.storage.Get(&account, sql); err != nil {
		return &account, err
	}

	return &account, nil
}

// CreateAccount - Создание пользователя
func (r *Repository) CreateAccount(request *models.Account) (*models.Account, error) {
	result, err := r.storage.NamedExec(inserAccount, request)
	if err != nil {
		return nil, err
	}

	AccountID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// And we must create profile and stats to user...
	// TODO: Transfer to user controller?

	request.ID = utils.Uint32(AccountID)

	if _, err := r.storage.NamedExec(inserAccountSubsite, request); err != nil {
		return nil, err
	}

	account, err := r.GetAccountByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateAccountLogin - Создаем запись об авторизации пользователя
func (r *Repository) CreateAccountLogin(request *models.AccountLogin) error {
	if _, err := r.storage.NamedExec(inserAccountLogin, request); err != nil {
		return err
	}
	return nil
}

// GetLogin - Получает список авторизаций
func (r *Repository) GetLogin(request *models.LoginRequest) ([]*models.AccountLogin, error) {
	login := []*models.AccountLogin{}
	sql := selectLogin

	// Account
	sql = fmt.Sprintf("%s where al.account_id = '%d'", sql, request.AccountID)

	// Order
	sql = fmt.Sprintf("%s order by al.date desc", sql)

	// Sort and limit
	limit := request.Limit
	offset := request.Limit * (request.Page - 1)
	sql = fmt.Sprintf("%s limit %d offset %d", sql, limit, offset)

	if err := r.storage.Select(&login, sql); err != nil {
		return nil, err
	}

	return login, nil
}
