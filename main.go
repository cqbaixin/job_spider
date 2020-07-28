package main

import (
	_ "job_spider/init"
	"job_spider/spider"
)

func main(){
<<<<<<< HEAD
	spider := spider.NewWYSpider("php","ssss")
=======
	spider := spider.NewWYSpider("php","sss")
>>>>>>> master
	spider.Claw()
}
