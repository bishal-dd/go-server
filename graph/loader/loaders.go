package loader

import (
	"github.com/bishal-dd/go-server/graph/loader/userLoader"
	"github.com/bishal-dd/go-server/graph/model"
	"github.com/vikstrous/dataloadgen"
	"gorm.io/gorm"
)


type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(db *gorm.DB) *Loaders {
    // define the data loader
    return &Loaders{
        UserLoader: userLoader.UserLoader(db),
    }
}