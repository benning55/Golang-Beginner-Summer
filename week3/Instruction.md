# Week 3: Functions and Collections in Go

## Student Work Instructions

### Objective
Enhance the Income Expenses Tracker using Go, focusing on functions and slices.

### Setup
1. Will have two files named `main.go` and `operations.go`.

### Tasks

1. **Implement Basic Program Structure**
   - Refactor the `main` function to call functions for different operations.
   - Use an infinite loop for the menu system.

2. **Create Variables**
   - Declare `TotalIncome`, `TotalExpenses`, and `Transactions` as global variables.

3. **Implement Menu System**
   - Use a `for` loop to display the menu continuously.
   - Implement a `switch` statement to handle user choices:
     1. Add Income
     2. Add Expense
     3. View Summary
     4. View Transactions
     5. Exit

4. **Add Income Functionality**
   - Create `AddIncome` in `operations.go`.
   - Prompt user for income amount and validate input.
   - Update `TotalIncome` and append to `Transactions`.

5. **Add Expense Functionality**
   - Create `AddExpense` in `operations.go`.
   - Prompt user for expense amount and validate input.
   - Update `TotalExpenses` and append as a negative value to `Transactions`.

6. **Implement Summary View**
   - Create `ViewSummary` to display total income, total expenses, and net savings.

7. **Implement Transaction View**
   - Create `ViewTransactions` to display all transactions with their type (income or expense).

8. **Implement Exit Option**
   - Create `ExitProgram` to terminate the program using `os.Exit(0)`.

### Hints for Functions

- Use `fmt.Print` and `fmt.Scan` for user input and output.
- Use slices to manage transaction history.
- Ensure to handle invalid inputs gracefully.

### Running the Program

1. Open a terminal and navigate to your project directory.
2. Run the program using:
   ```sh
   go run .