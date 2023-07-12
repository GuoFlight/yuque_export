package main

import (
	"fmt"
	"log"
	"os"
	"yuque_export/common"
	"yuque_export/conf"
	"yuque_export/flag"
)

func Init() {
	// 解析命令行
	flag.InitFlag()

	// 解析配置文件
	conf.ParseConfig(*flag.PathConfFile)

	// 功能导航
	if *flag.Version {
		fmt.Println(conf.Version)
		return
	}

	// 创建导出目录
	err := os.MkdirAll(conf.GConf.FileDir, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 初始化
	Init()

	// 得到所有bookID
	bookIDs, err := common.GetAllBookIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(bookIDs)
	// return

	// 得到所有docID
	for _, bookID := range bookIDs {
		docIDs, err := common.GetAllDocIDs(bookID)
		if err != nil {
			fmt.Printf("获取bookID=%d的文档id失败：%s\n", bookID, err.Error())
			continue
		}
		// fmt.Println(docIDs)
		// return

		// 得到下载链接
		for _, docID := range docIDs {
			downloadUrl, err := common.API{}.GetDownloadLink(docID)
			if err != nil {
				fmt.Printf("获取docID=%d的下载链接失败：%s\n", docID, err.Error())
				continue
			}
			// fmt.Println(downloadUrl)
			// return
			// 下载文件
			err = common.API{}.DownloadDoc(downloadUrl)
			if err != nil {
				fmt.Printf("文件下载失败，错误：%s，url：%s\n", err, downloadUrl)
				continue
			}
		}
	}
}
