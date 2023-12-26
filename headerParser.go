package email

import (
	"net/mail"
	"strings"
	"time"
)

type headerParser struct {
	header *mail.Header
	err    error
}

func (hp headerParser) parseAddress(s string) (mailAddress *mail.Address) {
	if hp.err != nil {
		return nil
	}

	if strings.Trim(s, " \n") != "" {
		mailAddress, hp.err = mail.ParseAddress(s)

		return mailAddress
	}

	return nil
}

func (hp headerParser) parseAddressList(s string) (mailAddress []*mail.Address) {
	if hp.err != nil {
		return
	}

	if strings.Trim(s, " \n") != "" {
		mailAddress, hp.err = mail.ParseAddressList(s)
		return
	}

	return
}

func (hp headerParser) parseTime(s string) (mailTime time.Time) {
	s = strings.Trim(s, " ")
	if hp.err != nil || s == "" {
		return
	}

	formats := []string{
		time.RFC1123Z,
		"Mon, 2 Jan 2006 15:04:05 -0700",
		time.RFC1123Z + " (MST)",
		"Mon, 2 Jan 2006 15:04:05 -0700 (MST)",
	}

	for _, format := range formats {
		mailTime, hp.err = time.Parse(format, s)
		if hp.err == nil {
			return
		}
	}

	return
}

func (hp headerParser) parseMessageId(idString string) string {
	if hp.err != nil {
		return ""
	}

	return strings.Trim(idString, "<> ")
}

func (hp headerParser) parseMessageIdList(idString string) (result []string) {
	if hp.err != nil {
		return
	}

	for _, p := range strings.Split(idString, " ") {
		if strings.Trim(p, " \n") != "" {
			id := hp.parseMessageId(p)
			if len(id) > 0 {
				result = append(result, id)
			}
		}
	}

	return
}
