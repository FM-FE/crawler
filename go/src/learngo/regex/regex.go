package main

import (
	"fmt"
	"regexp"
)

//const text = `my email is fmfe@gmail.com.cn@com
//email1 = 123@qq.com
//email2 = kkk@def.com`

const str = `<a href="http://www.zhenai.com/zhenghun/aba" data-v-0c63b635>阿坝</a>`
const name  = `<a href="http://www.zhenai.com/zhenghun/[a-z][0-9]+" [^>]*>[^<]+</a>`
func main() {
	//re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)
	//fmt.Printf("email name is : \n")
	//for i := range match{
	//	fmt.Printf("%d : %s \n",i,match[i][1])
	//}
	//for _, m := range match {
	//	fmt.Println(m)
	//}

	contents := `<div class="m-btn purple" data-v-ff544c08>离异</div>
	<div class="m-btn purple" data-v-ff544c08>51岁</div>
	<div class="m-btn purple" data-v-ff544c08>魔羯座(12.22-01.19)</div>
	<div class="m-btn purple" data-v-ff544c08>160cm</div>
	<div class="m-btn purple" data-v-ff544c08>53kg</div>
	<div class="m-btn purple" data-v-ff544c08>工作地:阿坝茂县</div>
	<div class="m-btn purple" data-v-ff544c08>月收入:3千以下</div>
	<div class="m-btn purple" data-v-ff544c08>其他职业</div>
	<div class="m-btn purple" data-v-ff544c08>高中及以下</div>`

	a := []byte(contents)

	//re := regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>.+座.*</div>`)
	re := regexp.MustCompile(`<div class="m-btn purple" .*>(.+)座.*</div>`)

	// care about the [^>] and [^<], both of those define the limit of regexp

	s := re.FindAllSubmatch(a, -1)

	for _, m := range s {
		for _, n := range m {
			fmt.Printf("%s\n",string(n))
		}
	}
}
