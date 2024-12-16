package route

import (
	"server/internal/queue/processor"
	"server/internal/queue/tasks"

	"github.com/hibiken/asynq"
)

func EmailTaskRoute(mux *asynq.ServeMux, processor *processor.EmailTaskProcessor) {
	mux.HandleFunc(tasks.TypeEmailVerification, processor.HandleVerificationEmail)
	mux.HandleFunc(tasks.TypeEmailForgotPassword, processor.HandleForgotPasswordEmail)
	mux.HandleFunc(tasks.TypeEmailPharmacistAccount, processor.HandlePharmacistAccountEmail)
}