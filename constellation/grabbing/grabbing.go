package grabbing

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// ConstellationInfo 星座信息
type ConstellationInfo struct {
	LoveFortune    int64 // 爱情运势
	OverAllFortune int64 // 整体运势
	WealthFortune  int64 // 健康运势
}

const (
	baiyang   string = "https://www.meiguoshenpo.com/yunshi/baiyang/"   // 白羊座
	jinniu    string = "https://www.meiguoshenpo.com/yunshi/jinniu/"    // 金牛座
	shuangzi  string = "https://www.meiguoshenpo.com/yunshi/shuangzi/"  // 双子座
	juxie     string = "https://www.meiguoshenpo.com/yunshi/juxie/"     //  巨蟹座
	shizi     string = "https://www.meiguoshenpo.com/yunshi/shizi/"     // 狮子
	chunv     string = "https://www.meiguoshenpo.com/yunshi/chunv/"     // 处女
	tiancheng string = "https://www.meiguoshenpo.com/yunshi/tiancheng/" // 天秤座
	tianxie   string = "https://www.meiguoshenpo.com/yunshi/tianxie/"   // 天蝎座
	sheshou   string = "https://www.meiguoshenpo.com/yunshi/sheshou/"   // 射手
	mojie     string = "https://www.meiguoshenpo.com/yunshi/mojie/"     // 摩羯
	shuiping  string = "https://www.meiguoshenpo.com/yunshi/shuiping/"  // 水瓶
	shuangyu  string = "https://www.meiguoshenpo.com/yunshi/shuangyu/"  // 双鱼
)

// GetBaiyang 用于获取白羊座的星座信息
func GetBaiyang() *ConstellationInfo {
	return getData(baiyang)
}

// GetJinniu 用于获取金牛座的星座信息
func GetJinniu() *ConstellationInfo {
	return getData(jinniu)
}

// GetShuangzi 用于获取白羊座的星座信息
func GetShuangzi() *ConstellationInfo {
	return getData(shuangyu)
}

// GetJuxie  用于获取巨蟹座的信息
func GetJuxie() *ConstellationInfo {
	return getData(juxie)
}

// GetShizi 用于获取狮子座的信息
func GetShizi() *ConstellationInfo {
	return getData(shizi)
}

// GetChunv 用于获取处女座的信息
func GetChunv() *ConstellationInfo {
	return getData(chunv)
}

// GetTiancheng 用于获取天秤座的信息
func GetTiancheng() *ConstellationInfo {
	return getData(tiancheng)
}

// GetTianxie 用于获取天蝎座的信息
func GetTianxie() *ConstellationInfo {
	return getData(tianxie)
}

// GetSheshou 用于获取射手座的信息
func GetSheshou() *ConstellationInfo {
	return getData(sheshou)
}

// GetMojie 用于获取摩羯座的信息
func GetMojie() *ConstellationInfo {
	return getData(mojie)
}

// GetShuiping 用于获取水瓶座的信息
func GetShuiping() *ConstellationInfo {
	return getData(shuiping)
}

// GetShuangyu 用于获取双鱼座的星座信息
func GetShuangyu() *ConstellationInfo {
	return getData(shuangyu)
}

func getData(url string) *ConstellationInfo {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	pInstance := &ConstellationInfo{}
	doc.Find(".info ul li").Each(func(i int, s *goquery.Selection) {
		Description := s.Text()
		Class, _ := s.Find("em").Attr("class")
		if Description == "整体运势：" {
			pInstance.OverAllFortune, _ = strconv.ParseInt(Class[len(Class)-1:], 10, 64)
		} else if Description == "爱情运势：" {
			pInstance.LoveFortune, _ = strconv.ParseInt(Class[len(Class)-1:], 10, 64)
		} else if Description == "健康运势：" {
			pInstance.WealthFortune, _ = strconv.ParseInt(Class[len(Class)-1:], 10, 64)
		}
	})
	return pInstance
}
