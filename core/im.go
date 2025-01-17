package core

import (
	"time"
)

type Sender interface {
	GetUserID() int
	GetChatID() int
	GetImType() string
	GetMessageID() int
	GetUsername() string
	IsReply() bool
	GetReplySenderUserID() int
	GetRawMessage() interface{}
	SetMatch([]string)
	SetAllMatch([][]string)
	GetMatch() []string
	GetAllMatch() [][]string
	Get(...int) string
	GetContent() string
	IsAdmin() bool
	IsMedia() bool
	Reply(...interface{}) (int, error)
	Delete() error
	Disappear(lifetime ...time.Duration)
}

type Edit int
type Replace int

var E Edit
var R Replace

type Faker struct {
	Message interface{}
	matches [][]string
}

func (sender *Faker) GetContent() string {
	return ""
}

func (sender *Faker) GetUserID() int {
	return 0
}

func (sender *Faker) GetChatID() int {
	return 0
}

func (sender *Faker) GetImType() string {
	return ""
}

func (sender *Faker) GetMessageID() int {
	return 0
}

func (sender *Faker) GetUsername() string {
	return ""
}

func (sender *Faker) IsReply() bool {
	return false
}

func (sender *Faker) GetReplySenderUserID() int {
	return 0
}

func (sender *Faker) GetRawMessage() interface{} {
	return nil
}

func (sender *Faker) SetMatch(ss []string) {

}
func (sender *Faker) SetAllMatch(ss [][]string) {

}

func (sender *Faker) GetMatch() []string {
	return nil
}

func (sender *Faker) GetAllMatch() [][]string {
	return nil
}

func (sender *Faker) Get(index ...int) string {
	return ""
}

func (sender *Faker) IsAdmin() bool {
	return true
}

func (sender *Faker) IsMedia() bool {
	return false
}

func (sender *Faker) Reply(msgs ...interface{}) (int, error) {
	return 0, nil
}

func (sender *Faker) Delete() error {
	return nil
}

func (sender *Faker) Disappear(lifetime ...time.Duration) {

}
