package dingtalk

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	address = "https://oapi.dingtalk.com/robot/send?access_token="

	MsgTypeText                  = "text"
	MsgTypeLink                  = "link"
	MsgTypeMarkdown              = "markdown"
	MsgTypeIndependentActionCard = "independent_action_card"
	MsgTypeIntegratedActionCard  = "integrated_action_card"
	MsgTypeFeedCard              = "feed_card"

	keyMsgType    = "msgtype"
	keyAt         = "at"
	keyText       = MsgTypeText
	keyLink       = MsgTypeLink
	keyMarkdown   = MsgTypeMarkdown
	keyActionCard = "actionCard"
	keyFeedCard   = "feedcard"
)

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// TextMessage 文本类型消息
type TextMessage struct {
	Content string `json:"content"`
	At      At     `json:"at"`
}

// LinkMessage 链接类型消息
type LinkMessage struct {
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl"`
	Text       string `json:"text"`
	Title      string `json:"title"`
}

// Markdown 类型消息
type Markdown struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	At    At     `json:"at"`
}

// IntegratedActionCard 整体 ActionCard 类型
type IntegratedActionCard struct {
	BtnOrientation string `json:"btnOrientation"`
	HideAvatar     string `json:"hideAvatar"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Text           string `json:"text"`
	Title          string `json:"title"`
}

// IndependentActionCard 独立跳转 ActionCard 类型
type IndependentActionCard struct {
	BtnOrientation string                     `json:"btnOrientation"`
	Btns           []IndependentActionCardBtn `json:"btns"`
	HideAvatar     string                     `json:"hideAvatar"`
	Text           string                     `json:"text"`
	Title          string                     `json:"title"`
}

type IndependentActionCardBtn struct {
	ActionURL string `json:"actionURL"`
	Title     string `json:"title"`
}

// FeedCard 类型
type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

type FeedCardLink struct {
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
	Title      string `json:"title"`
}

type Request struct {
	MsgType     string
	AccessToken string

	At                    *At
	Text                  *TextMessage
	Link                  *LinkMessage
	Markdown              *Markdown
	IndependentActionCard *IndependentActionCard
	IntegratedActionCard  *IntegratedActionCard
	FeedCard              *FeedCard
}

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (r *Request) GetBody() ([]byte, error) {
	bodyDict := make(map[string]interface{})
	bodyDict[keyMsgType] = r.MsgType

	switch r.MsgType {
	case MsgTypeText:
		bodyDict[keyText] = r.Text
		bodyDict[keyAt] = r.At
	case MsgTypeLink:
		bodyDict[keyLink] = r.Link
	case MsgTypeMarkdown:
		bodyDict[keyMarkdown] = r.Markdown
		bodyDict[keyAt] = r.At
	case MsgTypeIndependentActionCard:
		bodyDict[keyActionCard] = r.IndependentActionCard
	case MsgTypeIntegratedActionCard:
		bodyDict[keyActionCard] = r.IntegratedActionCard
	case MsgTypeFeedCard:
		bodyDict[keyFeedCard] = r.FeedCard
	default:
		return nil, errors.New("invalid DingTalk type, " + r.MsgType)
	}
	bodyInBytes, err := json.Marshal(bodyDict)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal body")
	}
	return bodyInBytes, nil
}

func (r *Request) GetAccessToken() string {
	return r.AccessToken
}
