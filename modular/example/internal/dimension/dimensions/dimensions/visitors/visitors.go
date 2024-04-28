package visitors

import (
	"fmt"
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/core"
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/dimensions/exposes"
	"github.com/kercylan98/minotaur/modular/example/internal/dimension/dimensions/models"
"github.com/kercylan98/minotaur/toolkit/collection"
)
type Dimension struct {
	*core.Room                                   // 房间 Id
	visitors   map[string]*models.VisitorsMember // 所有访客
	visitorIds []string                          // 所有访客 OpenId
}

func (d *Dimension) OnInit(owner *core.Room) error {
	exposes.Visitors = d
	d.Room = owner
	d.visitors = make(map[string]*models.VisitorsMember)
	fmt.Println("visitors dimension initialized")
	return nil
}

func (d *Dimension) OnPreload() error {
	fmt.Println("visitors dimension preloaded")
	return nil
}

func (d *Dimension) OnMount() error {
	fmt.Println("visitors dimension mounted")
	return nil
}

func (d *Dimension) Count() int {
	return len(d.visitors)
}

func (d *Dimension) OpenIds() []string {
	return d.visitorIds
}

func (d *Dimension) Has(openId string) bool {
	return collection.KeyInMap(d.visitors, openId)
}

func (d *Dimension) Del(openId string) {
	member := d.Get(openId)
	if member == nil {
		return
	}
	delete(d.visitors, openId)
	collection.DropSliceByIndices(&d.visitorIds, member.OpenIdIdx)
}

func (d *Dimension) Get(openId string) *models.VisitorsMember {
	return d.visitors[openId]
}

func (d *Dimension) Add(member *models.VisitorsMember) {
	member.OpenIdIdx = len(d.visitorIds)
	d.visitorIds = append(d.visitorIds, member.OpenId)
	d.visitors[member.OpenId] = member
}
