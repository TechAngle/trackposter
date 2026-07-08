package ytdlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	t.Parallel()

	url := "https://example.com"
	req := NewRequest(url, "--test-arg val")

	assert.Equal(t, url, req.url)
	assert.Contains(t, req.args, "--test-arg")
	assert.Equal(t, "val", req.args["--test-arg"])
}

func TestCommandRequest_AddArgument(t *testing.T) {
	t.Parallel()

	req := NewRequest("https://test.com")

	req.AddArgument("--new-arg value", false)
	assert.Equal(t, "value", req.args["--new-arg"])

	req.AddArgument("--new-arg updated", false)
	assert.Equal(t, "value", req.args["--new-arg"])

	req.AddArgument("--new-arg updated", true)
	assert.Equal(t, "updated", req.args["--new-arg"])
}

func TestCommandRequest_BuildArguments(t *testing.T) {
	t.Parallel()

	req := NewRequest("https://test.com")
	args := req.BuildArguments()

	assert.NotEmpty(t, args)
	assert.Contains(t, args, "https://test.com")
	assert.Contains(t, args, "-i")
}

func TestCommandRequest_Validate(t *testing.T) {
	t.Parallel()

	req := NewRequest("invalid-url")
	err := req.Validate()
	require.Error(t, err)

	req = NewRequest("https://valid.com")
	err = req.Validate()
	require.NoError(t, err)
}
