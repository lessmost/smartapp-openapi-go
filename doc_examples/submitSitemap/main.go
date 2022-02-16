// 本示例基于百度智能小程序服务端开发者 OpenAPI-SDK-Go
// 使用该示例需要首先下载该 SDK，使用引导见：https://smartprogram.baidu.com/docs/develop/serverapi/introduction_for_openapi_sdk/
// 使用之前请先确认下 SDK 版本是否为最新版本，如不是，请下载最新版本使用
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	// openapisdk-go.demo/openapi 为示例项目名，实际使用时需要替换为实际的项目名
	"openapisdk-go.demo/openapi"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &openapi.SubmitSitemapRequest{
		AccessToken: "21.8fa343ebfa8c131dda4b5a061c60377e.5802878.7104054755.075265-46421478",                                   // 文档中对应字段：access_token，实际使用时请替换成真实参数
		AppID:       13166334,                                                                                                   // 文档中对应字段：app_id，实际使用时请替换成真实参数
		Desc:        "智能小程序示例",                                                                                                  // 文档中对应字段：desc，实际使用时请替换成真实参数
		Frequency:   "3",                                                                                                        // 文档中对应字段：frequency，实际使用时请替换成真实参数
		Type:        "1",                                                                                                        // 文档中对应字段：type，实际使用时请替换成真实参数
		URL:         "https://mbs1.bdstatic.com/searchbox/mappconsole/image/88682533/635c15eb-e803-0b47-02ed-3f251a7e8d51.json", // 文档中对应字段：url，实际使用时请替换成真实参数
	}

	resp, err := openapi.SubmitSitemap(reqParams)
	if err != nil {
		if _, ok := err.(*openapi.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*openapi.APIError); ok {
			// api error
			// 可能是参数错误或没有权限，建议根据错误信息自查或者社区发帖、加群反馈
			fmt.Println("api error, ", err)
		} else {
			// 其他错误
			fmt.Println("others error, ", err)
		}
	} else {
		fmt.Printf("%#v\n", resp)
	}
}
