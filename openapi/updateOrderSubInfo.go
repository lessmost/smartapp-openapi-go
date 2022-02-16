package openapi

import (
	"encoding/json"
	"fmt"
)

// UpdateOrderSubInfoRequest 请求结构体
type UpdateOrderSubInfoRequest struct {
	AccessToken string                              // 接口调用凭证
	OpenID      string                              // 用户 openId
	SceneID     string                              // 百度收银台分配的平台订单 ID，通知支付状态接口返回的 orderId
	SceneType   int64                               // 支付场景类型，开发者请默认传 2
	PmAppKey    string                              // 调起百度收银台的支付服务 appKey
	Data        []UpdateOrderSubInfoRequestDataItem // 请求数据
}

type UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail struct {
	Status     string `json:"Status"`     // 默认传 2
	SwanSchema string `json:"SwanSchema"` // 售后订单跳转地址，用以小程序跳转 Scheme
}
type UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem struct {
	Amount   int64  `json:"Amount"`   // 退款金额（单位：分），即100为1元
	ID       string `json:"ID"`       // 商品 ID
	Quantity int64  `json:"Quantity"` // 售后商品数量
}
type UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefund struct {
	Amount  string                                                                    `json:"Amount"`  // 退款总金额（单位：分），即100为1元。
	Product []UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem `json:"Product"` // 售后商品列表
}
type UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItem struct {
	CTime       int64                                                             `json:"CTime"`       // 创建时间（单位：秒）
	MTime       int64                                                             `json:"MTime"`       // 修改时间（单位：秒）
	OrderDetail UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail `json:"OrderDetail"` // 跳转到这个订单的详情结构
	OrderType   int64                                                             `json:"OrderType"`   // 退款类型，1(仅退款)，2(换货)，3(退款+退货)。
	Refund      UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefund      `json:"Refund"`      // 售后订单商品信息，详见 Data.Ext.SubsOrder.Item.Refund
	SubOrderID  string                                                            `json:"SubOrderID"`  // 售后订单 ID
	SubStatus   int64                                                             `json:"SubStatus"`   // 售后订单状态，同 Data.Ext.SubsOrder.Status 退换货枚举值一致
}
type UpdateOrderSubInfoRequestDataItemEXTSubsOrder struct {
	Items  []UpdateOrderSubInfoRequestDataItemEXTSubsOrderItemsItem `json:"Items"`  // 售后订单列表
	Status int64                                                    `json:"Status"` // 所有售后订单的状态汇总最终状态，详见 Data.Ext.SubsOrder.Status 退换货枚举值
}
type UpdateOrderSubInfoRequestDataItemEXT struct {
	SubsOrder UpdateOrderSubInfoRequestDataItemEXTSubsOrder `json:"SubsOrder"` // 子订单信息（退款、售后订单）
}
type UpdateOrderSubInfoRequestDataItem struct {
	BizAPPID   string                               `json:"BizAPPID"`   // 小程序 AppKey
	CateID     int64                                `json:"CateID"`     // 订单种类：1（实物）、2（虚拟物品）、5（快递服务类）、6（快递服务类无金额订单）、10（上门服务类）、11（上门服务类无金额订单）、15（酒店类）、20（票务类）、25（打车类）、26（打车类无金额订单）
	EXT        UpdateOrderSubInfoRequestDataItemEXT `json:"EXT"`        // 扩展信息
	ResourceID string                               `json:"ResourceID"` // 开发者接入的唯一订单 ID
}

// 响应结构体

type UpdateOrderSubInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST 请求参数中 BizAPPID
	CateID       string `json:"cate_id"`       // POST 请求参数中 CateID
	ResourceID   string `json:"resource_id"`   // POST 请求参数中 ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数（即请求是否成功， 0 为失败，非 0 为成功）
}

type UpdateOrderSubInfoResponse struct {
	Data      []UpdateOrderSubInfoResponsedataItem `json:"data"`       // 响应对象
	Errno     int64                                `json:"errno"`      // 错误码
	ErrMsg    string                               `json:"msg"`        // 错误信息
	ErrorCode int64                                `json:"error_code"` // openapi 错误码
	ErrorMsg  string                               `json:"error_msg"`  // openapi 错误信息
}

// UpdateOrderSubInfo
func UpdateOrderSubInfo(params *UpdateOrderSubInfoRequest) ([]UpdateOrderSubInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateOrderSubInfoResponsedataItem
	)
	respData := &UpdateOrderSubInfoResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/app/update/sub/info")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("open_id", params.OpenID)
	client.AddGetParam("scene_id", params.SceneID)
	client.AddGetParam("scene_type", fmt.Sprintf("%v", params.SceneType))
	client.AddGetParam("pm_app_key", params.PmAppKey)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"Data": params.Data,
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
