package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var visitedLink = make(map[string]struct{})

//Распарсит байты и вернет все ссылки (<a href="someLink"></a>)
func getLinks(data []byte) ([]string, error) {
	links := []string{}
	reg, err := regexp.Compile("href=\".+?\"")
	if err != nil {
		return nil, err
	}

	matched := reg.FindAll(data, -1)

	for _, link := range matched {
		links = append(links, string(link[6:len(link)-1]))
	}

	return links, nil
}

//Заменит ссылки на локальные
func mirror(data []byte) []byte {
	reg, _ := regexp.Compile("href=\".+?\"")

	data = reg.ReplaceAllFunc(data, func(m []byte) []byte{
		pUrl, _ := url.Parse(string(m[6:len(m)-1]))
		pUrl.Scheme = ""
		pUrl.Host = ""

		if len(pUrl.String()) > 0 {
			return []byte("href=\"" + pUrl.String()[1:] + "\"")
		}
		return []byte("href=\"./\"")
	})

	return data
}

//data links error
func request(url string, sFlag bool) ([]byte, []string, error) {
	links := []string{}

	//Отправляем запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	//Получаем данные
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if !sFlag{
		nameOfFile := strings.Split(url, "/")
		ioutil.WriteFile(nameOfFile[len(nameOfFile) - 1], data, 0700)
		return data, []string{}, nil
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") {
		//Ищем все ссылки на странице
		links, err = getLinks(data)
		if err != nil {
			return nil, nil, err
		}
	}

	makeFiles(url, contentType, data)

	return data, links, nil
}

func makeFiles(link string, contentType string, data []byte) error {
	pLink, err := url.Parse(link)
	if err != nil {
		//Если ссылка не парсится игнорим
		return nil
	}

	//Если контент html - создаем папку и в ней index.html
	if strings.HasPrefix(contentType, "text/html") {
		os.MkdirAll(pLink.Host + pLink.Path, 0700)
		//data = mirror(data)
		ioutil.WriteFile(pLink.Host + pLink.Path + "/index.html", data, 0700)
	} else {
		//Иначе создаем папку до последнего слеша и создаем там файл
		splitedPath := strings.Split(pLink.Path, "/")
		os.MkdirAll(pLink.Host + strings.Join(splitedPath[:len(splitedPath)-1], "/"), 0700)
		ioutil.WriteFile(pLink.Host + pLink.Path, data, 0700)
	}

	return nil
}

func wget(target string, sFlag bool) error {
	pTarget, err := url.Parse(target)
	if err != nil {
		return err
	}

	_, links, err := request(target, sFlag)
	if err != nil {
		return nil
	}

	for _, link := range links {
		pLink, err := url.Parse(link)
		if err != nil {
			return err
		}

		if pLink.Host == "" {
			pLink.Host = pTarget.Host
			pLink.Scheme = pTarget.Scheme
		}

		//Ссылка относительная или ведет на тот же ресурс или это css и без # якорей
		if (pTarget.Host == pLink.Host || pLink.Host == "") &&
			pLink.Scheme != "" &&
			!strings.Contains(pLink.String(), "#") ||
			strings.HasSuffix(pLink.String(), ".css") {

			//Рекурсивно скачиваем, непосещенные
			if _, ok := visitedLink[pLink.String()]; !ok {
				visitedLink[pLink.String()] = struct{}{}
				err := wget(pLink.String(), sFlag)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func Spinner() {
	for {
		for _, r := range `█▒▒▒▒▒▒▒▒▒ ███▒▒▒▒▒▒▒ █████▒▒▒▒▒ ███████▒▒▒ ██████████` {
			if r == ' ' {
				fmt.Print("\r")
			}
			fmt.Printf("%c", r)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func main() {
	//"https://ru.simplesite.com" //http://ftp.gnu.org/gnu/wget/wget-1.5.3.tar.gz
	target := flag.String("url", "", "urlForDownload")
	sflag := flag.Bool("s", false, "downloadSite")

	flag.Parse()

	go Spinner()
	err := wget(*target, *sflag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
