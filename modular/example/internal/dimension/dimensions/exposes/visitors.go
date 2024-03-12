package exposes

import (
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/dimensions/models"
)

var Visitors VisitorsExpose

type VisitorsExpose interface {
	// Count 访客数量
	Count() int

	// OpenIds 访客 OpenId 列表
	OpenIds() []string

	// Has 是否存在指定 OpenId 的访客
	Has(openId string) bool

	// Del 删除指定 OpenId 的访客
	Del(openId string)

	// Get 获取指定 OpenId 的访客
	Get(openId string) *models.VisitorsMember

	// Add 添加访客
	Add(member *models.VisitorsMember)
}
