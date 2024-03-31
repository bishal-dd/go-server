package userLoader

import (
	"context"

	"github.com/bishal-dd/go-server/graph/model"
	"gorm.io/gorm"
)

type userReader struct {
    db *gorm.DB
}

// getUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (u *userReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
    var users []*model.User
    if err := u.db.Where("id IN (?)", userIDs).Find(&users).Error; err != nil {
        return nil, []error{err}
    }
    return users, nil
}