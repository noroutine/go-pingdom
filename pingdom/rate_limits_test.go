package pingdom

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseLimit(t *testing.T) {
    example1 := "Remaining: 394 Time until reset: 3589"
    exampleLimit := parseLimit(example1)

    assert.Equal(t, 394, exampleLimit.Remaining)
    assert.Equal(t, 3589, exampleLimit.TimeUntilReset)
    assert.Nil(t, exampleLimit.Error)
}
