package openapi

import (
	"encoding/json"
)

// UpdateCouponBannerRequest 请求结构体
type UpdateCouponBannerRequest struct {
	AccessToken     string // 接口调用凭证
	CouponID        string // 卡券 ID
	PicURL          string // 卡券 banner 图片
	Title           string // 卡券 banner 图标题
	AppRedirectPath string // banner 图跳转的小程序页面路径
	BannerID        int64  // 卡券 banner 记录 id
}

// 响应结构体

type UpdateCouponBannerResponse struct {
	Data      bool   `json:"data"`       // true 退还成功 false 退还失败
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// UpdateCouponBanner
func UpdateCouponBanner(params *UpdateCouponBannerRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &UpdateCouponBannerResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/banner/update")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"couponId":        params.CouponID,
		"picUrl":          params.PicURL,
		"title":           params.Title,
		"appRedirectPath": params.AppRedirectPath,
		"bannerId":        params.BannerID,
	}
	bts, err := json.Marshal(postData)
	if err != nil {
		return defaultRet, err
	}
	client.SetBody(bts)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return respData.Data, nil
}
