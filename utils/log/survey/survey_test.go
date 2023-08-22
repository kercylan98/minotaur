package survey_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log/survey"
	"os"
	"testing"
	"time"
)

func TestRecord(t *testing.T) {
	_ = os.MkdirAll("./test", os.ModePerm)
	survey.RegSurvey("GLOBAL_DATA", "./test/global_data.log")
	survey.Record("GLOBAL_DATA", map[string]any{
		"joinTime": time.Now().Unix(),
		"action":   1,
	})
	survey.Flush()

	fmt.Println(survey.Sum("GLOBAL_DATA", time.Now(), "action"))
}
