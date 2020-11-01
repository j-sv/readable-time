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

var RFC3339 = time.RFC3339

type Time struct {
	time.Time
}

var Parse = time.Parse

// Clock returns the clock time of time day specified by `t` as a string.
// Even hours will be returned as "<h> o'clock", and noon and midnight will
// be returned as "noon" and "midnight" respectively.
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

// Hour returns the hour of the time `t` as a string. The formatting uses
// 12 hour time, meaning 13:00 would be returned as "one". 00:00 will be
// returned as "midnight". Formatting is tied to how the hour would be printed
// together with minutes. As 12:40 would be "twenty to one", the hour field alone
// would be returned as "one".
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

// Minute returns the minute of time `t` as a string rounded to the nearest
// 5 minute timestamp. 12:13 would be returned as "quarter past", while 12:12
// would be returned as "ten past".
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

// Day returns the ordinal day of month as a string. 2020-05-02 would e.g.
// return "second".
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

// Weekday returns the weekday of time `t` as a string. E.g. "wednesday".
func (t Time) Weekday() string {
	return strings.ToLower(t.Time.Weekday().String())
}

// Month returns the month of time `t` as a string. E.g. "may"
func (t Time) Month() string {
	return strings.ToLower(t.Time.Month().String())
}

// Now returns the current time
func Now() Time {
	return Time{time.Now()}
}

// FromTime takes a `Time` from the standard package and returns a `Time`
// from this package.
func FromTime(t time.Time) Time {
	return Time{t}
}
