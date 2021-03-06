package app

import (
	"fmt"
	"encoding/json"
	"Hanfu/utils"
	"Hanfu/conf"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Code2Session(c *gin.Context) {
	type JSONData struct {
		Code string `form:"code" binding:"required"`
		Version string `form:"version" binding:"required"`
	}
	var data JSONData
	if err := c.ShouldBind(&data); err != nil {
		utils.RES(c, utils.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	code:=data.Code
	// fmt.Println(code)

	//生成client 参数为默认
	var url=""
	if data.Version =="qq"{
		url="https://api.q.qq.com/sns/jscode2session?appid="+conf.QQConfig.APP_ID+"&secret="+conf.QQConfig.APP_SECRET+"&js_code="+code+"&grant_type=authorization_code"
	}else{
		url="https://api.weixin.qq.com/sns/jscode2session?appid="+conf.WXConfig.APP_ID+"&secret="+conf.WXConfig.APP_SECRET+"&js_code=" + code + "&grant_type=authorization_code"
	}
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		utils.RES(c, utils.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer  resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.RES(c, utils.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	// fmt.Println(string(body))

	myMap:=make(map[string]string)
	json.Unmarshal([]byte(body),&myMap)
	fmt.Println(myMap["openid"])
	
	info:=gin.H{
		"openid":myMap["openid"],
	}

	utils.RES(c, utils.SUCCESS,  gin.H{
		"info":info,
	})
}