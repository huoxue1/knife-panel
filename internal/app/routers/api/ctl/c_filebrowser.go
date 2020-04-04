package ctl

import (
	"github.com/gin-gonic/gin"
	"knife-panel/internal/app/bll"
	"knife-panel/internal/app/ginplus"
	"knife-panel/internal/app/schema"
)

func NewFileBrowser(browser bll.IFileBrowser) *FileBrowser {
	return &FileBrowser{
		FileBrowserBll: browser,
	}
}

type FileBrowser struct {
	FileBrowserBll bll.IFileBrowser
}

func (a *FileBrowser) List(c *gin.Context) {
	basePath := c.Query("basePath")

	result, err := a.FileBrowserBll.List(ginplus.NewContext(c), basePath)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResSuccess(c, result.Data)
}

func (a *FileBrowser) Download(c *gin.Context) {
	err := a.FileBrowserBll.Download(c)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nil)
}

func (a *FileBrowser) Upload(c *gin.Context) {
	var item schema.FileItem
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.FileBrowserBll.Upload(c)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nil)
}

func (a *FileBrowser) Delete(c *gin.Context) {
	err := a.FileBrowserBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
