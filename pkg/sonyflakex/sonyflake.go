package sonyflakex

import (
	log "github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

var flake *sonyflake.Sonyflake

func NewSonyflake() *sonyflake.Sonyflake {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
	return flake
}

func NextID() uint64 {
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return id
}
