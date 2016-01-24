package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

const (
	GETNEXTVOTE = "select * from getnextvote()"
	RESETVOTES  = "select resetVotes()"
	ARCHIVEVOTE = "select archiveVote($1::int)"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) InitDB(dbUrl string) error {
	var err error
	r.db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetNextVote() Vote {
	var vote Vote

	var voteid int
	var talkid string
	var userid int
	var rating int

	err := r.db.QueryRow(GETNEXTVOTE).Scan(&voteid, &talkid, &userid, &rating)

	parrot.Debug("GetNextVote: " + strconv.Itoa(voteid) + " " + talkid)

	if err != nil {
		parrot.Warn("No more rows", err)
		return Vote{ID: -1}
	}

	vote.ID = voteid
	vote.TalkId = talkid
	vote.UserId = userid
	vote.Rating = rating

	return vote
}

func (r *Repository) ResetVotes() error {
	_, err := r.db.Exec(RESETVOTES)
	return err
}

func (r *Repository) ArchiveVote(id int) error {
	_, err := r.db.Exec(ARCHIVEVOTE, id)
	return err
}
