package userLoader

import (
	"time"

	"github.com/bishal-dd/go-server/graph/model"
	"github.com/vikstrous/dataloadgen"
	"gorm.io/gorm"
)

func UserLoader(db *gorm.DB) *dataloadgen.Loader[string, *model.User] {
    // define the data loader
    ur := userReader{db: db}
    return dataloadgen.NewLoader(ur.getUsers, dataloadgen.WithWait(time.Millisecond))
}

