package funcs

import (
	"bot/botTool"
	. "bot/config"
	"bot/dbManager"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var config = dbManager.InitMysql("config", CONFIG_TOKEN, "config")

func GetMessgae(update *tgbotapi.Update) {
	if update.Message != nil {
		config.AddGroup(update, fmt.Sprint(update.Message.Chat.ID), update.Message.Chat.UserName, update.Message.Chat.Title, fmt.Sprint(update.Message.From.ID), update.Message.From.UserName, botTool.GetName(update))
	}
}

func getHistory(body *[]byte, date ...string) {
	var resp *http.Response
	if len(date) == 0 {
		resp, _ = http.Get("http://hao.360.cn/histoday")
	} else {
		resp, _ = http.Get(fmt.Sprintf("http://hao.360.cn/histoday/%s%s.html", date[0], date[1]))
	}
	defer resp.Body.Close()
	*body, _ = io.ReadAll(resp.Body)
}

func httpfix(url string) string {
	url = strings.TrimSpace(url)
	if url[0:4] != "http" || (url[5:8] != "://" && url[4:7] != "://") {
		url = "http://" + url
	}
	return url
}

func goHttp(url string, ip, port string, wg ...*sync.WaitGroup) (b string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println(string(debug.Stack()))
		}
	}()
	if wg != nil {
		defer wg[0].Done()
	}
	resp, err := http.Get(fmt.Sprintf("%s?ip=%s&port=%s", url, ip, port))
	if err != nil {
		log.Println("goHttp", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("GoHttp", err)
		return
	}
	b = string(body)

	return
}

func find(text string, arr [14]string) bool {
	for _, i := range arr {
		if text == i {
			return true
		}
	}
	return false
}

func index(text string, arr [14]string) int {
	for i, v := range arr {
		if text == v {
			return i
		}
	}
	return 13
}

func getToMap(url string) (res map[string]interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("GetTomap", err)
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&res)
	return
}

func getCoin(update *tgbotapi.Update, coinType string) {
	text := getToMap(fmt.Sprintf("https://api.huobi.pro/market/history/kline?symbol=%susdt&period=1min&size=1", coinType))
	msg := "正在查询，plz wait..."
	msgConfig, _ := botTool.SendMessage(update, &msg, true)
	msg = fmt.Sprintf("啊哈哈哈哈哈哈\n价格来咯！\n1.0 %s = %.2f USD", strings.ToUpper(coinType), text["data"].([]interface{})[0].(map[string]interface{})["open"].(float64))
	botTool.Edit(msgConfig, &msg)
}

func getReplyAt(update *tgbotapi.Update) string {
	return fmt.Sprintf("[%s](tg://user?id=%d)", botTool.GetReplyName(update), update.Message.ReplyToMessage.From.ID)
}

func getAt(update *tgbotapi.Update) string {
	return fmt.Sprintf("[%s](tg://user?id=%d)", botTool.GetName(update), update.Message.From.ID)
}
