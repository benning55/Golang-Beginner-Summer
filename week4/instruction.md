# Week 4: Structs, Interfaces, and Error Handling

## Student Work Instructions

### Objective
Enhance the Income Expenses Tracker using Go, focusing on structs, interfaces, and error handling.

### Setup
1. You will have three files named `main.go`, `transaction.go`, and `operations.go`.

### Tasks

1. **Implement Basic Program Structure**
   - Use the provided `main` function in `main.go`.
   - Implement an infinite loop for the menu system.

2. **Create Structs and Interfaces**
   - Define `Transaction`, `IncomeTransaction`, and `ExpenseTransaction` structs in `transaction.go`.
   - Implement the `FinancialOperation` interface with `Execute` and `GetAmount` methods.

3. **Create Variables**
   - Use the `IncomeExpenseTracker` struct to manage transactions and categories.

4. **Implement Menu System**
   - Use the provided `displayMenu` function.
   - Implement a `switch` statement to handle user choices:
     1. Add Income
     2. Add Expense
     3. View Summary
     4. View Transactions
     5. View Expenses By Category
     6. Exit

5. **Add Income Functionality**
   - Implement the `AddIncome` method in `operations.go`.
   - Prompt user for income amount and validate input.
   - Update `TotalIncome` and append to `Transactions`.

6. **Add Expense Functionality**
   - Implement the `AddExpense` method in `operations.go`.
   - Prompt user for expense amount and validate input.
   - Use `getCategory` function to get the expense category.
   - Update `TotalExpenses`, `Transactions`, and `ExpenseCategories`.

7. **Implement Summary View**
   - Implement the `ViewSummary` method to display total income, total expenses, and net savings.

8. **Implement Transaction View**
   - Implement the `ViewTransactions` method to display all transactions with their type (income or expense).

9. **Implement Expenses By Category View**
   - Implement the `ViewExpensesByCategory` method to display expenses grouped by category.

10. **Implement Exit Option**
    - Use the provided `ExitProgram` function to terminate the program.

### Key Concepts to Focus On

1. **Structs**: 
   - Implement and use structs to group related data together.
   - Understand how to define and use struct methods.

2. **Interfaces**: 
   - Define and implement interfaces to create flexible and reusable code.
   - Understand how interfaces can be used to define a set of actions.

3. **Error Handling**: 
   - Implement basic error handling in functions.
   - Create custom error types and use them in your code.

### Hints

- Use `fmt.Print` and `fmt.Scan` for user input and output.
- Remember to update both `Transactions` and `ExpenseCategories` when adding an expense.
- When displaying transactions, use conditional statements to differentiate between income and expenses.

### Running the Program

1. Open a terminal and navigate to your project directory.
2. Run the program using:
   ```sh
   go run .