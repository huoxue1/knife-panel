package bll

import (
	"context"
	"github.com/gin-gonic/gin"

	"knife-panel/internal/app/schema"
)

// IFileBrowser demo业务逻辑接口
type IFileBrowser interface {
	// 查询数据
	List(ctx context.Context, basePath string) (*schema.FileQueryResult, error)
	// 查询指定数据
	Download(c *gin.Context) error
	// 创建数据
	Upload(c *gin.Context) error
	// 删除数据
	Delete(ctx context.Context, id string) error
}
