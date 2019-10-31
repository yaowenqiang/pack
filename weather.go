package pack
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
)

func PrintWeather(zipCode int) {
/*
	resp, err := http.Get("http://wsf.cdyne.com//WeatherWS/" + 
	"Weather.asmx/GetCityWeatherByZIP?ZIP=" + strconv.Itoa(zipCode))
*/

//http://www.weather.com.cn/data/cityinfo/101010100.html

	resp, err := http.Get("http://www.weather.com.cn/data/cityinfo/" + strconv.Itoa(zipCode) + ".html")


	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("%#v", resp.Body)
	//data := make([]byte, resp.ContentLength)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		fmt.Println(err)
		return
	}


	var weather  Weather
	err = json.Unmarshal(body , &weather)

	if err != nil {
		fmt.Println(err)
		return
	}
    
	fmt.Println(weather)

}

type WeatherInfo  struct {
	Weatherinfo Weather
}

/*
type Weather  struct {
	City string
	Cityid string
	Temp1  float32
	Temp2 float32
	Weather string
	Img1 string
	Img2 string
	Ptime string
}

*/

//https://mholt.github.io/json-to-go/

type Weather struct {
	Weatherinfo struct {
		City string `json:"city"`
		Cityid string `json:"cityid"`
		Temp1 string `json:"temp1"`
		Temp2 string `json:"temp2"`
		Weather string `json:"weather"`
		Img1 string `json:"img1"`
		Img2 string `json:"img2"`
		Ptime string `json:"ptime"`
	} `json:"weatherinfo"`
}
