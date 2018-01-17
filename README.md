# GoThreading
# CS 341 Final Project: Go Programming Language

Final project for CS341 FALL 2017.
Created by Casey Charlesworth.
Program written in Go.
Examples of concurrency in Go using Go Routines.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.


### Prerequisites

Go needs to be installed
Use of terminal/ cmd prompt

### Installing


1. Go to https://golang.org/dl/ to install Go on your OS.

    Location of Go:
      Windows "C:\Go"
      Mac OS X & Linux "/usr/local/go/"

2. Set path (if path is not already set)
    ```
    Windows:
      SET PATH=%PATH%;C:\Go\bin
    macOS:
      vi ~/.bash_profile
      export PATH=$PATH:/usr/local/go/bin
      source $HOME/.bash_profile
    ```
3. Navigate to Go/src
    ```
    C:> cd Go/src

    ```
4. Create new folder
    ```
    mkdir hello

    ```
5. Save hello.go to new folder
6. Open cmd prompt OR terminal
7. Navigate to "hello" folder
    ```
    On Windows 10:
    C:> cd Go\src\hello
    C:\Go\src\hello>

    ```
8. Enter following commands to compile and then run:
    ```
    C:\Go\src\hello> go build
    C:\Go\src\hello> hello

    ```

Sample Output:
```
** Welcome to GO Bank **

Main Menu: Please select a transaction:
  Press 1 for Deposit
  Press 2 for Withdraw
  Press 3 for View Balance
  Press 0 to Exit

>> 2


*** Starting Safe Withdrawal Example ***
>> Account Details:
>> Name: John Doe
>> Balance: 500
>> Account Number: 81

Go Routine # 4 Balance: 500
Go Routine # 2 Balance: 400
Go Routine # 0 Balance: 300
Go Routine # 1 Balance: 200
Go Routine # 3 Balance: 100
New Balance: 100

Please select option:
 Press 1 For Unsafe Withdrawal Example
 Press 0 to Exit
>> 0
```

## Running the tests
Options:
Part 1
  (1) Press 1 to deposit money into the Account
      Enter in integer amount
  (2) Press 2 to see safe withdraw example
  (3) Press 3 to view Balance
  (4) Press 0 to Exit
Part 2
  (1) Press 1 to see Unsafe Withdraw example
  (2) Press 0 to Exit

### Break down into end to end tests
Part 1
  (1) will prompt for amount to deposit, display new balance, proceed to part 2
  (2) will display output for a safe withdraw, proceed to part 2
  (3) will display balance in the account, proceed to part 2
  (0) will exit the Program

Part 2
  (1) will display output for unsafe withdraw, exits
  (0) will exit the Program



```
Example Part 1:
(1) >> 1
    Enter Deposit Amount
    100
    New Balance: 600
(2) >> 2
    *** Starting Safe Withdrawal Example ***
    >> Account Details:
    >> Name: John Doe
    >> Balance: 500
    >> Account Number: 81

    Go Routine # 4 Balance: 500
    Go Routine # 2 Balance: 400
    Go Routine # 0 Balance: 300
    Go Routine # 1 Balance: 200
    Go Routine # 3 Balance: 100
    New Balance: 100
(3) >> 3
    Account Balance: 500
(0) >> 0

    C:\Go\src\hello>

Example Part 2:
(1) >> 1

  !!Warning your account will overdraw!!

  Go Routine # 1 Balance 495
  Go Routine # 16 Balance 475
  Go Routine # 2 Balance 320
  Go Routine # 3 Balance 270
  Go Routine # 4 Balance 250
  Go Routine # 5 Balance 245
  Go Routine # 6 Balance 195
  Go Routine # 7 Balance 175
  Go Routine # 8 Balance 170
  Go Routine # 9 Balance 120
  Go Routine # 10 Balance 100
  Go Routine # 11 Balance 95
  Go Routine # 12 Balance 45
  Go Routine # 13 Balance 25
  Go Routine # 14 Balance 20
  Go Routine # 15 Balance -30
  New Balance: -30
  ** Goodbye! **
(2) >> 0

    C:\Go\src\hello>
```

## Built With

* [Go]https://golang.org/dl/

## Author

* **Casey Charlesworth**
