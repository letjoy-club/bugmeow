package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"bugmeow/biz"

	lark "github.com/larksuite/oapi-sdk-go/v3"

	"github.com/gin-gonic/gin"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"

	"github.com/sbecker/gin-api-demo/util"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	content string `json:"content" binding:"required"`
}

type Event struct {
	message Message `json:"message" binding:"required"`
}

type Challenge struct {
	event interface{} `json:"event" binding:"required"`
}

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)

		entry := log.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    util.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func send(client *lark.Client, text string, chatId string) error {
	body := map[string]string{
		"receive_id": chatId,
		"msg_type":   "text",
		"content":    text,
	}
	res, err := client.Do(context.Background(),
		&larkcore.ApiReq{
			HttpMethod:                http.MethodPost,
			ApiPath:                   "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=chat_id",
			Body:                      body,
			SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant},
		},
	)
	if !strings.Contains(string(res.RawBody), "\"code\":0") || err != nil {
		return errors.New("发送失败")
	}
	return nil
}

func addRecord(client *lark.Client, text string, reporter string, processor string) error {
	status := "待定"

	if strings.HasPrefix(text, "【功能】") {
		status = "功能"
	}

	// 标签处理优化
	if strings.HasPrefix(text, "【优化】") {
		status = "优化"
	}
	content := map[string]interface{}{"bug内容": text, "状况": status, "报告人": []map[string]string{map[string]string{"id": reporter}}}
	if processor != "" {
		person := map[string]string{
			"id": processor,
		}
		people := [1]map[string]string{person}
		content["经办人"] = people
	}
	body := map[string]interface{}{
		"fields": content,
	}
	res, err := client.Do(context.Background(),
		&larkcore.ApiReq{
			HttpMethod:                http.MethodPost,
			ApiPath:                   "https://open.feishu.cn/open-apis/bitable/v1/apps/bascnAH9f5XWLhVt7HG432D1eLj/tables/tbl9QXXSKVkbUHCU/records",
			Body:                      body,
			SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant},
		},
	)

	if !strings.Contains(string(res.RawBody), "\"code\":0") || err != nil {
		return errors.New("填写表格失败")
	}
	return nil
}

func main() {
	r := gin.Default()
	r.Use(JSONLogMiddleware())
	r.Use(gin.Recovery())
	r.Use(biz.Handler())

	// 在已有 Gin 实例上注册消息处理路由
	r.POST("/", func(c *gin.Context) {
		// var tmp biz.ReceiveMessageEvent

		// if err := c.ShouldBindJSON(&tmp); err != nil {
		// 	c.JSON(400, gin.H{"code": -1})
		// 	return
		// }

		// raw := tmp.Event.Message.Content
		// var tmpMap map[string]string
		// if err := json.Unmarshal([]byte(raw), &tmpMap); err != nil {
		// 	c.JSON(400, gin.H{"code": -2})
		// 	return
		// }

		// // Lark 获得的信息
		// base := tmpMap["text"]
		// head := ""
		// startIndex := strings.Index(base, " ")
		// if startIndex != -1 {
		// 	head = base[:startIndex]
		// 	base = strings.Trim(base[startIndex:], " ")
		// }

		// // sender 给所有人发的时候 mute
		// if strings.Contains(head, "@_all") {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code":  0,
		// 		"error": "",
		// 	})
		// 	return
		// }

		// rawask := strings.Split(base, "@")
		// sender := tmp.Event.Sender.SenderID
		// place := tmp.Event.Message.ChatID
		// rawText := ""
		// rawProcessor := "" // 经办人 Id
		// if len(tmp.Event.Message.Mentions) > 1 {
		// 	rawProcessor = tmp.Event.Message.Mentions[1].ID.OpenID
		// }

		// fmt.Println(rawask)
		// // 处理的结果
		// resp := ""

		// // 创建client
		// client := lark.NewClient("cli_a327a22b1eb8500d", "aHcbZ75HvEj2dhfm7z4SUceeMIk5PotG")

		// if len(rawask) < 3 && len(rawask) > 0 {
		// 	resp = "OK，已经记录"
		// 	rawText = rawask[0]
		// } else {
		// 	resp = "bug 提交请按照 @xxx 问题 或 @xxx 问题 @经办人 的格式，其他格式暂不支持"
		// }

		// fmt.Println("rawText", rawText)

		// // 指令处理
		// if rawText != "" {
		// 	if rawText == "help" {
		// 		// 帮助
		// 		resp = "docs: 展示文档；xxx 问题 或 @xxx 问题 @经办人 添加记录"
		// 	} else if rawText == "docs" {
		// 		// 文档
		// 		resp = "https://zdjb5i2gev.feishu.cn/base/bascnAH9f5XWLhVt7HG432D1eLj?table=tbl9QXXSKVkbUHCU&view=vewhQYpSaw"
		// 	} else {
		// 		// 添加
		// 		if err := addRecord(client, rawText, sender["open_id"], rawProcessor); err != nil {
		// 			c.JSON(http.StatusOK, gin.H{
		// 				"code":  "-3",
		// 				"error": err,
		// 			})
		// 			return
		// 		}
		// 	}
		// }

		// response := map[string]string{
		// 	"text": "<at user_id=\"" + sender["open_id"] + "\">Tom</at> " + resp,
		// }

		// res_text, _ := json.Marshal(response)

		// if err := send(client, string(res_text), place); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code":  "-2",
		// 		"error": err,
		// 	})
		// 	return
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Hello world",
		// })
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		if file != nil {
			c.JSON(http.StatusOK, gin.H{
				"filename": file.Filename,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "file is not found",
			})
		}

	})
	r.POST("/challenge", func(c *gin.Context) {
		var tmp biz.ChallengeMessageEvent

		if err := c.ShouldBindJSON(&tmp); err != nil {
			c.JSON(400, gin.H{"code": -1})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"challenge": tmp.Challenge,
			"token":     tmp.Token,
			"type":      tmp.Type,
		})
	})

	r.Run(":8000")
}
