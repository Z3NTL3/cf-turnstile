package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/z3ntl3/turnstile"
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