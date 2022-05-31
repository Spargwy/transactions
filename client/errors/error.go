package errors

import (
	"log"
	"transactions/model"
)

func InternalErrorf(msg string, args ...interface{}) (*model.ErrorPayload, error) {
	viewerErr := model.ErrorPayload{
		Message: "Внутренняя ошибка. Попробуйте позже",
	}

	log.Printf("internal error: "+msg+"\n", args...)

	return &viewerErr, nil
}
