package main

import (
	_ "job_spider/init"
	"job_spider/spider"
)

func main(){
	spider := spider.NewWYSpider("php","sss")
	spider.Claw()
}
