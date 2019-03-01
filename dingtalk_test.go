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

	cases:= []*Request{
		&Request{
			MsgType:     MsgTypeText,
			AccessToken: "3946b02aae68e6a03138f60740645bb691bc802f1fa307acba9d2eddc1b516ae",
			Text: &TextMessage{
				Content: "test ding talk custom webhook robot",
				At: At{
					IsAtAll: true,
				},
			},
		},
		&Request{
			MsgType:     MsgTypeMarkdown,
			AccessToken: "3946b02aae68e6a03138f60740645bb691bc802f1fa307acba9d2eddc1b516ae",
			Markdown: &Markdown{
			    //Title: `test alert`,
			    Text: "#测试告警邮件\n在**2019-03-01T20:18:52+08:00**到**2019-03-01T20:18:52+08:00**期间\n发生了告警",
				At: At{
					IsAtAll: true,
				},
			},
		},
	}

	for _, tc := range cases {
		if err := td.SendMessage(context.Background(), tc); err != nil {
			t.Fatal(err)
		}
	}
}
