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
		Secret:         "phC446yva2mQ91hTBvQkaZLRLi_XrNUFZ0FKbXgJ6o0",
		Token:          "epwpgcrjiledhmfa",
		EncodingAESKey: "hsiberxbcbffhqohswqxehexouspqiwczuljjzbwkny",
		Cache: cache.NewRedis(cache.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		ExpireTime:      7200,
		IsCloseCache:    false,
		IsCustomizedApp: true,
		SuiteTicket:     "T38D7-gn9_wgguU3hz5kPtIanS5li0tF44cHvUc85TiRGbuvZtYMgHHEEonTLIcT",
		SuiteId:         "dk7c21f60738e98afa",
		SuiteSecret:     "j0MFZvbjs65lrns7vcq2i04VqHUlibW3DrdbUsM_Wmk",
		PermanentCode:   "uIAPlJlk1dgD5g8ciwcTDO9yMf5AawIlNbKaXIkbDXI",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)
}
