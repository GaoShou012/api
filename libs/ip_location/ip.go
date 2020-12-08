// ip 信息
package ip_location

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
)

type location struct {
	Country  string // 国家
	Province string // 省份
	City     string // 城市
}

var db *ipdb.City

func Init(filePath string) error {
	_db, err := ipdb.NewCity(filePath)
	if err != nil {
		return fmt.Errorf("init ip location db failed path=%s err=%v", filePath, err)
	}
	db = _db
	return nil
}

func GetLocation(ip string) (*location, error) {
	defaultLocation := &location{
		Country:  "中国",
		Province: "",
		City:     "",
	}

	loc, err := db.FindInfo(ip, "CN")
	if err != nil {
		return defaultLocation, err
	}

	defaultLocation.Country = loc.CountryName
	defaultLocation.Province = loc.RegionName
	defaultLocation.City = loc.CityName

	return defaultLocation, nil
}
