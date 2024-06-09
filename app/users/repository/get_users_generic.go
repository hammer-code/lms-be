package repository

import (
	"context"
	"encoding/json"

	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
)

// query untuk get user di database
func (repo *repository) GetUsersGenericConditions(ctx context.Context, filter domain.GetUserBy) (users []domain.User, err error) {
	db := repo.db.DB(ctx)

	// generic condition
	if filter.Role != "" {
		db = db.Where("role = ?", filter.Role)
	}

	if filter.ID != 0 {
		db = db.Where("id = ?", filter.ID)
	}

	if filter.Username != "" {
		db = db.Where("username = ?", filter.Username)
	}

	if filter.Email != "" {
		db = db.Where("email = ?", filter.Email)
	}

	err = db.Find(&users).Error
	if err != nil {
		c, _ := json.Marshal(filter)
		logrus.Error("repo.GetUsers: failed to get users use generic conditions", string(c))
		return
	}
	return
}
