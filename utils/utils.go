package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	okxBaseUrl   = "https://www.okx.com"
	okxApiId     = "479d5408200921213da5fd233c94e49b"
	okxApiKey    = "4cd88099-ca98-4c4c-9b0a-b6e3d6dbcfa5"
	okxApiSecret = "6143898B9549E37DC3A0DE11A26C9602"

	okxApiPassphrase = "Asd123456,./"
)

func OkxHttp(pathUrl string) {
	now := time.Now().UTC()
	timestamp := now.Format(time.RFC3339)
	prehashString := timestamp + "GET" + pathUrl
	// 4. 使用 HMAC SHA256 对预哈希字符串进行签名
	signature, err := generateSignature(prehashString, okxApiSecret)
	if err != nil {
		fmt.Println("Error generating signature:", err)
		return
	}
	// 5. 构建 HTTP 请求头
	headers := make(map[string]string)
	headers["OK-ACCESS-PROJECT"] = okxApiId
	headers["OK-ACCESS-KEY"] = okxApiKey
	headers["OK-ACCESS-SIGN"] = signature
	headers["OK-ACCESS-TIMESTAMP"] = timestamp
	headers["OK-ACCESS-PASSPHRASE"] = okxApiPassphrase
	fmt.Println(signature)
	fmt.Println(timestamp)
	url := okxBaseUrl + pathUrl

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	// 7. 处理响应
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(responseData))
}

func generateSignature(data, secretKey string) (string, error) {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func HmacSha256crypto(data, secretKey string) (string, error) {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}

func ConvertToEth(value *big.Int, decimals int) string {
	v := new(big.Float)
	v.SetInt(value)
	ethValue := new(big.Float).Quo(v, big.NewFloat(math.Pow10(decimals)))
	return ethValue.Text('f', decimals)
}
func ConvertStrToEth(value string, decimals int) string {
	v := new(big.Float)
	v.SetString(value)
	ethValue := new(big.Float).Quo(v, big.NewFloat(math.Pow10(decimals)))
	return ethValue.Text('f', decimals)
}

func ConvertStrToWei(value string, decimals int) string {
	v := new(big.Float)
	v.SetString(value)
	ethValue := new(big.Float).Mul(v, big.NewFloat(math.Pow10(decimals)))
	return ethValue.Text('f', 0)
}

func FormatGasPrice2Eth(value string, gasLimit uint64, convertDecimals int, formatDecimals int) string {
	v := new(big.Float)
	v.SetString(value)
	limit := new(big.Float)
	limit.SetUint64(gasLimit)
	ethValue := new(big.Float).Quo(v, big.NewFloat(math.Pow10(convertDecimals)))
	gasPrice := new(big.Float).Mul(ethValue, limit)
	return gasPrice.Text('f', formatDecimals)
}

func ConvertEth2Usd(value string, usd string) string {
	v := new(big.Float)
	v.SetString(value)
	u := new(big.Float)
	u.SetString(usd)
	res := new(big.Float).Mul(v, u)
	return res.Text('f', 6)
}

func AddPrefixIfNeeded(input string) string {
	if !strings.HasPrefix(input, "0x") {
		return "0x" + input
	}
	return input
}

func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func cryptoPwd(username, password string) string {
	mac := hmac.New(sha256.New, []byte(username))
	mac.Write([]byte(password))
	expectedMAC := mac.Sum(nil)
	fmt.Println(fmt.Sprintf("%x", expectedMAC))
	return fmt.Sprintf("%x", expectedMAC)
}

func GetFormatTime() string {
	currentTime := time.Now().Local()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

func GetFormatTime2Local(time2 time.Time, formatStr string) string {
	currentTime := time2.Local()
	formattedTime := currentTime.Format(formatStr)
	return formattedTime
}

func GetUnixTime2Local(time2 time.Time) time.Time {
	currentTime := time2.Local()
	//formattedTime := currentTime.Unix()
	return currentTime
}

func GetTodayTime() string {
	return time.Now().Local().Format(time.DateOnly)
}

func GetTodayMonthTime() string {
	return time.Now().Local().Format("2006-01")
}

// 时间戳 ms  返回 年月日
func GetDayTimeByParams(mss int64) string {
	ftStr := fmt.Sprintf("%d", mss)
	var ft time.Time
	if len(ftStr) == 10 {
		ft = time.Unix(mss, 0)
	} else {
		ft = time.UnixMilli(mss)
	}

	const layout = "2006-01-02"
	return ft.Format(layout)
}

func GetMonthTime() string {
	now := time.Now().Local()
	year, month, _ := now.Date()
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	return firstDayOfMonth.Format("2006-01")
}

func GetYesterdayTime() string {
	now := time.Now().Local()
	yesterday := now.AddDate(0, 0, -1)
	return yesterday.Format(time.DateOnly)
}

func Obj2String(obj interface{}) (string, error) {
	bt, err := json.Marshal(obj)
	if err != nil {
		//fmt.Errorf("obj转换错误!")
		return "", errors.New("obj转换错误")
	} else {
		return string(bt), nil
	}
}

// GMT-4 时区的时间
func GetGMTFTime(p int64) string {
	t := time.Unix(p, 0)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	newYorkTime := t.In(loc)
	return newYorkTime.Format("2006-01-02T15:04:05")
}

// 格式化时间 格式化时间YYYY-MM-DDTHH:mm:ss.SSSZ
func Ftime2UTC(t int64) string {

	t1 := time.Unix(t, 0)
	//t1 := time.Unix(unixTime-1*60*60*24, 0)
	//t2 := time.Unix(unixTime, 0)
	//startDate := t1.UTC().Format("2006-01-02T15:04:05.999Z")
	//endDate := t2.UTC().Format("2006-01-02T15:04:05.999Z")
	tt := t1.UTC()
	return tt.Format("2006-01-02T15:04:05.999Z")
}

// UTC-4 :2024-03-13T02:52:43-04:00
func GetUTCFTime(p int64) string {

	tb := time.Unix(p, 0)
	// 将时间对象转换为 UTC-4 时区
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	timeInUTCMinus4 := tb.In(loc)
	// 格式化输出时间
	return timeInUTCMinus4.Format(time.RFC3339)
}

func GetUTCFTimeEVO(p int64) string {
	tb := time.Unix(p, 0).UTC()
	// 格式化输出时间
	return tb.Format("2006-01-02T15:04:05.999Z")
}

// 将逗号分隔的字符串转换为整数切片
func StringToIntSlice(str string) ([]int, error) {
	strList := strings.Split(str, ",")
	intList := make([]int, len(strList))
	for i, s := range strList {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intList[i] = num
	}
	return intList, nil
}

// 判断数组中是否 存在指定元素
func IsContainsInList(list *[]any, target any) bool {
	for _, v := range *list {
		if v == target {
			return true
		}
	}
	return false
}

// 格式化时间 转时间戳 key:time.RFC3339
func Ftime2Unix(key string, value string) int64 {

	//layout := "2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(key, value)
	if err != nil {
		fmt.Println("解析时间字符串出错:", err)
		return 0
	}
	return t.Unix()
}

// 格式化时间 转时间戳 key:time.RFC3339 (毫秒)
func Ftime2UnixMilli(key string, ft string) int64 {
	t, err := time.Parse(key, ft)
	if err != nil {
		fmt.Println("解析时间字符串出错:", err)
		return 0
	}
	return t.UnixMilli()
}

// 数组转字符串 带[]
func Arr2Str(arrs []interface{}) string {
	return strings.Join(strings.Fields(fmt.Sprint(arrs)), ",")
}

func MaskString(s string) string {
	length := len([]rune(s))
	if length <= 2 {
		return string([]rune(s)[0]) + "*"
	}

	masked := []rune(s)
	var pad = min(3, max(1, length/3)) //两边保留字符 1-3个字符
	pad = max(1, min(length/2-1, pad))
	for i := pad; i < length-pad; i++ {
		masked[i] = '*'
	}

	return string(masked)
}

var defaultPriority = 0

func MakeGameTypePriority(gameid string) string {
	return fmt.Sprintf("{ \"%s\" : %d }", gameid, defaultPriority)
}

func TruncateTransformInAmount(amount float64) float64 {
	// 将数乘以100，使用Floor函数向下取整，然后再除以100，从而保留两位小数
	truncatedNum := math.Floor(amount*100) / 100
	return truncatedNum
}

// stack 返回当前的堆栈跟踪信息
func stack() string {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	return string(buf)
}

type IPInfo struct {
	Country string `json:"country"`
}

func GetCountryByIP(ip string) (string, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	// 发起HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}

	// 解析JSON数据
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}
	return ipInfo.Country, nil
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func GenerateAmigoUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("错误:", err)
		return GenerateAmigoUuid()
	}
	// 将字节数组转换为32位的十六进制字符串
	id := hex.EncodeToString(b)
	return id
}

func ValidateWhiteList(list string, sep string, white string) bool {
	if list == "" {
		return true
	}
	strlist := strings.Split(list, sep)
	// 检查白名单
	for _, str := range strlist {
		if white == str {
			return true
		}
	}

	return false
}
func ConvertToBeijingTime(timestampMillis int64) time.Time {
	// Convert milliseconds to seconds
	seconds := timestampMillis / 1000
	// Convert to time.Time object
	t := time.Unix(seconds, 0)
	// Define the Beijing timezone
	beijingTime, _ := time.LoadLocation("Asia/Shanghai")
	// Convert time to Beijing time
	return t.In(beijingTime)
}

func SHA256(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
