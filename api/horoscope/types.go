package horoscope

import "encoding/xml"

type Horo struct {
	XMLName xml.Name `xml:"horo"`
	Date    struct {
		Yesterday  string `xml:"yesterday,attr"`
		Today      string `xml:"today,attr"`
		Tomorrow   string `xml:"tomorrow,attr"`
		Tomorrow02 string `xml:"tomorrow02,attr"`
	} `xml:"date"`
	Aries struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"aries"`
	Taurus struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"taurus"`
	Gemini struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"gemini"`
	Cancer struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"cancer"`
	Leo struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"leo"`
	Virgo struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"virgo"`
	Libra struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"libra"`
	Scorpio struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"scorpio"`
	Sagittarius struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"sagittarius"`
	Capricorn struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"capricorn"`
	Aquarius struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"aquarius"`
	Pisces struct {
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"pisces"`
}

type HoroscopeSign struct {
	Yesterday  string `xml:"yesterday"`
	Today      string `xml:"today"`
	Tomorrow   string `xml:"tomorrow"`
	Tomorrow02 string `xml:"tomorrow02"`
}
