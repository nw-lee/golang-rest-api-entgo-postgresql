package model

import (
	"strings"

	"github.com/quavious/golang-docker-forum/ent"
	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *Post) PasswordHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)
	return string(hash), err
}

func (p *Post) PasswordCheck(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p.Password))
	return err
}

func HideInfoPost(p *ent.Post) {
	p.Password = "undefined"
	p.UserIP = strings.Join(strings.Split(p.UserIP, ".")[:2], ".")
}
