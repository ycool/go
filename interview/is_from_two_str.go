package main

import (
    "fmt"
)

func IsFromTwoString(str1, str2, str3 string) bool {
	if len(str1) + len(str2) != len(str3) {
		return false
	}
	if len(str1) == 0 && len(str2) == 0 && len(str3) == 0 {
		return true
	}
	if len(str1) > 0 && len(str2) > 0 && str1[0] == str2[0] && 
		str1[0] == str3[0] {
		return IsFromTwoString(str1[1:], str2, str3[1:]) || 
			IsFromTwoString(str1, str2[1:], str3[1:]) 
	}
	if len(str1) > 0 && str1[0] == str3[0] {
		return IsFromTwoString(str1[1:], str2, str3[1:])
	}
	if len(str2) > 0 && str2[0] == str3[0] {
		return IsFromTwoString(str1, str2[1:], str3[1:])
	}
	return false
}

func main() {
	var s1 string = "abc"
	var s2 string = "def"
	var s3 string = "abcdef"

    fmt.Println(s1, ",", s2, ",", s3, ":", IsFromTwoString(s1, s2, s3))
    s3 = "adbecf"
    fmt.Println(s1, ",", s2, ",", s3, ":", IsFromTwoString(s1, s2, s3))
    s3 = "abdecf"
    fmt.Println(s1, ",", s2, ",", s3, ":", IsFromTwoString(s1, s2, s3))
    s3 = "adefbc"
    fmt.Println(s1, ",", s2, ",", s3, ":", IsFromTwoString(s1, s2, s3))
    s3 = "abcdeg"
    fmt.Println(s1, ",", s2, ",", s3, ":", IsFromTwoString(s1, s2, s3))
}
 