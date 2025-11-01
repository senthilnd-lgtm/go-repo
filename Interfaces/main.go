package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}
type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

type SimpleAccountNotifier struct{}

func (n SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("simple account notifier created", "username", account.Username, "email", account.Email)
	return nil
}

type EmailAccountNotifier struct{}

func (n EmailAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("New Email account notifier created ", "username", account.Username, "email", account.Email)
	return nil
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {

	var account Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode", "err", err)
		return
	}

	//if err := notifyAccountCreated(account); err != nil {

	// this is common logic, we are not touching this code every time for new notifier type
	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("Failed to notify ", "err", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)

}

// EMail
// SMS
// Telegram
// if we implement here then need to change each time below code , so used interface and methods
// func notifyAccountCreated(account Account) error {
// 	slog.Info("new account created", "username", account.Username, "Email", account.Email)

// 	return nil
// }

func main() {

	mux := http.NewServeMux()

	// accountHandler := &AccountHandler{
	// 	AccountNotifier: SimpleAccountNotifier{},
	// }

	accountHandler := &AccountHandler{
		AccountNotifier: EmailAccountNotifier{},
	}

	mux.HandleFunc("POST /account", accountHandler.handleCreateAccount)
	http.ListenAndServe(":3000", mux)
}
