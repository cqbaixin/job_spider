package main

import (
	_ "job_spider/init"
	"job_spider/spider"
)

func main(){
	wySpider := spider.NewWYSpider("php","ssss")
	hbSpider := spider.NewHBSpider("php","ssss")
	hbSpider.Claw()
	wySpider.Claw()
}
