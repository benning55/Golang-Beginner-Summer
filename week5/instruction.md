# Week 5: File I/O and Concurrency in Go

## Student Work Instructions

### Objective
Enhance the Income Expenses Tracker using Go, focusing on File I/O and Concurrency.

### Setup
1. You will have five files named `main.go`, `operations.go`, `transaction.go`, `fileio.go`, and `concurrency.go`.

### Tasks

1. **Implement File I/O (fileio.go)**
   - Create functions for saving and loading data:
     - `SaveData(tracker *IncomeExpenseTracker) error`
     - `LoadData(tracker *IncomeExpenseTracker) error`
   - Use JSON format for data serialization.
   - Handle potential errors during file operations.

2. **Implement Concurrency Features (concurrency.go)**
   - Create an auto-save function:
     - `AutoSave(tracker *IncomeExpenseTracker, interval time.Duration)`
   - Implement a real-time calculation function:
     - `RealTimeCalculation(tracker *IncomeExpenseTracker, updateChan chan<- string)`

3. **Modify Main Program (main.go)**
   - Load data from file at program startup.
   - Start the auto-save goroutine.
   - Initialize and start the real-time calculation goroutine.
   - Implement a mechanism to display real-time updates.

4. **Update Operations (operations.go)**
   - Modify `AddIncome` and `AddExpense` methods to trigger a save after each transaction.
   - Add methods to calculate total income and expenses.

5. **Implement Error Handling**
   - Add appropriate error handling for file operations and data processing.
   - Display meaningful error messages to the user.

### Key Concepts to Focus On

1. **File I/O**: 
   - Reading from and writing to files.
   - Using `encoding/json` for data serialization.

2. **Concurrency**: 
   - Understanding and using goroutines.
   - Implementing channels for communication between goroutines.
   - Using `select` statement for handling multiple channel operations.

3. **Error Handling**: 
   - Proper error checking and reporting in file operations.
   - Graceful handling of potential errors in concurrent operations.

### Hints

- Use `os.ReadFile` and `os.WriteFile` for file operations.
- Remember to close channels when they are no longer needed.
- Use `time.Ticker` for periodic operations like auto-save.
- Implement proper synchronization to avoid race conditions when accessing shared data.

### Running the Program

1. Open a terminal and navigate to your project directory.
2. Run the program using:
   ```sh
   go run main.go .

### Final Complete Program
1. Run the program using:
   ```sh
   go build .
2. Run the program using:
   ```sh
   ./income-expenses-tracker
