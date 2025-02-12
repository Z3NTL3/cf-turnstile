package turnstile_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/z3ntl3/turnstile"
)

// Output unit test:
// === RUN   TestClient
//
//	turnstile_test.go:33: res: &{Success:true ErrorCodes:[] ChallengeTS:2025-02-11T18:24:27.146Z Hostname:example.com}
//	turnstile_test.go:33: res: &{Success:false ErrorCodes:[invalid-input-response] ChallengeTS: Hostname:}
//	turnstile_test.go:33: res: &{Success:false ErrorCodes:[timeout-or-duplicate] ChallengeTS: Hostname:}
//
// --- PASS: TestClient (0.15s)
// PASS
// ok      github.com/SimpaiX-net/turnstile        0.585s
func TestClient(t *testing.T) {
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
			t.Fatalf("error: %s", err)
		}

		t.Logf("res: %+v\n", res)
	}
}
