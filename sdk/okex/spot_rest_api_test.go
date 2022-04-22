package okex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSpotAccounts(t *testing.T) {
	c := NewTestCli