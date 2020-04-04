package schema

import "time"

// FileItem demo对象
type FileItem struct {
	Id         string    `json:"id"`    // 记录ID
	Name       string    `json:"name" ` // 名称
	Dir        bool      `json:"dir"`
	ModifyTime time.Time `json:"modifyTime"` // 创建时间
}

// FileQueryResult demo对象查询结果
type FileQueryResult struct {
	Data []*FileItem
}
