package transaction_handlers

import (
	"net/http"

	"simple-mpesa/app/account"
	"simple-mpesa/app/auth"
	"simple-mpesa/app/errors"
	"simple-mpesa/app/models"
	"simple-mpesa/app/ports"
	"simple-mpesa/app/routing/responses"

	"github.com/gofiber/fiber/v2"
)

// Deposit allows user to deposit or credit their account.
func Deposit(txnAdapter ports.TransactorPort) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		var userDetails auth.UserAuthDetails
		if details, ok := ctx.Locals("userDetails").(auth.UserAuthDetails); !ok {
			return errors.Error{Code: errors.EINTERNAL}
		} else {
			userDetails = details
		}

		var p account.DepositParams
		_ = ctx.BodyParser(&p)

		depositor := models.TxnCustomer{
			UserType: userDetails.UserType,
			UserID: userDetails.UserID,
		}
		err := txnAdapter.Deposit(depositor, p.AgentNumber, p.Amount)
		if err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(responses.TransactionResponse())
	}
}

// Withdraw allows user to withdraw or debit their account.
func Withdraw(txnAdapter ports.TransactorPort) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		var userDetails auth.UserAuthDetails
		if details, ok := ctx.Locals("userDetails").(auth.UserAuthDetails); !ok {
			return errors.Error{Code: errors.EINTERNAL}
		} else {
			userDetails = details
		}

		var p account.WithdrawParams
		_ = ctx.BodyParser(&p)

		withdrawer := models.TxnCustomer{
			UserID:   userDetails.UserID,
			UserType: userDetails.UserType,
		}
		err := txnAdapter.Withdraw(withdrawer, p.AgentNumber, p.Amount)
		if err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(responses.TransactionResponse())
	}
}

func Transfer(txnAdapter ports.TransactorPort) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		return nil
	}
}
