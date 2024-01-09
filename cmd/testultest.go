package main

import (
	"bytes"
	"html"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"snippetbox.kaiyrbekovadilet.net/internal/models/mocks"
)

var csrfTokenRX = regexp.MustCompile(`<input type="hidden" name="csrf_token" value="(.+)" />`)

func extractCSRFToken(t *testing.T, body string) string {
	matches := csrfTokenRX.FindStringSubmatch(body)

	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}

	return html.UnescapeString(string(matches[1]))
}