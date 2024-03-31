package loader

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Middleware(db *gorm.DB) gin.HandlerFunc {
    // return a middleware that injects the loader to the request context
    return func(c *gin.Context) {
        loader := NewLoaders(db)
        c.Set("loaders", loader)
        c.Next()
    }
}