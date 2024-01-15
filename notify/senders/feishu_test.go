package senders

import (
	"github.com/kercylan98/minotaur/notify/notifies"
	"testing"
)

func TestFeiShu_Push(t *testing.T) {
	fs := NewFeiShu("https://open.feishu.cn/open-apis/bot/v2/hook/bid")

	rt := notifies.NewFeiShu(notifies.FeiShuMessageWithRichText(notifies.NewFeiShuRichText().Create("zh_cn", "标题咯").AddText("哈哈哈").Ok()))
	if err := fs.Push(rt); err != nil {
		panic(err)
	}
}
