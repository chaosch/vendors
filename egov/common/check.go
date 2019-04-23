package common

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func CheckIdCard(idNo string) (error) {

	var id_card [18]byte // 'X' == byte(88)， 'X'在byte中表示为88
	var id_card_copy [17]byte

	var id_card_string string

	id_card_string = idNo

	if len(id_card_string) < 18 {
		return errors.New("必须要输入18位的身份证号码")
	}

	// 将字符串，转换成[]byte,并保存到id_card[]数组当中
	for k, v := range []byte(id_card_string) {
		id_card[k] = byte(v)
	}

	//复制id_card[18]前17位元素到id_card_copy[]数组当中
	for j := 0; j < 17; j++ {
		id_card_copy[j] = id_card[j]

		//fmt.Println(byte2int(id_card[j]))
	}
	/*
		fmt.Println(byte2int(id_card[17]))
		fmt.Println(string(id_card[17]))
	*/
	/*
		y := check_id(id_card_copy)
		fmt.Println(y)
	*/
	reg, err := regexp.Compile("^(\\d{6})(\\d{8})(.*)")

	if err != nil {
		return err
	}

	if reg.MatchString(idNo) == true {
		submatch := reg.FindStringSubmatch(idNo)
		t, err := time.ParseInLocation("20060102", submatch[2],time.Local)
		if err != nil {
			return err
		}
		minDate,err:=time.ParseInLocation("20060102", "19170101",time.Local)
		if err!=nil{
			return err
		}
		maxDate,err:=time.ParseInLocation("20060102", "20100101",time.Local)
		if err!=nil{
			return err
		}
		if t.Before(minDate)||t.After(maxDate){
			return errors.New("出生日期不合法")
		}
		//fmt.Println(submatch)
		//fmt.Println("date is ", submatch[2])
	}

	res, str := verify_id(check_id(id_card_copy), byte2int(id_card[17]))
	if !res {
		return errors.New(str)
	}
	return nil
}

func byte2int(x byte) byte {
	if x == 88 {
		return 'X'
	}
	return (x - 48) // 'X' - 48 = 40;
}

func check_id(id [17]byte) int {
	arry := make([]int, 17)

	//强制类型转换，将[]byte转换成[]int ,变化过程
	// []byte -> byte -> string -> int
	//将通过range 将[]byte转换成单个byte,再用强制类型转换string()，将byte转换成string
	//再通过strconv.Atoi()将string 转换成int 类型
	for index, value := range id {
		arry[index], _ = strconv.Atoi(string(value))
	}
	/*
		for p := 0; p < len(arry); p++ {
			fmt.Println("arry[", p, "]", "=", arry[p])
		}
	*/

	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		//fmt.Println("id =", i, byte2int(id[i]), wi[i])
		res += arry[i] * wi[i]
	}

	//fmt.Println("res = ", res)

	return (res % 11)
}

func verify_id(verify int, id_v byte) (bool, string) {
	var temp byte
	var i int
	a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}

	for i = 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			//fmt.Println("verify_id id",)
			// if a18[i] == 'X' ,let convert it to type string
			if (a18[i] == 88 ) {
				//fmt.Println("计算得到身份证最后一位是 ", string(a18[i]))
			} else {
				//fmt.Println("计算得到身份证最后一位是 ", a18[i])
			}
			//fmt.Println(i, temp)
			break
		}
	}
	//if id_v == 'X', let's convert it to type string
	//if (id_v == 88) {
	//	fmt.Println("身份证最后一位是 ", string(id_v))
	//} else {
	//	fmt.Println("身份证最后一位是  ", id_v) // id_v是身份证的最后一位
	//}

	if temp == id_v {

		return true, "验证成功"
	}

	return false, "验证失败"
}

const (
	black=30
	red=31
	green=32
	yellow=33
	blue=34
	pink=35
	cyan=36
	white=37
)

func Println(foreaground int,format string, a ...interface{}){
	fmt.Printf("%c[0;0;%dm%s%c[0m\n",0x1B,foreaground,fmt.Sprintf(format,a...), 0x1B)
}

func PrintRed(format string, a ...interface{}){
	Println(red,format,a...)
}

func PrintGreen(format string, a ...interface{}){
	Println(green,format,a...)
}

func Print(format string, a ...interface{}){
	Println(0,format,a...)
}

