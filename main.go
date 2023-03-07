package main

import (
	"bookget/config"
	"bookget/router"
	"bufio"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()

	//配置初始化
	if !config.Init(ctx) {
		os.Exit(0)
	}

	//单个URL
	if config.Conf.DUrl != "" {
		ExecuteCommand(ctx, 1, config.Conf.DUrl)
		log.Print("Download complete.\n")
		return
	}

	//批量URLs
	if config.Conf.UrlsFile != "" {
		//加载配置文件
		bs, err := os.ReadFile(config.Conf.UrlsFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		mUrls := strings.Split(string(bs), "\n")
		iCount := 0
		for i, sUrl := range mUrls {
			if sUrl == "" {
				continue
			}
			ExecuteCommand(ctx, i+1, sUrl)
			iCount++
		}
		log.Print("Download complete.\n")
		log.Printf("下载完成，共 %d 个任务，请到 %s 目录下查看。\n", iCount, config.Conf.SaveFolder)
		return
	}

	iCount := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter an URL:")
		fmt.Print("-> ")
		sUrl, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Printf("Error: %w \n", err)
			break
		}
		iCount++
		ExecuteCommand(ctx, iCount, sUrl)
	}
	log.Print("Download complete.\n")
	log.Printf("下载完成，共 %d 个任务，请到 %s 目录下查看。\n", iCount, config.Conf.SaveFolder)
}

func ExecuteCommand(ctx context.Context, i int, sUrl string) {
	if sUrl == "" || !strings.HasPrefix(sUrl, "http") {
		fmt.Println("URL Not found.")
		return
	}
	sUrl = strings.Trim(sUrl, "\r\n")
	u, err := url.Parse(sUrl)
	if err != nil {
		fmt.Printf("URL Error:%+v\n", err)
		return
	}
	msg, err := router.FactoryRouter(u.Host, []string{sUrl})
	if err != nil {
		fmt.Println(err)
		return
	}
	if msg != nil {
		fmt.Printf("%+v\n", msg)
	}
	return
}
