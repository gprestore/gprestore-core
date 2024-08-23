package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/pkg/random"
)

func TestRandomString(t *testing.T) {
	r := random.String(8)
	log.Println(r)
}

func TestRandomNumber(t *testing.T) {
	r := random.Number(5)
	log.Println(r)
}
