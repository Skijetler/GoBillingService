// Code generated by ent, DO NOT EDIT.

package account

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCurrencyType holds the string denoting the currency_type field in the database.
	FieldCurrencyType = "currency_type"
	// FieldFrozen holds the string denoting the frozen field in the database.
	FieldFrozen = "frozen"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "accounts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldCurrencyType,
	FieldFrozen,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultFrozen holds the default value on creation for the "frozen" field.
	DefaultFrozen bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// CurrencyType defines the type for the "currency_type" enum field.
type CurrencyType string

// CurrencyType values.
const (
	CurrencyTypeRUB CurrencyType = "RUB"
	CurrencyTypeUSD CurrencyType = "USD"
	CurrencyTypeEUR CurrencyType = "EUR"
)

func (ct CurrencyType) String() string {
	return string(ct)
}

// CurrencyTypeValidator is a validator for the "currency_type" field enum values. It is called by the builders before save.
func CurrencyTypeValidator(ct CurrencyType) error {
	switch ct {
	case CurrencyTypeRUB, CurrencyTypeUSD, CurrencyTypeEUR:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for currency_type field: %q", ct)
	}
}
