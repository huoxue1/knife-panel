package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"knife-panel/internal/app/schema"
	"os"
	"path"
)

// NewFileBrowser 创建demo
func NewFileBrowser() *FileBrowser {
	return &FileBrowser{}
}

// FileBrowser 示例程序
type FileBrowser struct {
}

// List 查询数据
func (a *FileBrowser) List(ctx context.Context, basePath string) (*schema.FileQueryResult, error) {
	fileList, e := ioutil.ReadDir(basePath)
	if e != nil {
		return nil, e
	}
	fileItems := make([]*schema.FileItem, 0)
	for _, v := range fileList {
		if len(v.Name()) == 0 {
			continue
		}
		fileItem := schema.FileItem{
			Id:         path.Join(basePath, v.Name()),
			Name:       v.Name(),
			Dir:        v.IsDir(),
			ModifyTime: v.ModTime(),
		}
		fileItems = append(fileItems, &fileItem)
	}
	return &schema.FileQueryResult{Data: fileItems}, nil
}

// Download 查询指定数据
func (a *FileBrowser) Download(c *gin.Context) error {
	filename := c.Param("id")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filename)
	return nil
}

// Upload 创建数据
func (a *FileBrowser) Upload(c *gin.Context) error {
	file, _ := c.FormFile("file")
	dst := c.PostForm("dst")
	return c.SaveUploadedFile(file, dst)
}

// Delete 删除数据
func (a *FileBrowser) Delete(ctx context.Context, id string) error {
	if info, err := os.Stat(id); err == nil {
		if info.IsDir() {
			return os.RemoveAll(id)
		} else {
			return os.Remove(id)
		}
	} else {
		return err
	}
}
