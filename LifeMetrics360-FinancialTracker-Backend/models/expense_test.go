package models

import "testing"

func TestExpenseValidate(t *testing.T) {
	tests := []struct {
		name      string
		expense   Expense
		wantError bool
	}{
		{

			name: "valid expense",
			expense: Expense{
				Merchant: "Supermarket",
				Amount:   150,
				Date:     "2023-10-31",
				Category: "Groceries",
			},
			wantError: false,
		},
		{
			name: "invalid expense with empty merchant",
			expense: Expense{
				Amount:   150,
				Date:     "2023-10-31",
				Category: "Groceries",
			},
			wantError: true,
		},
		{
			name: "invalid expense with empty amount",
			expense: Expense{
				Merchant: "Power Company",
				Date:     "2023-10-31",
				Category: "Groceries",
			},
			wantError: true,
		},
		{
			name: "invalid expense with empty date",
			expense: Expense{
				Merchant: "Power Company",
				Amount:   10,
				Category: "Groceries",
			},
			wantError: true,
		},
		{
			name: "invalid expense with empty category",
			expense: Expense{
				Merchant: "Power Company",
				Amount:   10,
				Date:     "2023-10-31",
			},
			wantError: true,
		},
		{
			name: "invalid expense with invalid date",
			expense: Expense{
				Merchant: "Power Company",
				Amount:   10,
				Date:     "2023-10",
				Category: "Groceries",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.expense.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Expense.Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}

}
