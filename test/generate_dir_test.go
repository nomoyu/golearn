package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var rootDir string
var separator string
var jsonData map[string]any

const fileName = "dir.json"

/*加载json文件*/
func loadJson() {
	separator = string(filepath.Separator)
	workDir, _ := os.Getwd()
	rootDir = workDir[:strings.LastIndex(workDir, separator)]
	jsonBytes, _ := os.ReadFile(workDir + separator + fileName)
	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

/*解析json中的map 解析路径*/
func parseMap(data map[string]any, dir string) {
	for k, v := range data {
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if dir != "" {
					path := dir + separator + path
					if k == "text" {
						dir = path
					}
					createDir(path)
				} else {
					dir = path
					createDir(path)
				}

			}
		case []any:
			{
				parseArray(v.([]any), dir)
			}
		}
	}

}

/*解析路径，处理子路径*/
func parseArray(data []any, dir string) {
	for _, v := range data {
		mapv, _ := v.(map[string]any)
		parseMap(mapv, dir)
	}
}

/*创建文件目录*/
func createDir(path string) {
	if path == "" {
		return
	}
	err := os.MkdirAll(rootDir+separator+path, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())
	}
	fmt.Println(path)
}

/*测试函数入口*/
func TestGenerateDir(t *testing.T) {
	// 初始化目录
	//loadJson()
	//parseMap(jsonData, "")

	//jwt测试
	TestJwt()
}
