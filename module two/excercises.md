# Lesson 2.1: Go Syntax and Structure

## Exercise: Age Classification Program

This program prompts the user for their age and then classifies them as a "Child," "Teenager," or "Adult."

```go
package main

import (
    "fmt"
)

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age < 13 {
        fmt.Println("You are a Child.")
    } else if age >= 13 && age < 18 {
        fmt.Println("You are a Teenager.")
    } else {
        fmt.Println("You are an Adult.")
    }
}
```

---

### Lesson 2.2: Functions in Go

#### Exercise: Factorial Function with Error Handling

This program defines a function to calculate the factorial of a number and includes error handling for negative inputs.

```go
package main

import (
    "errors"
    "fmt"
)

// factorial calculates the factorial of a given number.
// It returns an error if the number is negative.
func factorial(n int) (int, error) {
    if n < 0 {
        return -1, errors.New("cannot calculate factorial of a negative number")
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}

func main() {
    var num int
    fmt.Print("Enter a number to calculate its factorial: ")
    fmt.Scan(&num)

    result, err := factorial(num)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Factorial of %d is %d\n", num, result)
    }
}
```

---

### Lesson 2.3: Arrays, Slices, and Maps

#### Exercise: Student Grades Program

This program uses a map to store student names and grades. It allows the user to add new students and retrieve a student’s grade by name.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Initialize a map to store student names and their grades
    studentGrades := make(map[string]string)

    reader := bufio.NewReader(os.Stdin)

    // Adding new students to the map
    for {
        fmt.Print("Enter student name (or type 'done' to finish): ")
        name, _ := reader.ReadString('\n')
        name = strings.TrimSpace(name)

        if name == "done" {
            break
        }

        fmt.Print("Enter grade for ", name, ": ")
        grade, _ := reader.ReadString('\n')
        grade = strings.TrimSpace(grade)

        studentGrades[name] = grade
    }

    // Retrieve and print a student's grade
    fmt.Print("Enter a student's name to get their grade: ")
    studentName, _ := reader.ReadString('\n')
    studentName = strings.TrimSpace(studentName)

    grade, exists := studentGrades[studentName]
    if exists {
        fmt.Printf("%s's grade is %s\n", studentName, grade)
    } else {
        fmt.Printf("No grade found for %s\n", studentName)
    }
}
```

---

### Lesson 2.4: Error Handling in Go

#### Exercise: Simple ATM Machine Simulation

This program simulates a simple ATM machine that allows users to deposit and withdraw money, with error handling for insufficient funds.

```go
package main

import (
    "errors"
    "fmt"
)

// Account struct represents a simple bank account.
type Account struct {
    balance float64
}

// Deposit adds an amount to the account balance.
func (a *Account) Deposit(amount float64) {
    a.balance += amount
    fmt.Printf("Deposited: $%.2f\n", amount)
    fmt.Printf("Current balance: $%.2f\n", a.balance)
}

// Withdraw subtracts an amount from the account balance.
// It returns an error if there are insufficient funds.
func (a *Account) Withdraw(amount float64) error {
    if amount > a.balance {
        return errors.New("insufficient funds")
    }
    a.balance -= amount
    fmt.Printf("Withdrew: $%.2f\n", amount)
    fmt.Printf("Current balance: $%.2f\n", a.balance)
    return nil
}

func main() {
    account := Account{balance: 100.0} // Start with an initial balance

    fmt.Println("Welcome to the ATM")
    var choice int
    for {
        fmt.Println("1. Deposit")
        fmt.Println("2. Withdraw")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var depositAmount float64
            fmt.Print("Enter amount to deposit: ")
            fmt.Scan(&depositAmount)
            account.Deposit(depositAmount)

        case 2:
            var withdrawAmount float64
            fmt.Print("Enter amount to withdraw: ")
            fmt.Scan(&withdrawAmount)
            err := account.Withdraw(withdrawAmount)
            if err != nil {
                fmt.Println("Error:", err)
            }

        case 3:
            fmt.Println("Thank you for using the ATM. Goodbye!")
            return

        default:
            fmt.Println("Invalid option. Please choose 1, 2, or 3.")
        }
    }
}
```
