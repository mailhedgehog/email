package email

import (
	"io"
	"mime/multipart"
	"strings"
)

// Attachment represents file data with filename, content type and data (as an io.Reader)
type Attachment struct {
	Filename    string
	ContentType string
	Data        io.Reader
}

func isAttachment(part *multipart.Part) bool {
	return part.FileName() != ""
}

func decodeAttachment(part *multipart.Part) (at Attachment, err error) {
	filename := decodeMimeSentence(part.FileName())
	decoded, err := decodeContent(part, part.Header.Get("Content-Transfer-Encoding"))
	if err != nil {
		return
	}

	at.Filename = filename
	at.Data = decoded
	at.ContentType = strings.Split(part.Header.Get("Content-Type"), ";")[0]

	return
}
