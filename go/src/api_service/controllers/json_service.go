package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"api_service/base"
	"strconv"
	"log"
)

type JsonService struct {
	MessageStore *base.MessageStore
}

func (js *JsonService) Serve(port int) {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hI!")
		})

	http.HandleFunc(
		"/message/send",
		func(w http.ResponseWriter, r *http.Request) {
			handleSendMessage(w, r, js.MessageStore)
		})

	http.HandleFunc(
		"/message/create",
		func(w http.ResponseWriter, r *http.Request) {
			handleCreateMessageChain(w, r, js.MessageStore)
		})

	http.HandleFunc(
		"/messages/fetch_by_user",
		func(w http.ResponseWriter, r *http.Request) {
			handleFetchMessages(w, r, js.MessageStore)
		})

	http.HandleFunc(
		"/messages",
		func(w http.ResponseWriter, r *http.Request) {
			handleMessages(w, r, js.MessageStore)
		})


	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func handleSendMessage(w http.ResponseWriter, r *http.Request, ms *base.MessageStore) {
	senderId, err := strconv.Atoi(r.FormValue("sender_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}
	msgChainId, err := strconv.Atoi(r.FormValue("message_chain_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}
	content := r.FormValue("content")
	messageType := r.FormValue("msg_type")

	ms.SendMessage(senderId, msgChainId, content, messageType)
	fmt.Fprintf(w, "Message Sent %v", ms.MessageChains[msgChainId])
}

func handleCreateMessageChain(w http.ResponseWriter, r *http.Request, ms *base.MessageStore) {
	inquirerId, err := strconv.Atoi(r.FormValue("inquirer_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}
	sellerId, err := strconv.Atoi(r.FormValue("seller_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}
	listingId, err := strconv.Atoi(r.FormValue("listing_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}
	content := r.FormValue("content")
	messageType := r.FormValue("msg_type")

	ms.CreateMessageChain(inquirerId, sellerId, listingId, content, messageType)
	fmt.Fprintf(w, "Created Chain %v", ms.MessageChains)
}

func handleFetchMessages(w http.ResponseWriter, r *http.Request, ms *base.MessageStore) {

	// TODO(ryan): for now, just get all the messages
	// start, err := strconv.Atoi(r.FormValue("start"))
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "Bad parameters")
	// }
	// limit, err := strconv.Atoi(r.FormValue("limit"))
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "Bad parameters")
	// }

	userId, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Bad parameters")
		return
	}

	// TODO(ryan): right now, it is traversing through all the messages 
	// and seeing if the message chain belongs to the user.
	// Later on, once we figure out an "inbox system" (user model will
	// contain all of the message chain ids he owns, so user can just
	// query based on message chain ids)
	// Luckily, traversal in go is lightning quick as opposed to ruby.
	msgChainsToReturn := make([]*base.MessageChain, 0)
	for _, messageChain := range ms.MessageChains {
		if messageChain.SellerId == userId {
			msgChainsToReturn = append(msgChainsToReturn, messageChain)
		}
	}
	j, err := json.Marshal(msgChainsToReturn)
	if err == nil {
		fmt.Fprintf(w, string(j))
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Could not encode JSON. err: %v", err)
	}

	// Log the request
	log.Printf(
		"%v Message Chains returned to user with id %v", 
		len(msgChainsToReturn), userId)
}

func handleMessages(w http.ResponseWriter, r *http.Request, ms *base.MessageStore) {

}
