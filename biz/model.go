package biz

/*
	General Model
*/
type Header struct {
	EventID    string `json:"event_id"`
	EventType  string `json:"event_type"`
	CreateTime string `json:"create_time"`
	Token      string `json:"token"`
	AppID      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TenantAccessTokenRequest struct {
	APPID     string `json:"app_id"`
	APPSecret string `json:"app_secret"`
}

type TenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	Expire            int    `json:"expire"`
	TenantAccessToken string `json:"tenant_access_token"`
}

/*
	Chat API Model
*/

type CreateChatRequest struct {
	Avatar                 string     `json:"avatar,omitempty"`
	Name                   string     `json:"name,omitempty"`
	Description            string     `json:"description,omitempty"`
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`
	OwnerId                string     `json:"owner_id,omitempty"`
	ChatMode               string     `json:"chat_mode,omitempty"`
	ChatType               string     `json:"chat_type,omitempty"`
	External               bool       `json:"external,omitempty"`
	JoinMessageVisibility  string     `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string     `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string     `json:"membership_approval,omitempty"`
}

type CreateChatResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    *CreateChatRespBody `json:"data"`
}

type CreateChatRespBody struct {
	ChatId                 string     `json:"chat_id,omitempty"`
	Avatar                 string     `json:"avatar,omitempty"`
	Name                   string     `json:"name,omitempty"`
	Description            string     `json:"description,omitempty"`
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`
	OwnerId                string     `json:"owner_id,omitempty"`
	OwnerIdType            string     `json:"owner_id_type,omitempty"`
	AddMemberPermission    string     `json:"add_member_permission,omitempty"`
	ShareCardPermission    string     `json:"share_card_permission,omitempty"`
	AtAllPermission        string     `json:"at_all_permission,omitempty"`
	EditPermission         string     `json:"edit_permission,omitempty"`
	ChatMode               string     `json:"chat_mode,omitempty"`
	ChatType               string     `json:"chat_type,omitempty"`
	ChatTag                string     `json:"chat_tag,omitempty"`
	External               bool       `json:"external,omitempty"`
	TenantKey              string     `json:"tenant_key,omitempty"`
	JoinMessageVisibility  string     `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string     `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string     `json:"membership_approval,omitempty"`
	ModerationPermission   string     `json:"moderation_permission,omitempty"`
}

type I18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"`
	EnUs string `json:"en_us,omitempty"`
	JaJp string `json:"ja_jp,omitempty"`
}

type ChatMembersInviteRequest struct {
	IdList []string `json:"id_list,omitempty"`
}

type ChatMembersInviteResponse struct {
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Data    *ChatMembersInviteRespBody `json:"data"`
}

type ChatMembersInviteRespBody struct {
	InvalidIDList []string `json:"invalid_id_list"`
}

type UpdateChatRequest struct {
	Avatar                 string     `json:"avatar,omitempty"`
	Name                   string     `json:"name,omitempty"`
	Description            string     `json:"description,omitempty"`
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`
	AddMemberPermission    string     `json:"add_member_permission,omitempty"`
	ShareCardPermission    string     `json:"share_card_permission,omitempty"`
	AtAllPermission        string     `json:"at_all_permission,omitempty"`
	EditPermission         string     `json:"edit_permission,omitempty"`
	OwnerId                string     `json:"owner_id,omitempty"`
	JoinMessageVisibility  string     `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string     `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string     `json:"membership_approval,omitempty"`
}

type UpdateChatResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

type GetChatInfoResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    *GetChatInfoResponseBody `json:"data"`
}

type GetChatInfoResponseBody struct {
	Avatar                 string     `json:"avatar,omitempty"`
	Name                   string     `json:"name,omitempty"`
	Description            string     `json:"description,omitempty"`
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`
	OwnerId                string     `json:"owner_id,omitempty"`
	OwnerIdType            string     `json:"owner_id_type,omitempty"`
	AddMemberPermission    string     `json:"add_member_permission,omitempty"`
	ShareCardPermission    string     `json:"share_card_permission,omitempty"`
	AtAllPermission        string     `json:"at_all_permission,omitempty"`
	EditPermission         string     `json:"edit_permission,omitempty"`
	ChatMode               string     `json:"chat_mode,omitempty"`
	ChatType               string     `json:"chat_type,omitempty"`
	ChatTag                string     `json:"chat_tag,omitempty"`
	External               bool       `json:"external,omitempty"`
	TenantKey              string     `json:"tenant_key,omitempty"`
	JoinMessageVisibility  string     `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string     `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string     `json:"membership_approval,omitempty"`
	ModerationPermission   string     `json:"moderation_permission,omitempty"`
}

/*
	Message API Model
*/

type CreateMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

type CreateMessageResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *MessageItem `json:"data"`
}

type MessageItem struct {
	MessageID  string         `json:"message_id,omitempty"`
	RootID     string         `json:"root_id,omitempty"`
	ParentID   string         `json:"parent_id,omitempty"`
	MsgType    string         `json:"msg_type,omitempty"`
	CreateTime string         `json:"create_time,omitempty"`
	UpdateTime string         `json:"update_time,omitempty"`
	Deleted    bool           `json:"deleted,omitempty"`
	ChatID     string         `json:"chat_id,omitempty"`
	Sender     *MessageSender `json:"sender,omitempty"`
	Body       *MessageBody   `json:"body,omitempty"`
}

type MessageBody struct {
	Content string `json:"content,omitempty"`
}

type MessageSender struct {
	ID         string `json:"id,omitempty"`
	IDType     string `json:"id_type,omitempty"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

type ReceiveEventEncrypt struct {
	Encrypt string `json:"encrypt" form:"encrypt"`
}

type DecryptToken struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

type Event struct {
	Schema string      `json:"schema"`
	Header Header      `json:"header"`
	Event  interface{} `json:"event"`
}

type ReceiveMessageEvent struct {
	Schema string       `json:"schema"`
	Header Header       `json:"header"`
	Event  MessageEvent `json:"event"`
}

type MessageEvent struct {
	Sender  Sender  `json:"sender"`
	Message Message `json:"message"`
}

type Sender struct {
	SenderID   map[string]string `json:"sender_id"`
	SenderType string            `json:"sender_type"`
	TenantKey  string            `json:"tenant_key"`
}

type Message struct {
	MessageID   string     `json:"message_id"`
	RootID      string     `json:"root_id"`
	ParentID    string     `json:"parent_id"`
	CreateTime  string     `json:"create_time"`
	ChatID      string     `json:"chat_id"`
	ChatType    string     `json:"chat_type"`
	MessageType string     `json:"message_type"`
	Content     string     `json:"content"`
	Mentions    []*Mention `json:"mentions,omitempty"`
}

type Mention struct {
	Key       string  `json:"key,omitempty"`
	ID        *UserID `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	TenantKey string  `json:"tenant_key,omitempty"`
}

type MentionV1 struct {
	Key       string `json:"key,omitempty"`
	ID        string `json:"id,omitempty"`
	IDType    string `json:"id_type,omitempty"`
	Name      string `json:"name,omitempty"`
	TenantKey string `json:"tenant_key,omitempty"`
}

type UserID struct {
	UserID  string `json:"user_id,omitempty"`
	OpenID  string `json:"open_id,omitempty"`
	UnionID string `json:"union_id,omitempty"`
}

type GetMessageHistoryResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *GetMessageHistoryBody `json:"data"`
}

type GetMessageHistoryBody struct {
	HasMore   bool                            `json:"has_more"`
	PageToken string                          `json:"page_token,omitempty"`
	Items     []*GetMessageHistoryMessageItem `json:"items,omitempty"`
}

type GetMessageHistoryMessageItem struct {
	MessageID      string         `json:"message_id,omitempty"`
	RootID         string         `json:"root_id,omitempty"`
	ParentID       string         `json:"parent_id,omitempty"`
	MsgType        string         `json:"msg_type,omitempty"`
	CreateTime     string         `json:"create_time,omitempty"`
	UpdateTime     string         `json:"update_time,omitempty"`
	Deleted        bool           `json:"deleted,omitempty"`
	ChatID         string         `json:"chat_id,omitempty"`
	Sender         *MessageSender `json:"sender,omitempty"`
	Body           *MessageBody   `json:"body,omitempty"`
	Mentions       []*MentionV1   `json:"mentions,omitempty"`
	UpperMessageID string         `json:"upper_message_id,omitempty"`
}

type UploadImageResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    *UploadImageResponseBody `json:"data,omitempty"`
}

type UploadImageResponseBody struct {
	ImageKey string `json:"image_key"`
}

/*
	Message card
*/
type CardContent struct {
	Config   *CardConfig   `json:"config,omitempty"`
	Header   *CardHeader   `json:"header,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode,omitempty"`
	EnableForward  bool `json:"enable_forward,omitempty"`
}

type CardHeader struct {
	Title    *CardText `json:"title,omitempty"`
	Template string    `json:"template,omitempty"`
}

type CardElement struct {
	Tag          string       `json:"tag"`
	Text         *CardText    `json:"text,omitempty"`
	Fields       []*CardField `json:"fields,omitempty"`
	ImgKey       string       `json:"img_key,omitempty"`
	Mode         string       `json:"mode,omitempty"`
	Alt          *CardText    `json:"alt,omitempty"`
	CustomWidth  int          `json:"custom_width,omitempty"`
	CompactWidth int          `json:"compact_width,omitempty"`
	Preview      bool         `json:"preview,omitempty"`
	Title        *CardText    `json:"title,omitempty"`
}

type CardNote struct {
	Tag      string        `json:"tag,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type CardActionBlock struct {
	Tag     string        `json:"tag,omitempty"`
	Layout  string        `json:"layout,omitempty"`
	Actions []interface{} `json:"actions,omitempty"`
}

type CardText struct {
	Tag     string `json:"tag,omitempty"`
	Content string `json:"content,omitempty"`
	Lines   int    `json:"lines,omitempty"`
}

type CardField struct {
	IsShort bool      `json:"is_short,omitempty"`
	Text    *CardText `json:"text,omitempty"`
}

type CardButton struct {
	Tag     string            `json:"tag,omitempty"`
	Text    *CardText         `json:"text,omitempty"`
	Type    string            `json:"type,omitempty"`
	Value   map[string]string `json:"value,omitempty"`
	Confirm *CardConfirm      `json:"confirm,omitempty"`
}

type CardConfirm struct {
	Title *CardText `json:"title,omitempty"`
	Text  *CardText `json:"text,omitempty"`
}

type CardSelectMenu struct {
	Tag           string            `json:"tag"`
	PlaceHolder   *CardText         `json:"placeholder,omitempty"`
	InitialOption string            `json:"initial_option,omitempty"`
	Options       []*CardOption     `json:"options,omitempty"`
	Value         map[string]string `json:"value,omitempty"`
	Confirm       []*CardConfirm    `json:"confirm,omitempty"`
}

type CardOption struct {
	Text     *CardText `json:"text,omitempty"`
	Value    string    `json:"value"`
	URL      string    `json:"url,omitempty"`
	MultiURL *CardURL  `json:"multi_url,omitempty"`
}

type CardURL struct {
	URL        string `json:"url"`
	AndroidURL string `json:"android_url"`
	IosURL     string `json:"ios_url"`
	PcURL      string `json:"pc_url"`
}

type CardSplitLine struct {
	Tag string `json:"tag"`
}
