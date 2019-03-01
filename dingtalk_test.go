package dingtalk

import (
	"context"
	"testing"
)

func TestDingTalk_GetBody(t *testing.T) {
	r := &Request{
		MsgType: MsgTypeIndependentActionCard,
		IndependentActionCard: &IndependentActionCard{
			BtnOrientation: "xxx",
			Btns: []IndependentActionCardBtn{
				{
					ActionURL: "www1",
					Title:     "title in btns 1",
				},
				{
					ActionURL: "www2",
					Title:     "title in btns 2",
				},
			},
			HideAvatar: "false",
			Text:       "xxx",
			Title:      "outer title",
		},
	}

	body, _ := r.GetBody()
	t.Logf("%s", string(body))
}

func TestDingTalk_SendMessage(t *testing.T) {
	td := NewClient()

	if err := td.SendMessage(context.Background(), &Request{
		MsgType:     MsgTypeText,
		AccessToken: "3946b02aae68e6a03138f60740645bb691bc802f1fa307acba9d2eddc1b516ae",
		Text: &TextMessage{
			Content: "test ding talk custom webhook robot",
			At: At{
				IsAtAll: true,
			},
		},
	}); err != nil {
		t.Fatal(err)
	}
}
