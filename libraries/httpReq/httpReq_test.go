package httpReq

import "testing"

func TestWechatReq_SendWechat(t *testing.T) {
	wr := WechatReq{}
	wr.ToUser = "yubingtian"
	res, err := wr.SendWechat("单元测试")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(res)
}