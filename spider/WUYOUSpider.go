package spider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/jinzhu/gorm"
	"job_spider/common"
	"job_spider/init/mysql"
	"job_spider/model"
	"regexp"
	"time"
)

type wySpider struct {
	Class string
	Url string
	FromType string
}

func NewWYSpider(class string,url string)*wySpider{
	return &wySpider{
		Class: class,
		Url:   url,
		FromType:"51",
	}
}

func (wy *wySpider)Claw()  {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.64"),
	)
	// Find and visit all links
	c.OnHTML(".dw_table .el", func(e *colly.HTMLElement) {
		name := e.ChildText(".t1")
		reg := regexp.MustCompile(`(php)|(PHP)`)
		regReult := reg.FindAllString(name,-1)
		if(len(regReult)<=0){
			return
		}
		company := e.ChildText(".t2")
		address := e.ChildText(".t3")
		salary := e.ChildText(".t4")
		minSalary,maxSalary,error1 := common.GetMaxMinSalary(salary)
		if error1 != nil {
			println(error1.Error())
		}
		timeStr := e.ChildText(".t5")
		fmt.Println(fmt.Sprintf("jobName:%s,companyName:%s,address:%s,minSalary:%d,maxSalary:%d,time:%s",
			name,
			company,
			address,
			minSalary,
			maxSalary,
			timeStr,
		),
		)
		publishAt,_ := time.Parse("2006-01-02",fmt.Sprintf("%d-%s",time.Now().Year(),timeStr))
		job := &model.Job{
			Name:      name,
			Company:   company,
			Address:   address,
			MinSalary: int64(minSalary),
			MaxSalary: int64(maxSalary),
			Class:     "php",
			Status:    0,
			PublishAt: publishAt,
			FromType: 		wy.FromType,
		}
		findJob := new(model.Job)
		err := mysql.DB.Where("name = ? and company = ? and min_salary = ? and max_salary = ? and from_type = ? ",name,company,minSalary,maxSalary,wy.FromType).First(&findJob).Error
		if(err != gorm.ErrRecordNotFound){
			mysql.DB.Model(&findJob).Update("status",1)
		}else {
			mysql.DB.Save(job)
		}
	})

	c.OnHTML("#rtNext", func(e *colly.HTMLElement){
		nextUrl := e.Attr("href")
		if(len(nextUrl)>0){
			c.Visit(nextUrl)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Body", string(r.Body[:]))
		body,err := common.StrToUtf8(fmt.Sprintf("%s",r.Body))
		if(err != nil){
			fmt.Println(err.Error())
			return
		}
		r.Body = []byte(body)
	})
	c.Visit("https://search.51job.com/list/060000,000000,0000,00,1,99,php,2,1.html")
}
