package response

import (
	"deall/cmd/lib/customError"
	errorCustomStatus "deall/pkg/error"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type rest struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success Serve Data With Success Status
func Success(w http.ResponseWriter, data interface{}) {
	result := rest{
		Error:   false,
		Status:  200,
		Message: "SUCCESS",
		Data:    data,
	}

	responses, err := json.Marshal(result)
	if err != nil {
		logrus.Error(err)
		Error(w, customError.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Status)
	w.Write(responses)
	return
}

// Error Serve Data With Error Status
func Error(w http.ResponseWriter, err error) {
	var result = rest{}
	customErrors, _ := err.(*errorCustomStatus.Error)

	if customErrors == nil {
		result = rest{
			Error:   true,
			Status:  500,
			Message: customError.ErrInternalServerError.Title,
			Data:    customError.ErrInternalServerError.Detail,
		}
	} else {
		status, err := strconv.Atoi(customErrors.Status)
		if err != nil {
			logrus.Error(err)
			Error(w, customError.ErrInternalServerError)
			return
		}
		result = rest{
			Error:   true,
			Status:  status,
			Message: customErrors.Title,
			Data:    customErrors.Detail,
		}
	}

	responses, err := json.Marshal(result)
	if err != nil {
		logrus.Error(err)
		Error(w, customError.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Status)
	w.Write(responses)
}
