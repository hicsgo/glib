package glib

import (
	"strings"
	"reflect"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 转换格式 eg:将user_id转换成UserId
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ConvertName(mysqlName string) string {
	sqlNames := StringToStringSlice(mysqlName, "_")
	goName := ""
	for _, sqlName := range sqlNames {
		goName += FirstToUpper(sqlName)
	}
	return goName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串分隔源字符串为字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToStringSlice(sourceString string, args ...string) []string {
	result := make([]string, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将字符串首字母大写
 * eg:(user_id转换成User_id)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] -= 32
			upperStr += string(vv[i]) // + string(vv[i+1])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}


/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将字符串首字母小写 eg:(User_id转换成user_id)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FirstToLower(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] += 32
			upperStr += string(vv[i]) // + string(vv[i+1])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据mysql类型返回go对应的类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GoTypeByMysqlType(mysqlTye string) interface{} {
	var goType interface{}
	switch mysqlTye {
	case "bool":
		goType = reflect.Bool
		break
	case "varchar":
		goType = reflect.String
		break
	case "text":
		goType = reflect.String
		break
	case "longtext":
		goType = reflect.String
		break
	case "char":
		goType = reflect.String
		break
	case "date":
		goType = reflect.Int32
		break
	case "datetime":
		goType = reflect.Struct
		break
	case "time":
		goType = reflect.Struct
		break
	case "tinyint":
		goType = reflect.Int64
		break
	case "smallint":
		goType = reflect.Int64
		break
	case "decimal":
		goType = reflect.Float64
		break
	case "int":
		goType = reflect.Int64
		break
	default:
		goType = reflect.String
		break
	}
	return goType
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据bitString返回bool Slice
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func BitStringToBoolSlice(str string) []bool {
	s := make([]bool, 0, len(str)*8)
	for i := 0; i < len(str); i++ {
		for bit := 7; bit >= 0; bit-- {
			set := (str[i]>>uint(bit))&1 == 1
			s = append(s, set)
		}
	}
	return s
}

