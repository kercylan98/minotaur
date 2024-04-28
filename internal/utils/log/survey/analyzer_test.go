package survey_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log/survey"
	"testing"
)

func TestIncrementAnalyze(t *testing.T) {
	path := `./test/day.2023-09-06.log`

	reader := survey.IncrementAnalyze(path, func(analyzer *survey.Analyzer, record survey.R) {
		switch record.GetString("type") {
		case "open_conn":
			analyzer.IncreaseValueNonRepeat("开播人数", record, 1, "live_id")
		case "report_rank":
			analyzer.IncreaseValue("开始游戏次数", 1)
			analyzer.Increase("开播时长", record, "game_time")
			analyzer.Sub(record.GetString("live_id")).IncreaseValue("开始游戏次数", 1)
			analyzer.Sub(record.GetString("live_id")).Increase("开播时长", record, "game_time")
		case "statistics":
			analyzer.IncreaseValueNonRepeat("活跃人数", record, 1, "active_player")
			analyzer.IncreaseValueNonRepeat("评论人数", record, 1, "comment_player")
			analyzer.IncreaseValueNonRepeat("点赞人数", record, 1, "like_player")
			analyzer.Sub(record.GetString("live_id")).IncreaseValueNonRepeat("活跃人数", record, 1, "active_player")
			analyzer.Sub(record.GetString("live_id")).IncreaseValueNonRepeat("评论人数", record, 1, "comment_player")
			analyzer.Sub(record.GetString("live_id")).IncreaseValueNonRepeat("点赞人数", record, 1, "like_player")

			giftId := record.GetString("gift.gift_id")
			if len(giftId) > 0 {
				giftPrice := record.GetFloat64("gift.price")
				giftCount := record.GetFloat64("gift.count")
				giftSender := record.GetString("gift.gift_senders")

				analyzer.IncreaseValue("礼物总价值", giftPrice*giftCount)
				analyzer.IncreaseValueNonRepeat(fmt.Sprintf("送礼人数_%s", giftId), record, 1, giftSender)
				analyzer.IncreaseValue(fmt.Sprintf("礼物总数_%s", giftId), giftCount)

				analyzer.Sub(record.GetString("live_id")).IncreaseValue("礼物总价值", giftPrice*giftCount)
				analyzer.Sub(record.GetString("live_id")).IncreaseValueNonRepeat(fmt.Sprintf("送礼人数_%s", giftId), record, 1, giftSender)
				analyzer.Sub(record.GetString("live_id")).IncreaseValue(fmt.Sprintf("礼物总数_%s", giftId), giftCount)
			}

		}
	})

	for i := 0; i < 10; i++ {
		report, err := reader()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(report.FilterSub("warzone0009"))
	}
}
