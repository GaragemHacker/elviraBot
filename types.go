package main

//"fmt"
//"encoding/json"

//Struct to config bot
type Config struct {
	GaragemChatId     int    `json:"garagemChatId`
	BotToken          string `json:"botToken"`
	SecretPath        string `json:"secretPath"`
	ReceiveStatusPath string `json:"receiveStatusPath"`
}

//User This object represents a Telegram user or bot.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

//Chat This object represents a chat.
type Chat struct {
	ID int `json:"id"`
	//	Type      string `json:"type"`
	//	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FiliSize int    `json:"file_size"`
}

//Audio This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FiliSize  int    `json:"file_size"`
}

//Document This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

//Sticker This object represents a sticker.
type Sticker struct {
	FileID   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   string    `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}

//Video This object represents a video file.
type Video struct {
	FileID   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   string    `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

//Voice This object represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

//Contact This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
}

//Location This object represents a point on the map.
type Location struct {
	Longitude string `json:"longitude"`
	Latitude  int    `json:"latitude"`
}

//Venue This object represents a venue.
type Venue struct {
	Location    Location `json:"location"`
	Title       string   `json:"title"`
	Address     string   `json:"address"`
	ForsquareID string   `json:"forsquare_id"`
}

//UserProfilePhotos This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

//File This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
//It is guaranteed that the link will be valid for at least 1 hour.
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

//ReplyKeyboardMarkup This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool           `json:"resize_keyboard"`
	OneTimeKeyboard bool           `json:"one_time_keyboard"`
	Selective       bool           `json:"selective"`
}

//KeyboardButton This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

// removed telegram api
/*type From struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
*/

//InlineQuery This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"User"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
}

//ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageID string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

//CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	Data            string  `json:"data"`
}

//Update This object represents an incoming update.
type Update struct {
	UpdateID           int                `json:"update_id"`
	Message            Message            `json:"message"`
	EditedMessage      Message            `json:"edited_message"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
}

//Message This object represents a message.
type Message struct {
	MessageID             int             `json:"message_id"`
	From                  User            `json:"from"`
	Date                  int64           `json:"date"`
	Chat                  Chat            `json:"chat"`
	ForwardFrom           User            `json:"forward_from"`
	ForwardFromChat       Chat            `json:"forward_from_chat"`
	ForwardDate           int64           `json:"forward_date"`
	ReplyToMessage        *Message        `json:"reply_to_message"`
	EditDate              int64           `json:"edit_date"`
	Text                  string          `json:"text"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 Audio           `json:"audio"`
	Document              Document        `json:"document"`
	Photo                 []PhotoSize     `json:"photo"`
	Sticker               Sticker         `json:"sticker"`
	Video                 Video           `json:"video"`
	Voice                 Voice           `json:"voice"`
	Caption               string          `json:"caption"`
	Contact               Contact         `json:"contact"`
	Location              Location        `json:"location"`
	Venue                 Venue           `json:"venue"`
	NewChatMember         User            `json:"new_chat_member"`
	LeftChatMember        User            `json:"left_chat_member"`
	NewChatTitle          string          `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`
	GroupChatCreated      bool            `json:"group_chat_created"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created"`
	ChannelChatCreated    bool            `json:"channel_chat_created"`
	MigrateToChatID       int             `json:"migrate_to_chat_id"`
	MigrateFromChatID     int             `json:"migrate_from_chat_id"`
	PinnedMessage         *Message        `json:"pinned_message"`
}

//MessageEntity This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	URL    string `json:"url"`
	User   User   `json:"user"`
}
