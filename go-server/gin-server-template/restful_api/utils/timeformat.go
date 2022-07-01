package utils

import (
	"time"
)

func DateFormat(fmt, date time.Time) string {

	//fmt.Println(time.Now())
	/*
		opt := make(map[string]string)
		opt["Y+"] = fmt.Sprintf("%d", date.Year())
		opt["m+"] = fmt.Sprintf("%d", date.Month())
		opt["d+"] = fmt.Sprintf("%d", date.Day())
		opt["H+"] = fmt.Sprintf("%d", date.Hour())
		opt["M+"] = fmt.Sprintf("%d", date.Minute())
		opt["S+"] = fmt.Sprintf("%d", date.Minute())

		for k := range opt {
			//match, _ := regexp.MatchString("("+k+")", fmt)
			//ret = new RegExp("(" + k + ")").exec(fmt);

			fmt.Println(k)
		}
	*/

	return ""
}

/*
function dateFormat(fmt, date) {
    let ret;
    const opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}
*/
