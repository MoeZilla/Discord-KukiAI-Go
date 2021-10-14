package kuki

import (
	"fmt"
	"time"
        "net/url"

	"github.com/bwmarrin/discordgo"

	"github.com/valyala/fasthttp"
       
        "github.com/valyala/fastjson"
)

var client = &fasthttp.Client{}


func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return 
	}
	response := kequest("GET", fmt.Sprintf("http://kukiapi.xyz/api/apikey=%s/%s/%s/messags=%s", kukikey, botname, owner name, url.QueryEscape(m.Content)))
	kk, err := fastjson.Parser.Parse(string(response.Body()))
	if err != nil {
		session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("json data", err.Error()))
		return
	}
	session.ChannelMessageSend(msg.ChannelID, string(kk.GetStringBytes("reply")))
}

func kequest(method, uri string) *fasthttp.Response {
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(method)
	request.Header.SetRequestURI(uri)
	response := fasthttp.AcquireResponse()
	err := Client.DoTimeout(request, response, 15*time.Second)
	if err != nil {
		fmt.Printf("https error", err.Error())
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
		return nil
	}
	return response
}

