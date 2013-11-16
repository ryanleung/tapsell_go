/*
MessageStore is responsible for handling messaging in a 
high-throughput, low-latency manner. 

TODO(ryan): add description of class, similar to the catalog_store.go in luv.
http://stackoverflow.com/questions/166884/array-vs-linked-list
Later on, we will implement linked list functionality, but for now,
use slices to make json conversion easier.
http://golang.org/pkg/container/list/#example_


*/

package base

import (
	"time"
	// "container/list"
)

const (
	MESSAGE_TYPE_INQUIRY	= "inquiry"
	MESSAGE_TYPE_DEFAULT	= "default"
)

type Message struct {
	// MessageId int
	SenderId int 		`json:"sender_id"`
	MessageChainId int 	`json:"msg_chain_id"`
	Content string 		`json:"content"`
	MessageType string 	`json:"msg_type"`
	TimeDate time.Time 	`json:"time"`
}

type MessageChain struct {
	InquirerId int 		`json:"inquirer_id"`
	SellerId int 		`json:"seller_id"`
	ListingId int 		`json:"listing_id"`
	Messages []*Message `json:"messages"`
}

type MessageStore struct {
	// message chain id => message chain
	MessageChains map[int]*MessageChain
}

func NewMessageStore() *MessageStore {
	ms := &MessageStore{
		MessageChains: make(map[int]*MessageChain),
	}
	return ms
}

func (ms *MessageStore) CreateMessageChain(inquirerId int, sellerId int, listingId int, content string, messageType string) {
	// instantiate a new message chain
	mc := &MessageChain{
		InquirerId: inquirerId,
		SellerId: sellerId,
		ListingId: listingId,
		Messages: make([]*Message, 0),
	}

	// create first message, who is always going to be from the inquirer
	message := &Message{
		SenderId: inquirerId,
		MessageChainId: len(ms.MessageChains),
		Content: content,
		MessageType: messageType,
		TimeDate: time.Now(),
	}

	// append message to messages in message chain
	mc.Messages = append(mc.Messages, message)

	// TODO(ryan): temporarily, just keep using msgChainId based on length of messageChains
	ms.MessageChains[len(ms.MessageChains)] = mc
}

func (ms *MessageStore) SendMessage(senderId int, messageChainId int, content string, messageType string) {
	// fetch message chain with messageChainId
	mc := ms.MessageChains[messageChainId]

	// construct a new message given the params
	message := &Message{
		SenderId: senderId,
		MessageChainId: len(ms.MessageChains),
		Content: content,
		MessageType: messageType,
		TimeDate: time.Now(),
	}

	// append message to messages in message chain
	mc.Messages = append(mc.Messages, message)
}