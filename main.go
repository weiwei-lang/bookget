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
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx := context.Background()

	//配置初始化
	if !config.Init(ctx) {
		os.Exit(0)
	}

	//終端運行：单个URL
	if config.Conf.DUrl != "" {
		ExecuteCommand(ctx, 1, config.Conf.DUrl)
		log.Println("Download complete.")
		return
	}
	//終端運行：批量URLs
	if config.Conf.UrlsFile != "" {
		taskForUrls()
		return
	}
	//雙擊運行
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
	log.Println("Download complete.")
}

func taskForUrls() {
	//加载配置文件
	bs, err := os.ReadFile(config.Conf.UrlsFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	mUrls := strings.Split(string(bs), "\n")
	sortUrls := make(map[string][]string)
	for _, sUrl := range mUrls {
		sUrl = strings.Trim(sUrl, "\r\n")
		if sUrl == "" || !strings.HasPrefix(sUrl, "http") {
			continue
		}
		u, err := url.Parse(sUrl)
		if err != nil {
			continue
		}
		sortUrls[u.Host] = append(sortUrls[u.Host], sUrl)
	}
	for domain, sUrls := range sortUrls {
		wg.Add(1)
		go func(id string, data []string) {
			msg, err := router.FactoryRouter(id, data)
			if err != nil {
				fmt.Println(err)
				return
			}
			if msg != nil {
				fmt.Printf("%+v\n", msg)
			}
		}(domain, sUrls)
	}
	wg.Wait()
	log.Println("Download complete.")
	return
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
