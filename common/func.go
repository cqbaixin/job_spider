package common

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"time"
)

func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func StrToUtf8(str string)(string, error ){
	b, err := GbkToUtf8([]byte(str))
	if err != nil {
		println(err.Error())
		return "",err
	}
	str = string(b)
	return  str,nil
}

// transform UTF-8 string to GBK string and replace it, if transformed success, returned nil error, or died by error message
func StrToGBK(str *string) error {
	b, err := Utf8ToGbk([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return  nil
}


func GetMaxMinSalary(str string)(int,int,error){
	byteStr := []rune(str)
	strLast := byteStr[len(byteStr)-1:]
	strDay := byteStr[len(byteStr)-3]
	mouthNum := 1000.0
	isYear := 0
	regSalary := regexp.MustCompile(`(\-|\+)?\d+(\.\d+)?`)
	res := regSalary.FindAllString(str,-1)
	if(len(str)<=0){
		return  0,0,errors.New("获取工资失败,失败字符串是:"+str)
	}
	min,_ := strconv.ParseFloat(res[0], 64)
	max,_ :=  strconv.ParseFloat(res[1], 64)
	if(string(strLast) == "年"){
		isYear = 1
	}
	if(string(strDay) == "万"){
		mouthNum = 10000
	}
	if(isYear == 1){
		min = min * (mouthNum) /12
		max = max * (mouthNum) /12
	}else {
		min = min * (mouthNum)
		max = max * (mouthNum)
	}
	return int(math.Abs(min)),int(math.Abs(max)),nil
}


func GetHBMaxMinSalary(str string)(int,int,error){
	regSalary := regexp.MustCompile(`(\-|\+)?\d+(\.\d+)?`)
	res := regSalary.FindAllString(str,-1)
	if(len(str)<=0){
		return  0,0,errors.New("获取工资失败,失败字符串是:"+str)
	}
	min,_ := strconv.ParseFloat(res[0], 64)
	max,_ :=  strconv.ParseFloat(res[1], 64)
	return int(math.Abs(min)),int(math.Abs(max)),nil
}

func GetHBPublishAT(str string)time.Time  {
	var t time.Time
	var err error
	switch str {
	case "今天":
		t,err = time.Parse("20060102",time.Now().Format("20060102"))
	case "昨天":
		t,err = time.Parse("20060102",time.Now().AddDate(0,0,-1).Format("20060102"))
	default:
		t,err = time.Parse("2006/01/02",fmt.Sprintf("%d/%s",time.Now().Year(),str))
	}
	if(err != nil){
		t,_ = time.Parse("20060102",time.Now().Format("20060102"))
	}
	return t
}