package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/justinas/nosurf"
)

type templateData struct {
	CurrentYear int
	Snippet *models.Snippet
	Snippets []*models.Snippet
	Form any
	Flash string
	IsAuthenticated bool
	CSRFToken string
	User *models.User
}
