package tests

import (
	"fmt"
	"github.com/NICEXAI/WeChatCustomerServiceSDK"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/cache"
	"testing"
)

func main(m *testing.M) {
	fmt.Println("test main")
}

func TestClient(t *testing.T) {
	client, err := WeChatCustomerServiceSDK.New(WeChatCustomerServiceSDK.Options{
		CorpID:         "ww559e6956be3ad889",
		Secret:         "uIAPlJlk1dgD5g8ciwcTDO9yMf5AawIlNbKaXIkbDXI",
		Token:          "qV4ZYTr2185",
		EncodingAESKey: "rYvgG3HAWCGmNZ41fcv7ydUCt2pHT1SPLgf5JfJ3bkR",
		Cache: cache.NewRedis(cache.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		ExpireTime:      7200,
		IsCloseCache:    false,
		IsCustomizedApp: false,
		SuiteTicket:     "4P8Lga4vk1N3O_kRiPYE1bjuxclwVgYxMAJJtiX9NjI7auwGmAQ6xtIedhu78mOX",
		SuiteId:         "dk7c21f60738e98afa",
		SuiteSecret:     "j0MFZvbjs65lrns7vcq2i04VqHUlibW3DrdbUsM_Wmk",
		PermanentCode:   "Uxa1FUWhWhQsnPr4G5rJPZDPtvEwXpKCT0x2qmxp73c",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	list, err := client.AccountList()
	if err != nil {
		return
	}
	fmt.Println(list)
}
