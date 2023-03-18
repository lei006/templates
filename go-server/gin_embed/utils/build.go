package utils

/*
const char* build_time(void)
{
static const char* psz_build_time = __TIME__;
    return psz_build_time;
}
const char* build_date(void)
{
static const char* psz_build_date = __DATE__ ;
    return psz_build_date;
}

*/
import "C"
import (
	"time"
)

var (
	buildTime = C.GoString(C.build_time())
	buildDate = C.GoString(C.build_date())
)

var dateLayout = "2006-01-02"
var datetimeLayout = "2006-01-02 15:04:05"

func _buildTime() string {
	return buildTime
}

func _buildDate() string {
	return buildDate
}

func BuildDateTime() time.Time {
	// 日期转时间
	const shortForm = "Jan 02 2006"
	t, _ := time.Parse(shortForm, _buildDate())

	// 析出：日期时间
	build_time := t.Format(dateLayout) + " " + _buildTime()
	duetimecst, _ := time.ParseInLocation(datetimeLayout, build_time, time.Local)
	return duetimecst
}

func BuildTime() string {
	t := BuildDateTime()
	return t.Format(datetimeLayout)
}

func BuildDate() string {
	t := BuildDateTime()
	return t.Format(dateLayout)
}
