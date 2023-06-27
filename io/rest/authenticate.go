package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/okian/servo/v2/kv"
	"github.com/okian/servo/v2/lg"
	"strings"
	"time"
)

var unauthorized = errors.New("unauthorized")

func authenticateParkingAdmin(ctx context.Context, sessionToken string) (string, string, error) {
	if len(sessionToken) == 0 {
		return "", "", unauthorized
	}
	var data string
	err := kv.Get(ctx, sessionToken, &data)
	if err != nil {
		return "", "", err
	}
	dataParts := strings.Split(data, "_")
	if len(dataParts) < 2 {
		return "", "", unauthorized
	}
	if dataParts[0] == "pAdmin" {
		sessionToken := uuid.NewString()
		lg.Debug(sessionToken)
		kv.Set(ctx, sessionToken, fmt.Sprintf("pAdmin_%s", dataParts[1]), time.Second*120)
		return dataParts[1], sessionToken, nil
	} else {
		return "", "", unauthorized
	}
}

func authenticateSystemAdmin(ctx context.Context, sessionToken string) (string, string, error) {
	if len(sessionToken) == 0 {
		return "", "", unauthorized
	}
	var data string
	err := kv.Get(ctx, sessionToken, &data)
	if err != nil {
		return "", "", err
	}
	dataParts := strings.Split(data, "_")
	if len(dataParts) < 2 {
		return "", "", unauthorized
	}
	if dataParts[0] == "sAdmin" {
		sessionToken := uuid.NewString()
		lg.Debug(sessionToken)
		kv.Set(ctx, sessionToken, fmt.Sprintf("sAdmin_%s", dataParts[1]), time.Second*120)
		return dataParts[1], sessionToken, nil
	} else {
		return "", "", unauthorized
	}
}

func authenticateUser(ctx context.Context, sessionToken string) (string, string, error) {
	if len(sessionToken) == 0 {
		return "", "", unauthorized
	}
	var data string
	err := kv.Get(ctx, sessionToken, &data)
	if err != nil {
		return "", "", err
	}
	dataParts := strings.Split(data, "_")
	if len(dataParts) < 2 {
		return "", "", unauthorized
	}
	if dataParts[0] == "user" {
		sessionToken := uuid.NewString()
		lg.Debug(sessionToken)
		kv.Set(ctx, sessionToken, fmt.Sprintf("user_%s", dataParts[1]), time.Second*120)
		return dataParts[1], sessionToken, nil
	} else {
		return "", "", unauthorized
	}
}
