package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"job_spider/common"
	"job_spider/init/mysql"
	"job_spider/model"
	"regexp"
)

type hbSpider struct {
	Class string
	Url string
	FromType string
}

func NewHBSpider(class string,url string)*hbSpider{
	return &hbSpider{
		Class: class,
		Url:   url,
		FromType:	"HB",
	}
}

func (spider *hbSpider)Claw()  {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.64"),
	)
	// Find and visit all links
	c.OnHTML("#job_list_table .postIntro", func(e *colly.HTMLElement) {
		company := e.ChildText(".title > a")
		e.ForEach(".postIntroLx", func(i int, ee *colly.HTMLElement) {
			name := ee.ChildText(".name")
			reg := regexp.MustCompile(`(php)|(PHP)`)
			regReult := reg.FindAllString(name,-1)
			if(len(regReult)<=0){
				return
			}
			address := e.ChildText(".address")
			salary := e.ChildText(".money")
			minSalary,maxSalary,error1 := common.GetHBMaxMinSalary(salary)
			var timeStr string
			ee.ForEach("p > span.job_time", func(j int, eee *colly.HTMLElement) {
				timeStr = eee.Text
			})
			if error1 != nil {
				println(error1.Error())
			}
			fmt.Println(fmt.Sprintf("jobName:%s,companyName:%s,address:%s,minSalary:%d,maxSalary:%d,time:%s",
				name,
				company,
				address,
				minSalary,
				maxSalary,
				common.GetHBPublishAT(timeStr),
			),
			)
			publishAt := common.GetHBPublishAT(timeStr)
			job := &model.Job{
				Name:      name,
				Company:   company,
				Address:   address,
				MinSalary: int64(minSalary),
				MaxSalary: int64(maxSalary),
				Class:     "php",
				Status:    0,
				PublishAt: publishAt,
				FromType: 		"HB",
			}
			findJob := new(model.Job)
			err := mysql.DB.Where("name = ? and company = ? and min_salary = ? and max_salary = ? and from_type = ?",name,company,minSalary,maxSalary,spider.FromType).First(&findJob).Error
			if(err != gorm.ErrRecordNotFound){
				mysql.DB.Model(&findJob).Update("status",1)
			}else {
				mysql.DB.Save(job)
			}
		})

	})
	c.OnHTML("#job_list_table > div.postSchList > div.zright > p.milpage > a.arr_r", func(e *colly.HTMLElement){
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
	c.Visit("https://www.huibo.com/jobsearch/?key=php")
}