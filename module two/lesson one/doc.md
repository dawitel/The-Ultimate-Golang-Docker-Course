# Module 2: Go Fundamentals

## Overview

Welcome to Module 2 of *The Ultimate Golang + Docker Course for Beginners*! This module takes a deep dive into the core elements of the Go programming language. We’ll cover Go’s unique syntax, explore various function types and how to use them effectively, understand data structures like arrays, slices, and maps, and go over error handling to ensure robust applications. By the end of this module, you’ll be comfortable with Go fundamentals and ready to build small, structured applications.

---

## Learning Objectives

By the end of this module, you will be able to:

- Master Go’s syntax, variable types, and control structures.
- Create various types of functions in Go, understanding their roles and behaviors.
- Effectively work with data structures, including arrays, slices, and maps.
- Implement best practices for error handling, understanding the `error` interface and custom error creation.

---

## Lessons Overview

This module is structured into four main lessons:

1. **Lesson 2.1: Go Syntax and Structure**
2. **Lesson 2.2: Functions in Go**
3. **Lesson 2.3: Arrays, Slices, and Maps**
4. **Lesson 2.4: Error Handling in Go**

Each lesson combines theoretical explanations, code samples, and exercises to reinforce learning.

---

## Lesson 2.1: Go Syntax and Structure

### What You’ll Learn

- **Go Project Structure**: Understanding Go’s project organization and workflow.
- **Variable and Constant Declarations**: How to declare and initialize variables and constants.
- **Control Structures**: How to use `if`, `for`, and `switch` statements.

### Key Concepts

Go’s syntax is designed for simplicity and clarity, helping to create clean, maintainable code. Unlike languages with complex syntax, Go promotes readability and simplicity, making it ideal for beginner and advanced developers alike.

#### Code Structure and Workflow

A standard Go project organizes files and folders to keep code modular and maintainable. Each project has a `main.go` file, which contains the entry point of the program (the `main()` function). Go also uses packages to organize reusable code, and you’ll often find Go code organized into folders like `cmd/`, `pkg/`, and `internal/`.

#### Variable Declaration

Go has a few unique features for declaring variables:

- **Explicit Declaration**: Use `var` to declare variables explicitly. This is helpful for readability, especially when specifying types.

  ```go
  var age int = 25
  ```

- **Short Declaration**: Use `:=` to declare and initialize variables in one step, which Go infers based on the assigned value.

  ```go
  name := "John Doe"
  ```

- **Constants**: Constants are declared with `const` and are immutable after declaration.

  ```go
  const Pi = 3.14159
  ```

#### Control Structures

Control structures like `if`, `for`, and `switch` let you control the flow of execution:

1. **If statements**:

   ```go
   if age > 18 {
       fmt.Println("Adult")
   } else {
       fmt.Println("Not an adult")
   }
   ```

2. **For loops**: Go’s only looping construct is the `for` loop, which can act as a traditional loop, a while loop, or an infinite loop.

   ```go
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }
   ```

3. **Switch statements**: Go’s `switch` is more flexible than in other languages, allowing you to handle multiple cases at once and omit the `break` statement.

   ```go
   switch day := 3; day {
   case 1:
       fmt.Println("Monday")
   case 2:
       fmt.Println("Tuesday")
   default:
       fmt.Println("Other day")
   }
   ```

---

## Lesson 2.2: Functions in Go

### What You’ll Learn

- **Function Types**: Basic functions, functions with multiple return values, anonymous functions, and higher-order functions.
- **Parameters and Return Values**: Passing values to functions and understanding scope and closures.
- **Error Handling in Functions**: Using the `error` type in functions.

### Key Concepts

Functions are one of Go’s key building blocks, allowing you to encapsulate logic and make code reusable and modular. Go has several types of functions, each serving unique purposes in the code.

#### Types of Functions

1. **Basic Functions**: A simple function that performs a task and returns a result. Functions in Go are declared using the `func` keyword.

   ```go
   func greet(name string) string {
       return "Hello, " + name
   }
   ```

2. **Functions with Multiple Return Values**: Go allows functions to return multiple values, a common pattern for functions that return a result and an error.

   ```go
   func divide(a, b float64) (float64, error) {
       if b == 0 {
           return 0, errors.New("cannot divide by zero")
       }
       return a / b, nil
   }
   ```

3. **Anonymous Functions**: These are functions without a name, useful for short tasks, especially as arguments to higher-order functions.

   ```go
   add := func(x, y int) int {
       return x + y
   }
   fmt.Println(add(3, 4)) // Output: 7
   ```

4. **Higher-Order Functions**: Go functions can be passed as arguments or returned from other functions. This makes it easier to create reusable code and manage complex logic.

   ```go
   func applyOperation(x, y int, operation func(int, int) int) int {
       return operation(x, y)
   }
   ```

### Exercise

Write a function that calculates the factorial of a number. Include error handling for negative inputs. This will reinforce understanding of basic function structures and error handling.

---

## Lesson 2.3: Arrays, Slices, and Maps

### What You’ll Learn

- **Arrays**: Fixed-length collections of a specific data type.
- **Slices**: Dynamic, flexible data structures for managing collections.
- **Maps**: Key-value pairs, like dictionaries in other languages.

### Key Concepts

Go’s data structures allow you to manage collections of data efficiently, with optimized memory usage.

#### Arrays

Arrays have a fixed length and store elements of a single data type. Their immutability makes them useful for static datasets.

   ```go
   var numbers [5]int
   numbers[0] = 1
   ```

#### Slices

Slices are more flexible and can grow dynamically. They’re built on top of arrays but can resize as needed.

   ```go
   fruits := []string{"Apple", "Banana", "Cherry"}
   fruits = append(fruits, "Date")
   ```

#### Maps

Maps store data in key-value pairs, making it easy to retrieve and manipulate data based on keys.

   ```go
   person := map[string]string{"name": "Alice", "city": "Wonderland"}
   fmt.Println(person["name"]) // Output: Alice
   ```

### Exercise

Create a program that uses a map to store student names and grades. Allow the user to add new students and retrieve a student’s grade by name.

---

## Lesson 2.4: Error Handling in Go

### What You’ll Learn

- **Error as a Type**: Go treats errors as values, meaning functions can return errors alongside data.
- **Custom Error Types**: Go allows you to define custom error types for more descriptive error handling.
- **Best Practices for Error Handling**: Tips for clear, maintainable error handling.

### Key Concepts

In Go, errors are values that can be passed around just like other data. This philosophy encourages developers to handle errors explicitly and prevent runtime crashes.

#### Basic Error Handling

Use the `error` interface to handle and manage errors gracefully.

   ```go
   func divide(a, b float64) (float64, error) {
       if b == 0 {
           return 0, errors.New("cannot divide by zero")
       }
       return a / b, nil
   }
   ```

#### Custom Error Types

Define custom error types for errors that require specific handling, like insufficient funds in a bank transaction.

   ```go
   type InsufficientFundsError struct {
       Amount float64
   }

   func (e *InsufficientFundsError) Error() string {
       return fmt.Sprintf("insufficient funds: $%.2f needed", e.Amount)
   }
   ```

### Exercise

Write a program that simulates a simple ATM machine. It should allow users to deposit and withdraw money, with error handling for insufficient funds.

---

## Project Assignments

1. **To-Do List Application**: Build a command-line to-do list application, implementing functions for adding, completing, and deleting tasks.
2. **Number Guessing Game**: Create a number guessing game where the player tries to guess a randomly generated number within a set number of attempts.

---

## Module Summary

In this module, you covered essential Go elements, gaining skills to:

- Write code using Go’s syntax and control structures.
- Create different types of functions and understand function returns.
- Work with arrays, slices, and maps effectively.
- Handle errors gracefully using Go’s `error` type and custom error handling.
