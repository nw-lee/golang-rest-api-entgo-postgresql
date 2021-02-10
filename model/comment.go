package model

import (
	"strings"

	"github.com/quavious/golang-docker-forum/ent"
)

func HideInfoComment(cmt *ent.Comment) {
	cmt.Password = "undefined"
	cmt.UserIP = strings.Join(strings.Split(cmt.UserIP, ".")[:2], ".")
}
