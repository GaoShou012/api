package utils

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NatsClient(dns string, options ...interface{}) (*nats.Conn, error) {
	u, err := url.Parse(dns)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	username := u.User.Username()
	password, _ := u.User.Password()

	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	//var poolMax int
	//{
	//	arr, ok := params["PoolMax"]
	//	if !ok {
	//		poolMax = runtime.NumCPU()
	//	} else {
	//		num, err := strconv.Atoi(arr[0])
	//		if err != nil {
	//			return nil, err
	//		}
	//		poolMax = num
	//	}
	//}

	//var poolMin int
	//{
	//	arr, ok := params["PoolMin"]
	//	if !ok {
	//		poolMin = runtime.NumCPU()
	//	} else {
	//		num, err := strconv.Atoi(arr[0])
	//		if err != nil {
	//			return nil, err
	//		}
	//		poolMin = num
	//	}
	//}
	fmt.Println("i am addr:",addr)
	var timeout int
	{
		arr, ok := params["Timeout"]
		if !ok {
			timeout = 12
		} else {
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, err
			}
			timeout = num
		}
	}

	addrs := strings.Split(addr,",")
	fmt.Println("addrs:",addrs)
	// new redis_sortdset client
	_options := &nats.Options{
		Url:                         "",
		Servers:                     addrs,
		NoRandomize:                 false,
		NoEcho:                      false,
		Name:                        "",
		Verbose:                     false,
		Pedantic:                    false,
		Secure:                      false,
		TLSConfig:                   nil,
		AllowReconnect:              false,
		MaxReconnect:                0,
		ReconnectWait:               100 * time.Millisecond,
		CustomReconnectDelayCB:      nil,
		ReconnectJitter:             0,
		ReconnectJitterTLS:          0,
		Timeout:                     time.Duration(timeout) * time.Second,
		DrainTimeout:                0,
		FlusherTimeout:              0,
		PingInterval:                0,
		MaxPingsOut:                 0,
		ClosedCB:                    nil,
		DisconnectedCB:              nil,
		DisconnectedErrCB:           nil,
		ReconnectedCB:               nil,
		DiscoveredServersCB:         nil,
		AsyncErrorCB:                nil,
		ReconnectBufSize:            0,
		SubChanLen:                  0,
		UserJWT:                     nil,
		Nkey:                        "",
		SignatureCB:                 nil,
		User:                        username,
		Password:                    password,
		Token:                       "",
		TokenHandler:                nil,
		Dialer:                      nil,
		CustomDialer:                nil,
		UseOldRequestStyle:          false,
		NoCallbacksAfterClientClose: false,
	}
	natsClient, err := _options.Connect()

	//_dns := fmt.Sprintf("nats://%s", addr)
	//natsClient, err := nats.Connect(_dns, nil)
	if err != nil {
		return nil, err
	}
	return natsClient, nil
}
