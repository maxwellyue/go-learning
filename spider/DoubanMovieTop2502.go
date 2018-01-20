/*

使用goquery爬取豆瓣电影top250

与DoubanTopMovie250.go不同的只是不使用正则，

而是使用goquery

 */
package main
import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)



func parseUsingGoQuery(start string) []Movie{
	res := make([]Movie, 25)

	url := "https://movie.douban.com/top250?start=" + start + "&filter="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	doc.Find("div.item").Each(func(i int, selection *goquery.Selection) {
		originComment := selection.Find("div.star span").Last().Text()
		originComment = removeLast(originComment, 3)
		name := selection.Find("span.title").First().Text()
		score, _ := strconv.ParseFloat(selection.Find("span.rating_num").First().Text(), 64)
		commentCount, _ := strconv.ParseInt(originComment, 10, 32)
        rank, _ := strconv.ParseInt(selection.Find("div.pic em").First().Text(), 10, 32)

		fmt.Println(selection.Find("div.pic em").First().Text() + "  " + name + "  " + selection.Find("span.rating_num").First().Text() + "  " + originComment)

		var movie Movie
		movie.rank = rank
		movie.name = name
		movie.score = score
		movie.commentCount = commentCount
		res[i] = movie
	})

	return res
}

//移除字符串str的最后l个字符
func removeLast(str string, l int) string{
	rs := []rune(str)
	length := len(rs)

	if l < 0 || l > length {
		panic("l is wrong")
	}
    end := length - l
	return string(rs[0:end])
}

func doUsingGoQuery()  {
	//top250，每页是25个，一共10页
	for i := 0; i < 10; i++{
		movies := parseUsingGoQuery(strconv.Itoa(i*25))
		insertDb(movies)
	}
}


func main() {
	t1 := time.Now() // get current time
	doUsingGoQuery()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)

}

