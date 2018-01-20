/*

爬取豆瓣Top250的电影
并存在数据库中

抓取url:
https://movie.douban.com/top250
这是个分页的
第一页：
https://movie.douban.com/top250?start=0&filter=
第二页：
https://movie.douban.com/top250?start=25&filter=
...


抓取字段：电影名称、评分、评价人数

参考自：
http://blog.csdn.net/u013421629/article/details/72722302
原文是将结果输出到Excel文件中，
本程序是将结果输出到数据库，
其他未做改动，
因为刚接触go
仅仅为了熟悉go语法


//建表如下
CREATE TABLE `douban_movie_top250` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `score` varchar(255) DEFAULT NULL COMMENT '评分',
  `comment_count` int(11) DEFAULT NULL COMMENT '评价人数',
  `create_time` datetime DEFAULT NULL COMMENT '爬取时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}

//定义电影
type Movie struct {
	name         string
	score        float64
	commentCount int64
	rank int64
}

//定义 Spider get的方法
func (keyword Spider) get() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)

}

//获取网页数据
func getHtml(start string) string {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	url := "https://movie.douban.com/top250?start=" + "&filter="
	spider := &Spider{url, header}
	return spider.get()
}

//解析网页内容
func parseHtml(html string) []Movie {
	//定义一个切片，初始大小为25
	res := make([]Movie, 25)

	//排名
	//pattern1 := `<em class>.*</em>`
	//rp1 := regexp.MustCompile(pattern1)
	//find_txt1 := rp1.FindAllStringSubmatch(html, -1)

	//评价人数
	pattern2 := `<span>(.*?)人评价</span>`
	rp2 := regexp.MustCompile(pattern2)
	find_txt2 := rp2.FindAllStringSubmatch(html, -1)

	//评分
	pattern3 := `property="v:average">(.*?)</span>`
	rp3 := regexp.MustCompile(pattern3)
	find_txt3 := rp3.FindAllStringSubmatch(html, -1)

	//电影名称
	pattern4 := `alt="(.*?)" src="`
	rp4 := regexp.MustCompile(pattern4)
	find_txt4 := rp4.FindAllStringSubmatch(html, -1)

	for i := 0; i < len(find_txt2); i++ {
		//fmt.Printf("%s %s %s %s\n", find_txt1[i][1],find_txt4[i][1], find_txt3[i][1], find_txt2[i][1])
		fmt.Printf("%s %s %s\n", find_txt4[i][1], find_txt3[i][1], find_txt2[i][1])

		score, _ := strconv.ParseFloat(find_txt3[i][1], 64)
		commentCount, _ := strconv.ParseInt(find_txt2[i][1], 10, 32)
		//rank, _ := strconv.ParseInt(find_txt1[i][1], 10, 32)

		var movie Movie
		//movie.rank = rank
		movie.name = find_txt4[i][1]
		movie.score = score
		movie.commentCount = commentCount
		res[i] = movie
	}

	return res
}

//插入数据库
func insertDb(movies []Movie) {
	//获取数据库连接
	db, e := sql.Open("mysql", "root:@/spider_practice")
	if e != nil {
		panic(e)
	}

	//Begin函数内部会去获取连接
	tx, _ := db.Begin()

	for i := 0; i < len(movies); i++ {
		movie := movies[i]
		_, e := tx.Exec("INSERT INTO douban_movie_top250(rank, name,score,comment_count,create_time) values(?,?,?,?,?)", movie.rank, movie.name, movie.score, movie.commentCount, time.Now())
		if e != nil {
			//如果插入失败，则结束整个程序
			panic(e)
		}
	}

	//最后释放tx内部的连接
	tx.Commit()
}

//
func do()  {
	//top250，每页是25个，一共10页
	for i := 0; i < 10; i++{
		html := getHtml(strconv.Itoa(i*25))
		movies := parseHtml(html)
		insertDb(movies)
	}
}



func main() {
	t1 := time.Now() // get current time
	do()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)

}

