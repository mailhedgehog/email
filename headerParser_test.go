package email

import (
	"errors"
	"github.com/mailhedgehog/gounit"
	"net/mail"
	"testing"
	"time"
)

func TestParseMessageId_emptyOnError(t *testing.T) {
	hp := headerParser{nil, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsString("", hp.parseMessageId("<foo-bar >  "))

	hp = headerParser{&mail.Header{"test": []string{}}, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsString("", hp.parseMessageId("  <foo-bar>   "))
}

func TestParseMessageId_trimmedOnSuccess(t *testing.T) {
	hp := headerParser{}
	(*gounit.T)(t).AssertEqualsString("foo-bar", hp.parseMessageId("<foo-bar >  "))

	hp = headerParser{header: &mail.Header{"test": []string{}}}
	(*gounit.T)(t).AssertEqualsString("foo-bar", hp.parseMessageId("  <foo-bar>   "))
}

func TestParseMessageIdList_emptyOnError(t *testing.T) {
	hp := headerParser{nil, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsInt(0, len(hp.parseMessageIdList("<foo-bar > <foo-baz > <foo-bar > ")))

	hp = headerParser{&mail.Header{"test": []string{}}, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsInt(0, len(hp.parseMessageIdList("<foo-bar > <foo-baz > <foo-bar > ")))
}

func TestParseMessageIdList_listOnSuccess(t *testing.T) {
	hp := headerParser{}
	ids := hp.parseMessageIdList("<foo-bar > <foo-baz > <foo-bar > ")
	(*gounit.T)(t).AssertEqualsInt(3, len(ids))

	hp = headerParser{header: &mail.Header{"test": []string{}}}
	ids = hp.parseMessageIdList("<foo-bar > <foo-baz > <foo-bar > ")
	(*gounit.T)(t).AssertEqualsInt(3, len(ids))
}

func TestParseTime_emptyOnError(t *testing.T) {
	hp := headerParser{nil, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsString(time.Time{}.String(), hp.parseTime("Tue, 26 Dec 2023 15:04:05 -0800").String())

	hp = headerParser{&mail.Header{"test": []string{}}, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsString(time.Time{}.String(), hp.parseTime("Tue, 26 Dec 2023 15:04:05 -0800").String())
}

func TestParseTime_timeOnSuccess(t *testing.T) {
	hp := headerParser{}
	(*gounit.T)(t).AssertEqualsString("2023-12-26 15:04:05 -0800 -0800", hp.parseTime("Tue, 26 Dec 2023 15:04:05 -0800").String())

	hp = headerParser{header: &mail.Header{"test": []string{}}}
	(*gounit.T)(t).AssertEqualsString("2023-12-26 15:04:05 -0800 -0800", hp.parseTime("Tue, 26 Dec 2023 15:04:05 -0800").String())
}

func TestParseAddressList_emptyOnError(t *testing.T) {
	hp := headerParser{nil, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsInt(0, len(hp.parseAddressList("<foo@test.com>, \"John Snow\"<bar@test.com>")))

	hp = headerParser{&mail.Header{"test": []string{}}, errors.New("test error")}
	(*gounit.T)(t).AssertEqualsInt(0, len(hp.parseAddressList("<foo@test.com>, \"John Snow\"<bar@test.com>")))
}

func TestParseAddressList_listOnSuccess(t *testing.T) {
	hp := headerParser{}
	addresses := hp.parseAddressList("<foo@test.com>, \"John Snow\"<bar@test.com>")
	(*gounit.T)(t).AssertEqualsInt(2, len(addresses))

	hp = headerParser{header: &mail.Header{"test": []string{}}}
	addresses = hp.parseAddressList("<foo@test.com>, \"John Snow\"<bar@test.com>")
	(*gounit.T)(t).AssertEqualsInt(2, len(addresses))
	(*gounit.T)(t).AssertEqualsString("John Snow", addresses[1].Name)
}

func TestParseAddress_emptyOnError(t *testing.T) {
	hp := headerParser{nil, errors.New("test error")}
	(*gounit.T)(t).AssertNil(hp.parseAddress("\"John Snow\"<bar@test.com>"))

	hp = headerParser{&mail.Header{"test": []string{}}, errors.New("test error")}
	(*gounit.T)(t).AssertNil(hp.parseAddress("\"John Snow\"<bar@test.com>"))
}

func TestParseAddress_addressOnSuccess(t *testing.T) {
	hp := headerParser{}
	(*gounit.T)(t).AssertEqualsString("bar@test.com", hp.parseAddress("\"John Snow\"<bar@test.com>").Address)

	hp = headerParser{header: &mail.Header{"test": []string{}}}
	(*gounit.T)(t).AssertEqualsString("John Snow", hp.parseAddress("\"John Snow\"<bar@test.com>").Name)
}
