package schema

import "time"

// FileItem demo对象
type FileItem struct {
	Id         string    `json:"id"`    // 记录ID
	Name       string    `json:"name" ` // 名称
	Dir        bool      `json:"dir"`
	Size       int64     `json:"size"`
	ModifyTime time.Time `json:"modifyTime"` // 创建时间
}
type FileItems []*FileItem

func (p FileItems) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p FileItems) Less(i, j int) bool {
	return p[i].Dir
}

// 交换数据
func (p FileItems) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// FileQueryResult demo对象查询结果
type FileQueryResult struct {
	Data []*FileItem
}
