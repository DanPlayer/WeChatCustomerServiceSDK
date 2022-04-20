package WeChatCustomerServiceSDK

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
)

const (
	// 获取调用凭证AccessToken
	getTokenAddr = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	// 获取自建应用凭证
	getSuiteToken = "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"
	// 获取自建应用调用企业凭证AccessToken
	getCorpTokenAddr = "https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token?suite_access_token=%s"
)

type SuiteAccessTokenSchema struct {
	BaseModel
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

type GetSuiteAccessTokenOptions struct {
	SuiteId     string `json:"suite_id"`
	SuiteSecret string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

func (r *Client) getSuiteToken(options GetSuiteAccessTokenOptions) (info SuiteAccessTokenSchema, err error) {
	data, err := util.HttpPost(getSuiteToken, options)
	if err != nil {
		return
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type GetCorpTokenOptions struct {
	AuthCorpid    string `json:"auth_corpid"`
	PermanentCode string `json:"permanent_code"`
}

// AccessTokenSchema 获取调用凭证响应数据
type AccessTokenSchema struct {
	BaseModel
	AccessToken string `json:"access_token"` // 获取到的凭证，最长为512字节
	ExpiresIn   int    `json:"expires_in"`   // 凭证的有效时间（秒）
}

// GetAccessToken 获取调用凭证access_token
func (r *Client) GetAccessToken() (info AccessTokenSchema, err error) {
	var data []byte
	if r.isCustomizedApp {
		var suiteToken SuiteAccessTokenSchema
		suiteToken, err = r.getSuiteToken(GetSuiteAccessTokenOptions{
			SuiteId:     r.suiteId,
			SuiteSecret: r.suiteSecret,
			SuiteTicket: r.suiteTicket,
		})
		if err != nil {
			return
		}
		data, err = util.HttpPost(fmt.Sprintf(getCorpTokenAddr, suiteToken.SuiteAccessToken), GetCorpTokenOptions{
			AuthCorpid:    r.corpID,
			PermanentCode: r.permanentCode,
		})
	} else {
		data, err = util.HttpGet(fmt.Sprintf(getTokenAddr, r.corpID, r.secret))
	}

	if err != nil {
		return info, err
	}
	fmt.Println(string(data))
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// RefreshAccessToken 刷新调用凭证access_token
func (r *Client) RefreshAccessToken() error {
	//初始化AccessToken
	tokenInfo, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	if err = r.setAccessToken(tokenInfo.AccessToken); err != nil {
		return err
	}
	r.accessToken = tokenInfo.AccessToken
	return nil
}

func (r *Client) initAccessToken() error {
	//如果关闭自动缓存则直接刷新AccessToken
	if r.isCloseCache {
		if err := r.RefreshAccessToken(); err != nil {
			return err
		}
		return nil
	}

	//判断是否已初始化完成，如果己初始化则直接返回当前实例
	token, err := r.getAccessToken()
	if err != nil {
		return NewSDKErr(50002)
	}
	if token == "" {
		if err = r.RefreshAccessToken(); err != nil {
			return err
		}
	} else {
		r.accessToken = token
	}
	return nil
}

func (r *Client) getAccessToken() (string, error) {
	return r.cache.Get("wechat:kf:" + r.corpID)
}

func (r *Client) setAccessToken(token string) error {
	return r.cache.Set("wechat:kf:"+r.corpID, token, r.expireTime)
}
