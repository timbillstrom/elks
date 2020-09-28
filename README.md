# 46Elks API wrapper built using Go

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2768bba67f3843b8a3ab20715211d970)](https://www.codacy.com/gh/timbillstrom/elks/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=timbillstrom/elks&amp;utm_campaign=Badge_Grade)

## Future features

- [x]   Send SMS
- [ ]   Send MMS
- [ ]   Place Automated Phone Calls
- [ ]   `// TODO… 🤔`

## Usage

* Get module: `go get -u github.com/timbillstrom/elks`

#### Send SMS

```go
package main

import (
    "os"
    "github.com/timbillstrom/elks"
)

func main() {
    client := elks.NewClient(
        os.Getenv("46_USERNAME"),
        os.Getenv("46_SECRET"),
        false, // No message will be sent when this is true.
    )
    res, err := client.SendMessage(&elks.SMS{
        From:    "Moose",
        To:      "+46…",
        Message: "Är du ute och älgar?",
    })
    ...
}
```
