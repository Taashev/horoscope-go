package horoscope

import (
	"encoding/xml"
	"fmt"
	horoscopeUtils "horoscope_console/internal/horoscope"
	"io/ioutil"
	"net/http"
)

func GetHoroscope(zodiacSign string) (string, error) {
	// API для получения гороскопа
	apiURL := "http://img.ignio.com/r/export/utf/xml/daily/com.xml"

	// Выполняем GET-запрос
	response, err := http.Get(apiURL)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	// Считываем ответ
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	var horoscope Horo

	err = xml.Unmarshal(body, &horoscope)

	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}

	// Проверяем, есть ли знак в словаре
	englishName, err := horoscopeUtils.ZodiacEnglishMapping(zodiacSign)

	if err != nil {
		return "", err
	}

	sign := map[string]HoroscopeSign{
		"aries":       horoscope.Aries,
		"taurus":      horoscope.Taurus,
		"gemini":      horoscope.Gemini,
		"cancer":      horoscope.Cancer,
		"leo":         horoscope.Leo,
		"virgo":       horoscope.Virgo,
		"libra":       horoscope.Libra,
		"scorpio":     horoscope.Scorpio,
		"sagittarius": horoscope.Sagittarius,
		"capricorn":   horoscope.Capricorn,
		"aquarius":    horoscope.Aquarius,
		"pisces":      horoscope.Pisces,
	}[englishName]

	return sign.Today, nil
}
