package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func gethubdns() {
	hostsPath := `C:\Windows\System32\drivers\etc\hosts`
	url := "https://raw.hellogithub.com/hosts"
	startMarker := "# GitHub520 Host Start"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求异常:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应数据失败:", err)
			return
		}

		file, err := os.OpenFile(hostsPath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return
		}
		defer file.Close()

		lines, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("读取文件内容失败:", err)
			return
		}

		content := string(lines)
		index := strings.Index(content, startMarker)
		if index == -1 {
			index = len(content)
		}

		newContent := content[:index] + string(data)
		err = file.Truncate(0) // 清空文件内容
		if err != nil {
			fmt.Println("清空文件内容失败:", err)
			return
		}

		_, err = file.Seek(0, 0) // 重置文件指针
		if err != nil {
			fmt.Println("重置文件指针失败:", err)
			return
		}

		_, err = file.WriteString(newContent) // 写入新内容
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}

		fmt.Println("写入成功")
	} else {
		fmt.Println("写入失败：响应状态码", resp.StatusCode)
	}
}

func main() {
	gethubdns()
}
