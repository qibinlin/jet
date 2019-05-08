// Query building functions for expression components
package sqlbuilder

import (
	"strconv"
	"strings"
	"time"
)

type intervalExpression struct {
	expressionInterfaceImpl
	duration time.Duration
}

const intervalSep = ":"

func (c *intervalExpression) serialize(statement statementType, out *queryData) error {
	out.writeString("INTERVAL '")

	duration := c.duration

	if duration < 0 {
		duration = -duration
		out.writeString("-")
	}

	hours := duration / time.Hour
	minutes := (duration % time.Hour) / time.Minute
	sec := (duration % time.Minute) / time.Second
	msec := (duration % time.Second) / time.Microsecond

	out.writeString(strconv.FormatInt(int64(hours), 10))
	out.writeString(intervalSep)
	out.writeString(strconv.FormatInt(int64(minutes), 10))
	out.writeString(intervalSep)
	out.writeString(strconv.FormatInt(int64(sec), 10))
	out.writeString(intervalSep)
	out.writeString(strconv.FormatInt(int64(msec), 10))
	out.writeString("' HOUR_MICROSECOND")

	return nil
}

//// Interval returns a representation of duration
//func Interval(duration time.Duration) expression {
//	intervalExp := &intervalExpression{
//		duration: duration,
//	}
//
//	intervalExp.expressionInterfaceImpl.parent = intervalExp
//
//	return intervalExp
//}

var likeEscaper = strings.NewReplacer("_", "\\_", "%", "\\%")

func EscapeForLike(s string) string {
	return likeEscaper.Replace(s)
}
