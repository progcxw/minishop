package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	uuid "github.com/satori/go.uuid"
)

func String2Int(val string) int {

	goodsId_int, err := strconv.Atoi(val)
	if err != nil {
		return -1
	} else {
		return goodsId_int
	}
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func Int642String(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Float642String(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

func GetUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

//the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func ExactMapValues2Int64Array(maparray []orm.Params, key string) []int64 {

	var vals []int64
	for _, value := range maparray {
		vals = append(vals, value[key].(int64))
	}
	return vals
}

func ExactMapValues2StringArray(maparray []orm.Params, key string) []string {

	var vals []string
	for _, value := range maparray {
		vals = append(vals, value[key].(string))
	}
	return vals
}

type PageData struct {
	NumsPerPage int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

func GetPageData(rawData []orm.Params, page int, size int) PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []orm.Params

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// Hex2Dec : 十六进制转十进制
func Hex2Dec(val string) int64 {
	n, err := strconv.ParseInt(val, 16, 0)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

// Contain : 判断某个元素是否在 slice,array ,map中
func Contain(target interface{}, obj interface{}) (bool, error) {
	targetVal := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		// 是否在slice/array中
		for i := 0; i < targetVal.Len(); i++ {
			if targetVal.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		// 是否在map key中
		if targetVal.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	default:
		fmt.Println(reflect.TypeOf(target).Kind())
	}

	return false, errors.New("not in this array/slice/map")
}
