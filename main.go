package main

import (
	"fmt"
	// "reflect"
	"os"
	"encoding/json"
	"sync"
	"github.com/gin-gonic/gin"
	"net/http"
	"backEnd/utils"
)

type menuInfo struct {
	MenuInfo []interface{} `json:"database"`
}

var mutex sync.Mutex

func loadFileTest(filePath string, v menuInfo) {
	mutex.Lock()
	data, err := os.ReadFile(filePath)	
	mutex.Unlock()
	if err != nil {
		fmt.Println(1)
		return
	}
	err = json.Unmarshal([]byte(data), &v)
  	if err != nil {
		fmt.Println(err)
    	return
  	}
	fmt.Println(v.MenuInfo)
}

func loadFile(filePath string, pageName string) (interface{}, bool) {
	mutex.Lock()
	file, err := os.ReadFile(filePath)
	mutex.Unlock()
	if err != nil {
		fmt.Println("读取文件错误")
		return nil, false
	}
	var menuList map[string]interface{}
	err = json.Unmarshal(file, &menuList)
	if err != nil {
		fmt.Println("解码json错误")
		fmt.Println(err)
		return nil, false
	}
	// fmt.Printf("value: %+v\n", newMenu)
	// var menu map[string]interface{}
	menu := menuList[pageName]
	// fmt.Println("menu: ")
	// fmt.Println(menu)
	return menu, true
	// fmt.Println(reflect.TypeOf(newMenu))
}
func GetPageName(ctx *gin.Context) {
	// 与前端约定好字符串
	pageName := ctx.Query("pageKind")
	fmt.Println(pageName)
	if len(pageName) == 0 {
		ctx.String(http.StatusBadRequest,"未得到页面名称，无法返回菜单配置")
		return
	}
	menu, err := loadFile("./data/config.json", pageName)
	if !err {
		fmt.Println("加载文件失败")
		return 
	}
	fmt.Println(menu)
	ctx.JSON(http.StatusOK, menu)
}
func main() {
	// _, err1 := loadFile("./data/config.json", "database")
	
	engine := gin.Default()
	engine.Use(utils.CORS())
	engine.GET("/getPageName", GetPageName)
	err := engine.Run("0.0.0.0:2345")
	if err != nil {
		panic(err)
	}

	// v := menuInfo{}
	// loadFileTest("./database/config.json", v)

}