package env

import "time"

func (e *Env) Now() time.Time {
	return time.Now()
}
