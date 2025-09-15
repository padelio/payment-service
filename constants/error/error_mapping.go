package error

import (
	errField "payment-service/constants/error/field"
	errFieldSchedule "payment-service/constants/error/fieldschedule"
	errTime "payment-service/constants/error/time"
)

func ErrorMapping(err error) bool {
	var (
		GeneralErrors       = GeneralErrors
		FieldErrors         = errField.FieldErrors
		FieldScheduleErrors = errFieldSchedule.FieldScheduleErrors
		TimeErrors          = errTime.TimeErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, FieldErrors...)
	allErrors = append(allErrors, FieldScheduleErrors...)
	allErrors = append(allErrors, TimeErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
