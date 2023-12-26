## MailHedgehog email message parser

Parse email string to appropriate struct for easy access data and manipulate it.
More information about fields you can find [here](rfc5322.txt)

NOTE: this package is fork of https://github.com/ArkaGPL/parsemail with small changes required for MailHedgehog

## Usage

```go
package main

import (
    "github.com/mailhedgehog/email"
)

func main() {
    email, err := email.Parse(ioReaderWithEmailData)
    // or
    email, err := email.Parse(strings.NewReader(emailRawStringData))
    if err != nil {
        // handle error
    }

    fmt.Print(email.Subject)
}
```

## Development

```shell
go mod tidy
go mod verify
go mod vendor
go test --cover
```

## Credits

- [Initial package code](https://github.com/ArkaGPL/parsemail)
- [![Think Studio](https://yaroslawww.github.io/images/sponsors/packages/logo-think-studio.png)](https://think.studio/)
