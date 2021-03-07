package session_test

import (
	"testing"

	"github.com/cherya/standardfile/internal/model"
	"github.com/cherya/standardfile/internal/server/session"
	"github.com/stretchr/testify/assert"
)

func TestUserSupportsJWT(t *testing.T) {
	u := &model.User{Version: "003"}
	assert.True(t, session.UserSupportsJWT(u))

	u.Version = "004"
	assert.False(t, session.UserSupportsJWT(u))
}

func TestUserSupportsSessions(t *testing.T) {
	u := &model.User{Version: "003"}
	assert.False(t, session.UserSupportsSessions(u))

	u.Version = "004"
	assert.True(t, session.UserSupportsSessions(u))
}
