package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateAccountNo() string {
	u := uuid.New()

	return strings.Join([]string{fmt.Sprint(u.ID())[:5], fmt.Sprint(time.Now().UnixMilli())[:5]}, "")
}
