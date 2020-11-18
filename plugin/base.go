package plugin

import (
	"bytes"
	"crypto/rand"
	"github.com/ChengjinWu/gojson"
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

func TimeTostring(t time.Time) (s string) {
	s = t.Format("2006-01-02 15:04:05")
	return
}

//StrToInt64
func StrToInt64(str string) int64 {
	i, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		log.Println(e)
		return 0
	}
	return i
}

//StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		log.Println(e)
		return 0
	}
	return i
}

//StrToUInt string 转int
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// StrToTime 字符串转time
func StrToTime(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(timeLayout, s, loc)
	return t
}

// StrToTime RFC3339
func StrRFC3339ToTime(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02T15:04:05Z07:00", s, loc)
	return t, err
}

// StrToDate 字符串转time
func StrToDate(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02"
	t, _ := time.ParseInLocation(timeLayout, s, loc)
	return t
}

//StrToTimePtr ..
func StrToTimePtr(str string) *time.Time {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return nil
	}
	t, e := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if e != nil {
		return nil
	}
	return &t
}

//IntToStr
func IntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

//Int64ToStr
func Int64ToStr(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func JsonCheck(data []byte) (err error) {
	err = gojson.CheckValid([]byte(data))
	return
}

// RandString 生成随机字符串
func CreateRandomStringWithTime(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	t_str := Int64ToStr(time.Now().Unix())
	return t_str + container
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 创建二维码
func CreateErweima(str string, path string) error {
	err := qrcode.WriteFile(str, qrcode.Medium, 256, path)
	return err
}

// 写文件
func WriteToFile(msg string, filePath string) (err error) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return
	} else {
		_, err = f.Write([]byte(msg))
	}
	return
}

//读文件
func ReadFile(filePath string) (contentByte []byte, err error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		return
	}
	contentByte, err = ioutil.ReadAll(f)
	return
}

//map get keys type int64
func MapGetKeysInt64(m map[int64]int64) []int64 {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]int64, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

//map get keys type string
func MapGetKeysString(m map[string]interface{}) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
