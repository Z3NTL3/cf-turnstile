# Go Bindigs for Cloudflare's Turnstile
This package provides you with a high-level, minimal and composable set of APIs for using Cloudflare's Turnstile for your Go backend services.

> Examples can be found in the sections below and unit tests reside at ``turnstile_test.go``

### Install
> ``go get github.com/z3ntl3/turnstile``

and after that you can import and use the package

### Running unit tests
``go test -v``

### Examples

```go
package main

import (
	"fmt"
	"net/http"
	"strings"

	turnstile "github.com/z3ntl3/cf-turnstile"
)

func main() {
	client := turnstile.TurnstileClient{
		Client: &http.Client{},
	}

	testingPlayground := []string{
		"passes:1x0000000000000000000000000000000AA",
		"fails:2x0000000000000000000000000000000AA",
		"spent:3x0000000000000000000000000000000AA",
	}

	for _, task := range testingPlayground {
		secret := strings.Split(task, ":")[1]

		res, err := client.Verify(turnstile.VerifyOpts{
			Secret:   secret,
			Response: "XXXX.DUMMY.TOKEN.XXXX",
		})
		if err != nil {
			fmt.Printf("error: %s", err)
		}

		fmt.Printf("res: %+v\n", res)
	}

	// res: &{Success:true ErrorCodes:[] ChallengeTS:2025-02-12T09:32:03.767Z Hostname:example.com}
	// res: &{Success:false ErrorCodes:[invalid-input-response] ChallengeTS: Hostname:}
	// res: &{Success:false ErrorCodes:[timeout-or-duplicate] ChallengeTS: Hostname:}
}
```
