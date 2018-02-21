package model

import "fmt"

type Profile struct {
	Name       string
	Gender     string // 性别
	Age        int
	Height     int
	Weight     int
	Income     string // 收入
	Marriage   string
	Education  string
	Occupation string
	Hokou      string // 户口
	Xinzuo     string
	House      string
	Car        string
}

func (p Profile) strings() string {
	str := fmt.Sprintf("姓名: %v 性别: %v 年龄: %v 身高: %v 体重: %v 收入: %v 婚姻状况: %v 学历: %v",
		p.Name, p.Gender, p.Age, p.Height, p.Weight, p.Income, p.Marriage, p.Education)
	return str
}
