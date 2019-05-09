package pingdom

import (
    "fmt"
    "net/http"
    "sync"
)

type RateLimitsHolder struct {
    sync.Mutex
    RateLimits
}

type RateLimits struct {
    Short Limit
    Long Limit
}

type Limit struct {
    Remaining int
    TimeUntilReset int
    Error error
}

func (rl *RateLimitsHolder) UpdateFromResponse (resp *http.Response) {
    rl.Lock()
    defer rl.Unlock()

    if resp == nil {
        return
    }

    rl.Long = parseLimit(resp.Header.Get("Req-Limit-Long"))
    rl.Short = parseLimit(resp.Header.Get("Req-Limit-Short"))
}

func parseLimit(val string) Limit {
    if val == "" {
        return Limit{ Error: fmt.Errorf("no value") }
    }

    var remaining int
    var timeUntilReset int

    if n, err := fmt.Sscanf(val, "Remaining: %d Time until reset: %d", &remaining, &timeUntilReset); err != nil {
        return Limit{ Error: err }
    } else {
        if n == 2 {
            return Limit{
                Remaining: remaining,
                TimeUntilReset: timeUntilReset,
            }
        } else {
            return Limit{ Error: fmt.Errorf("not enough values") }
        }
    }
}

func (rl *RateLimitsHolder) Get() RateLimits {
    rl.Lock()
    defer rl.Unlock()

    return rl.RateLimits
}
