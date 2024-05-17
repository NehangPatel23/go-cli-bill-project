# CLI Bill Project

A simple command-line application written in Go for managing bills. This application allows users to create a new bill, add items, add tips, and save the bill to a file.

## Features

- Create a new bill
- Add items to the bill
- Add tips to the bill
- Save the bill to a text file

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.15 or later)
- A terminal or command prompt

## Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/NehangPatel23/cli-bill-project.git
    cd cli-bill-project
    ```

2. **Create the Bills Directory**:
    ```sh
    mkdir bills
    ```

## Usage

1. **Run the Application**:
    ```sh
    go run main.go bill.go
    ```

2. **Follow the Prompts**:
    - Enter a name for the new bill.
    - Choose options to add items, add tips, or save the bill.

### Example

```zsh
nehangpatel@Nehangs-MacBook-Pro CLI Bill Project % go run main.go bill.go                                                     
Create a new bill name: Nehang's Bill
Created the bill -  Nehang's Bill
Choose an option (a - Add an item, s - Save the bill, t - Add a tip): a
Item Name: Pizza
Item Price: 9.75
Item Added -  Pizza 9.75
Choose an option (a - Add an item, s - Save the bill, t - Add a tip): a
Item Name: Pepsi
Item Price: 3.99
Item Added -  Pepsi 3.99
Choose an option (a - Add an item, s - Save the bill, t - Add a tip): a
Item Name: Curly Fries
Item Price: 4.50
Item Added -  Curly Fries 4.5
Choose an option (a - Add an item, s - Save the bill, t - Add a tip): t
Enter tip amount: $2.50
Tip Added -  2.5
Choose an option (a - Add an item, s - Save the bill, t - Add a tip): s
The bill was saved successfully!
You chose to save the bill to the file -  Nehang's Bill
```

The output bill file would look like this: [Sample Output]()