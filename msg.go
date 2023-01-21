package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"mime"
	"mime/multipart"
	"net/mail"
)

type rawMessage struct {
	id      int
	content string
}

var messageTemplate = template.Must(template.New("").Parse(`
{{ define "message" }}
  <div class="message">
    {{ if .Reblog }}
      <div class="reblog">
        {{ template "message" .Reblog }}
      </dev>
    {{ end }}
  </div>
  <div class="header">
    {{ .Author.Name }}
    <a href="{{ .Author.URL }}">@{{ .Author.DisplayName }}</a>
    <a href="{{ .URL }}" class="message_date">{{ .Date }}</a>
  </div>
  <div class="message_content">
    <div class="message_text">{{ .Text }}</div>
    {{ if .Media }}
       <div class="media">
         {{ range $i, $media := .Media }}
           {{ if $media.IsImage }}
             <p><img src="{{ $media.URL }}"><p> {{ $media.Description }}
           {{ else }}
             <a href="{{ $media.URL }}">video</a>
           {{ end }}
         {{ end }}
       </div>
    {{ end }}
  </div>
{{ end }}
<html>
  <head>
    <style>
      * { font-family: sans-serif; }
      .reblog { margin-left: 50px; margin-bottom: 20px; }
      .header { margin-bottom: 10px; }
      a { color: #656565; }
      .message_date { font-size: 50%; }
      .message_content { margin-left: 5px; }
    </style>
  </head>
  <body>
    {{ template "message" . }}
  </body>
</html>
`))

type msgAccount struct {
	Account     string `json:"acct"`
	DisplayName string `json:"display_name"`
	URL         string `json:"url"`
}

type msgMediaAttachment struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type msg struct {
	CreatedAt string `json:"created_at"`
	URL       string `json:"url"`
	Content   string `json:"content"`

	Account msgAccount `json:"account"`

	MediaAttachments []msgMediaAttachment `json:"media_attachments"`

	Reblog *msg `json:"reblog"`
}

func templateAuthor(account msgAccount) map[string]any {
	return map[string]any{
		"DisplayName": account.DisplayName,
		"URL":         account.URL,
	}
}

func templateReblog(msg *msg) map[string]any {
	if msg == nil {
		return nil
	}
	return templateMsg(*msg)
}

func templateMedia(media []msgMediaAttachment) []map[string]any {
	var out []map[string]any
	for _, m := range media {
		out = append(out, map[string]any{
			"IsImage":     m.Type == "image",
			"URL":         m.URL,
			"Description": m.Description,
		})
	}
	return out
}

func templateMsg(msg msg) map[string]any {
	return map[string]any{
		"Author": templateAuthor(msg.Account),
		"Reblog": templateReblog(msg.Reblog),
		"URL":    msg.URL,
		"Date":   msg.CreatedAt,
		"Text":   template.HTML(msg.Content),
		"Media":  templateMedia(msg.MediaAttachments),
	}
}

func formatMessage(rawMsg rawMessage, from, to *mail.Address) []byte {
	var msg msg
	if err := json.Unmarshal([]byte(rawMsg.content), &msg); err != nil {
		fatal("error parsing message %d: %v", rawMsg.id, err)
	}

	var buf bytes.Buffer
	if err := messageTemplate.ExecuteTemplate(&buf, "", templateMsg(msg)); err != nil {
		fatal("error formatting message %d: %v", rawMsg.id, err)
	}

	var messageBuf bytes.Buffer
	multipart := multipart.NewWriter(&messageBuf)

	partWriter, err := multipart.CreatePart(map[string][]string{"Content-Type": {"text/html"}})
	if err != nil {
		fatal("error creating message part: %v", err)
	}
	if _, err := partWriter.Write(buf.Bytes()); err != nil {
		fatal("error formatting message %d: %v", rawMsg.id, err)
	}

	if err := multipart.Close(); err != nil {
		fatal("error formatting message: %v", err)
	}

	return []byte("From: " + from.String() + "\r\n" +
		"To: " + to.String() + "\r\n" +
		"Subject: " + mime.QEncoding.Encode("UTF-8", msg.Account.DisplayName) + "\r\n" +
		"Content-Type: " + multipart.FormDataContentType() + "\r\n" +
		"\r\n" +
		messageBuf.String())
}
