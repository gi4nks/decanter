package main

import (
	"encoding/json"
)

type Vote struct {
	ID     int
	TalkId string
	UserId int
	Rating int
	Status string
}

func (c Vote) String() string {
	b, err := json.Marshal(c)

	if err != nil {
		parrot.Error("Warning", err)
		return "{}"
	}
	return string(b)
}

func (c Vote) ToVoteResource() VoteResource {
	return VoteResource{TalkId: c.TalkId, Rating: c.Rating}
}

func (c Vote) IsEmpty() bool {
	if c.ID == -1 {
		return true
	}
	return false
}

type VoteResource struct {
	TalkId string `json:"talkid"`
	Rating int    `json:"rating"`
}

func (c VoteResource) String() string {
	b, err := json.Marshal(c)

	if err != nil {
		parrot.Error("Warning", err)
		return "{}"
	}
	return string(b)
}
