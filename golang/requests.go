package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type info struct {
	PostId int    `json:"PostId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// type info struct {
// 	postId int
// 	id     int
// 	name   string
// 	email  string
// 	body   string
// }

type infos []info
type resInfo struct {
	infos
}

func GetUrl(id int) string {
	Uid := strconv.Itoa(id)
	return "https://jsonplaceholder.typicode.com/posts/" + Uid + "/comments"
}

func GetEmail(url string) ([]string, error) {
	var bodyInfo resInfo
	var email []string
	client := &http.Client{}
	res, _ := http.NewRequest(http.MethodGet, url, nil)
	res.Header.Add("Content-Type", "application/json")
	allInfo, err := client.Do(res)
	if err == nil {
		defer allInfo.Body.Close()
		body, err := ioutil.ReadAll(allInfo.Body)
		_ = json.Unmarshal(body, &bodyInfo)
		if err == nil {
			// TODO:根据转换的格式获取email信息，结构体书写有问题，暂无法获取
			// for _, i := range body {
			// 	email = append(email, i.email)
			// }
			// return email, nil
			return append(email, "123@test.com"), nil
		} else {
			return nil, errors.New("获取信息失败")
		}
	}
	return nil, errors.New("请求失败")
}

func main() {
	var emails [][]string
	for i := 1; i < 101; i++ {
		url := GetUrl(i)
		email, err := GetEmail(url)
		if err != nil {
			fmt.Println(err)
		}
		emails = append(emails, email)
	}
	fmt.Println(emails)
}
