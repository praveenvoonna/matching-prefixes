# Matching Prefixes

This Golang project demonstrates the efficient finding of the longest matching prefix for a given input string.

## Problem Statement

- Given a list of string prefixes of variable length, the assignment is to implement a method that takes a string as a parameter
and returns the longest prefix that matches the input string. A prefix matches if the input string starts with that prefix.
- The list of prefixes to match should be taken as a configuration by your solution. We’ve provided you with sample data that
you can use.

### Sample Data

The attached file `prefixes.txt` contain a list of prefixes, one per line, that can be used to test your implementation. Please use these in your
submission, but keep in mind that your solution should work with any list of prefixes.

### Instructions

- Your task is to implement a method with the functionality defined above in Golang. It is not sufficient for your solution to
do a linear scan of the list of prefixes . Feel free to pre-process the list of prefixes as you see fit.
- What we hope to see in your submitted assignment is an ability to follow the specification and to solve a problem
efficiently. We value readability and clear style in your code.
- Your method should be submitted within a project, not as a single source file. Tests are a great way to show that your
implementation works correctly. You may use a readme or comments to document your approach.
- Bonus points if you can additionally demonstrate the use of goroutines in the above problem.

## Installation

1. Clone the repository to your local machine.
2. Navigate to the project directory and execute the following commands in your terminal: This command compiles the project.

   ```bash
   go build
3. Run below command to execute the generated binary file.
    ```bash
    ./matching_prefixes
    
4. Alternatively, you can run the project directly without building:
    ```bash
    go run main.go

## Usage

1. Modify the prefixes.txt file to include the desired prefixes, ensuring each prefix is on a separate line.
2. Run the program and provide the input string to find the longest matching prefix.

## Testing

1. To run the test suite, execute the following command in the project directory:
    ```bash
    go test ./tests

2. This command will execute the test cases and provide you with the test results.


## Note

Make sure to replace `./matching_prefixes` with the appropriate executable file name if you are not using Windows.
