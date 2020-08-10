package time

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/divan/num2words"
)

const (
	fiveMinutes = 5.0
	pastHalf    = 32
)

type Time struct {
	time.Time
}

func (t Time) Clock() string {
	if t.minuteRounded() == 0 {
		switch t.hourRounded() {
		case 0:
			return "midnight"
		case 12:
			return "noon"
		default:
			return fmt.Sprintf("%s o'clock", t.Hour())
		}
	}

	return fmt.Sprintf("%s %s", t.Minute(), t.Hour())
}

func (t Time) minuteRounded() int {
	fivers := math.Round(float64(t.Time.Minute()) / fiveMinutes)
	return int(fivers*fiveMinutes) % 60
}

func (t Time) hourRounded() int {
	hour := t.Time.Hour()
	if t.Time.Minute() > pastHalf {
		hour++
	}

	return hour
}

func (t Time) Hour() string {
	hour := t.hourRounded()

	if hour == 0 {
		return "midnight"
	}

	if hour < 12 {
		return num2words.Convert(hour)
	}

	hour %= 12

	if hour == 0 {
		return "twelve" // num2words.Convert(12)
	}

	return num2words.Convert(hour)
}

func (t Time) Minute() string {
	asInt := t.minuteRounded()
	switch asInt {
	case 5, 10, 20, 25:
		return fmt.Sprintf("%s past", num2words.Convert(asInt))
	case 15:
		return "quarter past"
	}

	asInt = int(math.Abs(float64(asInt) - 60.0))
	switch asInt {
	case 30:
		return "half past"
	case 5, 10, 20, 25:
		return fmt.Sprintf("%s to", num2words.Convert(asInt))
	case 15:
		return "quarter to"
	}

	return ""
}

func (t Time) Day() string {
	var toStr func(int) string
	toStr = func(day int) string {
		tens := day % 10

		switch {
		case day == 1:
			return "first"
		case day == 2:
			return "second"
		case day == 3:
			return "third"
		case day == 4:
			return "fourth"
		case day == 5:
			return "fifth"
		case day == 6:
			return "sixth"
		case day == 7:
			return "seventh"
		case day == 8:
			return "eighth"
		case day == 9:
			return "ninth"
		case day == 12:
			return "twelfth"
		case day < 20:
			return fmt.Sprintf("%sth", num2words.Convert(day))
		case tens == 0:
			switch {
			case day == 20:
				return "twentieth"
			case day == 30:
				return "thirtieth"
			default:
				panic("NO!")
			}
		default:
			return fmt.Sprintf("%s-%s", num2words.Convert(day-tens), toStr(tens))
		}
	}

	return toStr(t.Time.Day())
}

func (t Time) Weekday() string {
	return strings.ToLower(t.Time.Weekday().String())
}

func (t Time) Month() string {
	return strings.ToLower(t.Time.Month().String())
}

func Now() Time {
	return Time{time.Now()}
}

func FromTime(t time.Time) Time {
	return Time{t}
}
