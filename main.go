package main

//接受一个图片url列表的文本文件, 将图片下载至指定的文件夹下,
//运行过程中需要打印出每张图片的url及其下载时间, 如果失败则输出错误

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//创建文本文件
/*func creTxt(fileName string){

	f,err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte ("http://via.placeholder.com/350/E8117F/ffffff?text=123"))
		checkErr(err)
	}

}
*/

//创建文件夹
func creFile(dirName string) {
	// 创建文件夹
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		log.Printf("mkdir failed![%v]\n", err)
	} else {
		log.Printf("mkdir success!\n")
	}
}

//读取文件
func getImg(Filename string, dirName string) {
	data, err := os.Open(Filename)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	f := bufio.NewReader(data)
	i := 1
	for {

		line, err := f.ReadString(byte('\n')) //按行读取文件

		if err != nil || io.EOF == err {
			break
		} else {
			start := time.Now()
			response, err := http.Get(strings.TrimSpace(line)) //获取图片,line去掉空行
			if err != nil {
				log.Printf("batch-downloader -i image-list.txt -o folder -c 10 , error:%s\n", err.Error())

			} else {
				log.Println(line) //输出每一行的url
				defer response.Body.Close()
				s := fmt.Sprintf("%s/%d.jpg", dirName, i)
				downloadFile, err := os.Create(s) //创建文件到指定目录
				if err != nil {
					panic(err)
				} else {
					io.Copy(downloadFile, response.Body) //copybody到文件
					end := time.Now()
					log.Println("Cost time: %v", end.Sub(start))
				}

			}

			i++
		}
	}
}

func main() {

	sta := `
    Start download photos!
Please wait for a few minutes.
	`
	fmt.Println(sta)   //打印开始提示信息
	_dir := "./Files1" //folder名
	// creFile(_dir)
	file := "tempfile1.txt" //源文件名
	go getImg(file, _dir)

	// downFile := make([]os.File, 300);
	//downFile := getImg(file, _dir)
}
