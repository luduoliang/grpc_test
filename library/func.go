package library

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//md5密码
func Md5(password string) string {
	w := md5.New()
	io.WriteString(w, password)                  //将str写入到w中
	passwordMd5 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return passwordMd5
}

//从字符串转换成int64方法
func Str2Int64(str string) int64 {
	strInt64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return strInt64
}

//从int64转换成字符串方法
func Int642Str(data int64) string {
	return strconv.FormatInt(data, 10)
}

//从float64转换成字符串方法
func Float642Str(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

//从字符串转换成float64方法
func Str2Float64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

//获取为int64的参数
func ParamsInt64(data interface{}) (result int64) {
	result = 0
	switch data.(type) { //多选语句switch
	case string:
		result = Str2Int64(data.(string))
	case float64:
		result = int64(data.(float64))
	}
	return result
}

//获取当前时间戳
func CurrentTime() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location).UTC().Unix()
}

//截取字符串，含中文
func Substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

//判断是否在数组中
func InArrayString(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//判断是否在数组中
func InArrayInt8(need int8, needArr []int8) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

//随机一段整数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

//字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

//生成干扰字符串
func RandString(len int) string {
	str := "1123456789QWERTYUIOPASDFGHJKLZXVBNMqwertyuioplkjhgfdsamnbvcxzz"
	var returnStr string = ""
	for i := 0; i < len; i++ {
		index := GenerateRangeNum(i+1, 57)
		returnStr = returnStr + str[index:index+1]
	}
	return returnStr
}

//反转数组
func ReverseArray(arr []interface{}) []interface{} {
	len := len(arr)
	tempArr := make([]interface{}, 0)
	for i := len - 1; i >= 0; i-- {
		tempArr = append(tempArr, arr[i])
	}
	return tempArr
}

//请求接口
func RequestUrl(queryUrl string, method string, data interface{}) (content string, err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp := &http.Response{}
	if method == "GET" {
		//如果是GET请求，把参数组装到url中去
		dataTemp := data.(map[string]string)
		if len(dataTemp) != 0 {
			index := 0
			for k, v := range dataTemp {
				if index == 0 {
					queryUrl = queryUrl + "?" + k + "=" + url.PathEscape(v)
				} else {
					queryUrl = queryUrl + "&" + k + "=" + url.PathEscape(v)
				}
				index = index + 1
			}
		}
		resp, err = client.Get(queryUrl)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
	} else {
		dataValues := url.Values{}
		for key, val := range data.(map[string]string) {
			dataValues.Add(key, val)
		}
		resp, err = client.Post(queryUrl, "application/x-www-form-urlencoded", strings.NewReader(dataValues.Encode()))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, location)
}

//时间戳转换时间
func GetTimeFromUnix(t int64) *time.Time {
	timeData := time.Unix(t, 0)
	return &timeData
}

//获取传入时间周一的日期
func GetFirstDateOfWeek(d time.Time) *time.Time {
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}
	t := d.AddDate(0, 0, offset)
	return &t
}

//获取当前时间向前推时间节点
func GetDateFormat() map[string]string {
	nowTime := time.Now()
	//当天日期
	todayTime := nowTime.Format("2006-01-02")

	//昨天日期
	yesterdayTime := nowTime.AddDate(0, 0, -1).Format("2006-01-02")

	//本周开始日期，结束日期
	thisWeekTime := GetFirstDateOfWeek(nowTime)
	thisWeekStart := thisWeekTime.Format("2006-01-02")
	thisWeekEnd := thisWeekTime.AddDate(0, 0, 6).Format("2006-01-02")

	//上周开始日期，结束日期
	lastWeekTime := GetFirstDateOfWeek(GetFirstDateOfWeek(nowTime).AddDate(0, 0, -2))
	lastWeekStart := lastWeekTime.Format("2006-01-02")
	lastWeekEnd := lastWeekTime.AddDate(0, 0, 6).Format("2006-01-02")

	//一周开始日期，结束日期
	oneWeekStart := nowTime.AddDate(0, 0, -6).Format("2006-01-02")
	oneWeekEnd := todayTime

	//二周开始日期，结束日期
	twoWeekStart := nowTime.AddDate(0, 0, -13).Format("2006-01-02")
	twoWeekEnd := todayTime

	//本月开始日期，结束日期
	thisMonthStart := GetFirstDateOfMonth(nowTime).Format("2006-01-02")
	thisMonthEnd := GetLastDateOfMonth(nowTime).Format("2006-01-02")

	//上月开始日期，结束日期
	lastMonthTime := GetFirstDateOfMonth(nowTime).AddDate(0, 0, -2)
	lastMonthStart := GetFirstDateOfMonth(lastMonthTime).Format("2006-01-02")
	lastMonthEnd := GetLastDateOfMonth(lastMonthTime).Format("2006-01-02")

	//上上月开始日期，结束日期
	lastLastMonthTime := GetFirstDateOfMonth(lastMonthTime).AddDate(0, 0, -2)
	lastLastMonthStart := GetFirstDateOfMonth(lastLastMonthTime).Format("2006-01-02")
	lastLastMonthEnd := GetLastDateOfMonth(lastLastMonthTime).Format("2006-01-02")

	timeFormat := map[string]string{
		"todayTime":          todayTime,          //今天日期
		"yesterdayTime":      yesterdayTime,      //昨天日期
		"thisWeekStart":      thisWeekStart,      //本周开始日期
		"thisWeekEnd":        thisWeekEnd,        //本周结束日期
		"lastWeekStart":      lastWeekStart,      //上周开始日期
		"lastWeekEnd":        lastWeekEnd,        //上周结束日期
		"oneWeekStart":       oneWeekStart,       //一周开始日期
		"oneWeekEnd":         oneWeekEnd,         //一周结束日期
		"twoWeekStart":       twoWeekStart,       //二周开始日期
		"twoWeekEnd":         twoWeekEnd,         //二周结束日期
		"thisMonthStart":     thisMonthStart,     //本月开始日期
		"thisMonthEnd":       thisMonthEnd,       //本月结束日期
		"lastMonthStart":     lastMonthStart,     //上月开始日期
		"lastMonthEnd":       lastMonthEnd,       //上月结束日期
		"lastLastMonthStart": lastLastMonthStart, //上上月开始日期
		"lastLastMonthEnd":   lastLastMonthEnd,   //上上月结束日期
	}

	return timeFormat
}

//生成订单
func CreateOrderNo(user_id uint) string {
	user_id += 10000
	user_id_str := strconv.Itoa(int(user_id))
	count := len(user_id_str)
	if count < 8 {
		for i := 0; i < (8 - count); i++ {
			user_id_str = "0" + user_id_str
		}
	}
	rand_str := strconv.Itoa(GenerateRangeNum(10000, 99999))
	order_no := time.Now().Format("20060102150405") + user_id_str + rand_str
	return order_no
}

//随机一个表情
func RandEmoji() string {
	emojiList := []string{
		"😎", "😍", "😱", "👉", "👏", "👊", "👍", "✌", "✨", "📢", "🔔", "💡", "⚡", "🎁", "🎉", "☁", "🌛", "💰", "⚠", "🍀", "🔥", "🎈", "🌴", "🍃", "🙆", "🏃", "🌈", "🌸", "😸", "😺", "😑", "😭", "😝", "😜", "😂", "😡", "😛", "🌛", "🌱", "🌾", "🚌", "🚎", "🚊", "🌄", "🗻", "💥", "📣", "👑", "❕", "❓", "🔫", "😩", "🍰", "👯", "🍣", "😁", "🙀", "😻", "🍤", "💫", "🌟", "😘", "🍓", "➰", "🌺", "💕", "🌜", "🙋", "💣", "👌", "😄", "😃", "😀", "😊", "☺", "😉", "😍", "😘", "😚", "😗", "😙", "😜", "😝", "😛", "😳", "😁", "😔", "😌", "😒", "😞", "😣", "😢", "😂", "😭", "😪", "😥", "😰", "😅", "😓", "😩", "😫", "😨", "😱", "😠", "😡", "😤", "😖", "😆", "😋", "😷", "😎", "😴", "😵", "😲", "😟", "😦", "😧", "😈", "👿", "😮", "😬", "😐", "😕", "😯", "😶", "😇", "😏", "😑", "😺", "😸", "😻", "😽", "😼", "🙀", "😿", "😹", "😾", "👹", "👺", "👂", "👀", "👃", "👅", "👄", "👍", "👎", "👌", "👊", "✊", "✌", "👋", "✋", "👐", "👆", "👇", "👉", "👈", "🙌", "🙏", "☝", "👏", "💪", "🚶", "🏃", "💃", "👫", "👪", "👬", "👭", "💏", "💑", "👯", "🙆", "🙅", "💁", "🙋", "💇", "💅", "👰", "🙎", "🙍", "🙇", "🎩", "👑", "👒", "👟", "👞", "👡", "👠", "👢", "👕", "👔", "👚", "👗", "🎽", "👖", "👘", "👙", "💼", "👜", "👝", "👛", "👓", "🎀", "🌂", "💄", "💋", "👣", "💎", "💍", "👑", "🔥", "✨", "🌟", "💫", "💥", "🎀", "🌂", "💄", "💛", "💙", "💜", "💚", "❤", "💔", "💗", "💓", "💕", "💖", "💞", "💘", "💌", "💋", "🎍", "💝", "🎎", "🎒", "🎓", "🎏", "🎆", "🎇", "🎐", "🎑", "🎃", "👻", "🎅", "🎄", "🎁", "🎋", "🎉", "🎊", "🎈", "🎌", "💐", "🌸", "🌷", "🍀", "🌹", "🌻", "🌺", "🍁", "🍃", "🍂", "🌿", "🌾", "🍄", "🌵", "🌴", "🌲", "🌳", "🌰", "🌱", "🌼", "🌐", "🌞", "🌝", "🌚", "🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘", "🌜", "🌛", "🌙", "🌍", "🌎", "🌏", "🌋", "🌌", "🌠", "⭐", "☀", "⛅", "☁", "⚡", "☔", "❄", "⛄", "🌀", "🌁", "🌈", "🌊", "🔥", "✨", "🌟", "💫", "💥", "💢", "💦", "💧", "💤", "💨", "☕", "🍵", "🍶", "🍼", "🍺", "🍻", "🍸", "🍹", "🍷", "🍴", "🍕", "🍔", "🍟", "🍗", "🍖", "🍝", "🍛", "🍤", "🍱", "🍣", "🍥", "🍙", "🍘", "🍚", "🍜", "🍲", "🍢", "🍡", "🍳", "🍞", "🍩", "🍮", "🍦", "🍨", "🍧", "🎂", "🍰", "🍪", "🍫", "🍬", "🍭", "🍯", "🍎", "🍏", "🍊", "🍋", "🍒", "🍇", "🍉", "🍓", "🍑", "🍈", "🍌", "🍐", "🍍", "🍠", "🍆", "🍅", "🌽", "🐶", "🐺", "🐱", "🐭", "🐹", "🐰", "🐸", "🐯", "🐨", "🐻", "🐷", "🐽", "🐮", "🐗", "🐵", "🐒", "🐴", "🐑", "🐘", "🐼", "🐧", "🐦", "🐤", "🐥", "🐣", "🐔", "🐍", "🐢", "🐛", "🐝", "🐜", "🐞", "🐌", "🐙", "🐚", "🐠", "🐟", "🐬", "🐳", "🐋", "🐄", "🐏", "🐀", "🐃", "🐅", "🐇", "🐉", "🐎", "🐐", "🐓", "🐕", "🐖", "🐁", "🐂", "🐲", "🐡", "🐊", "🐫", "🐪", "🐆", "🐈", "🐩", "🐾", "🙈", "🙉", "🙊", "💀", "👽", "😺", "😸", "😻", "😻", "😽", "😼", "🙀", "😿", "😹", "😾", "📰", "🎨", "🎬", "🎤", "🎧", "🎼", "🎵", "🎶", "🎹", "🎻", "🎷", "🎸", "👾", "🎮", "🃏", "🎴", "🀄", "🎲", "🎯", "🏈", "🏀", "⚽", "⚾", "🎾", "🎱", "🏉", "🎳", "⛳", "🚵", "🚴", "🏁", "🏇", "🏆", "🎿", "🏂", "🏊", "🏄", "🎣", "🔮", "🎥", "📷", "📹", "📼", "💿", "📀", "💽", "💾", "💻", "📱", "☎", "📞", "📟", "📠", "📡", "📺", "📻", "🔊", "🔉", "🔈", "🔇", "🔔", "🔕", "📢", "📣", "⏳", "⌛", "⏰", "⌚", "🔓", "🔒", "🔐", "🔑", "🔎", "💡", "🔦", "🔆", "🔅", "🔌", "🔋", "🔍", "🛁", "🛀", "🚿", "🚽", "🔧", "🔩", "🔨", "🚪", "🚬", "💣", "🔫", "🔪", "💊", "💰", "💴", "💵", "💷", "💶", "💳", "💸", "📲", "📧", "📥", "📤", "✉", "📩", "📨", "📯", "📫", "📪", "📬", "📭", "📮", "📦", "📝", "📄", "📃", "📑", "📊", "📈", "📉", "📜", "📋", "📅", "📆", "📇", "📁", "📂", "✂", "📌", "📎", "✒", "✏", "📏", "📐", "📕", "📗", "📘", "📙", "📓", "📔", "📒", "📚", "📖", "🔖", "📛", "🔬", "🏠", "🏡", "🏫", "🏢", "🏣", "🏥", "🏦", "🏪", "🏩", "🏨", "💒", "⛪", "🏬", "🏤", "🌇", "🌆", "🏯", "🏰", "⛺", "🏭", "🗼", "🗾", "🗻", "🌄", "🚢", "⛵", "🚤", "🚣", "⚓", "🚀", "✈", "💺", "🚁", "🚂", "🚊", "🚉", "🚞", "🚆", "🚄", "🚅", "🚈", "🚇", "🚝", "🚋", "🚃", "🚎", "🚌", "🚍", "🚙", "🚘", "🚗", "🚕", "🚖", "🚛", "🚚", "🚨", "🚓", "🚔", "🚒", "🚑", "🚐", "🚲", "🚡", "🚟", "🚠", "🚜", "💈", "🚏", "🎫", "🚦", "🚥", "⚠", "🚧", "🔰", "⛽", "🏮", "🎰", "♨", "🗿", "🎪", "🎭", "📍", "🚩", "🐁", "🐂", "🐅", "🐇", "🐉", "🐍", "🐎", "🐐", "🐒", "🐓", "🐕", "🐖", "♈", "♉", "♊", "♋", "♌", "♍", "♎", "♏", "♐", "♑", "♒", "♓", "🈯", "🈳", "🈵", "🈴", "🈲", "🉐", "🈹", "🈺", "🈶", "🈚", "🚾", "🅿", "🈷", "🈸", "🈂", "Ⓜ", "🉑", "㊙", "㊗", "🆑", "🆘", "🆔", "🔞", "🚫", "🆚", "🅰", "🅱", "🆎", "🅾", "❇", "🚻", "🚹", "🚺", "🚼", "🚰", "🚮", "♿", "🚭", "🛂", "🛄", "🛅", "🛃", "🚫", "🔞", "🚯", "🚱", "🚳", "🚷", "🚸", "⛔", "✳", "❇", "❎", "✅", "✴", "💟", "📳", "📴", "💠", "➿", "♻", "⛎", "0⃣", "1⃣", "2⃣", "3⃣", "4⃣", "5⃣", "6⃣", "7⃣", "8⃣", "9⃣", "🔟", "⬆", "⬇", "⬅", "➡", "🔣", "🔢", "🔠", "🔡", "🔤", "↗", "↖", "↘", "↙", "↔", "↕", "🔄", "◀", "▶", "🔼", "🔽", "↩", "↪", "ℹ", "⏪", "⏫", "⏬", "⤵", "⤴", "🆗", "🔀", "🔁", "🔂", "🆕", "🆙", "🆒", "🆓", "🆖", "📶", "🎦", "🈁", "🔯", "🏧", "💹", "💲", "💱", "™", "❌", "❗", "❓", "❕", "❔", "⭕", "🔝", "🔚", "🔙", "🔛", "🔜", "🔃", "🕛", "🕧", "🕐", "🕜", "🕑", "🕝", "🕒", "🕞", "🕓", "🕟", "🕔", "🕠", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕡", "🕢", "🕤", "🕥", "🕦", "➕", "➖", "➗", "♠", "♥", "♣", "♦", "💮", "💯", "✔", "☑", "🔘", "🔗", "➰", "〰", "〽", "🔱", "◼", "🔺", "🔲", "🔳", "⚫", "⚪", "🔴", "🔵", "🔻", "🔶", "🔷", "🔸", "🔹", "✖",
	}

	emojiLen := len(emojiList)
	rand.Seed(time.Now().Unix())
	index := rand.Intn(emojiLen - 1)

	return emojiList[index]
}

//请求接口:url为请求地址，method请求方法：get/post，response_type响应类型：string/json/xml，data:请求参数
func Curl(url string, method string, response_type string, data map[string]interface{}) (result interface{}, err error) {
	var req *httplib.BeegoHTTPRequest = &httplib.BeegoHTTPRequest{}
	if method == "get" {
		req = httplib.Get(url)
	} else if method == "post" {
		req = httplib.Post(url)
	}
	//支持https请求
	req = req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if len(data) > 0 {
		for k, v := range data {
			req = req.Param(k, v.(string))
		}
	}

	if response_type == "json" {
		err = req.ToJSON(&result)
		return result, err
	} else if response_type == "xml" {
		err = req.ToXML(&result)
		return result, err
	}

	result, err = req.String()
	return result, err
}
