package global

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"net/url"
)

func InitCasbinAdapter(dns string) error {
	u, err := url.Parse(dns)
	if err != nil {
		return err
	}
	databaseType := u.Scheme
	username := u.User.Username()
	password, _ := u.User.Password()
	_dns := fmt.Sprintf("%s:%s@(%s)%s?%s", username, password, u.Host, u.Path, u.RawQuery)
	a, err := gormadapter.NewAdapter(databaseType, _dns, true)
	if err != nil {
		return err
	}
	CasbinAdapter = a
	return nil
}
