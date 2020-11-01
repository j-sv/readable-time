package when

import (
	"fmt"

	"github.com/j-sv/readable-time/time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func Parse(timestamp string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, timestamp); err == nil {
		return time.FromTime(t), nil
	}

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(timestamp, time.Now().Time)
	if err != nil {
		return time.Time{}, err
	}

	if r == nil {
		return time.Time{}, fmt.Errorf("wasn't parsed into a time: %s", timestamp)
	}

	return time.FromTime(r.Time), nil
}
