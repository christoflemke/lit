package lit

import (
	"fmt"
	"time"
)

type Author struct {
	Name  string
	Email string
	Time  time.Time
}

func (author *Author) ToString() string {
	_, offset := author.Time.Zone()
	offset = offset / 36
	// TODO: time is wrong
	return fmt.Sprintf("%s <%s> %d +0%d", author.Name, author.Email, author.Time.UnixMilli(), offset)
}
