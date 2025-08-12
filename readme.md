# Rumble Advertising Center Golang Challenge

This challenge is running in a simulated environment and is using a simple one file approach to understand how you think about and solve problems. The boundaries are that the results need to be accessible via curl, the code needs to compile and run within this environment, and the instructions to get the results should be clear. 

## Objectives

This challenge has two components: written and code. For the written, create a well formatted markdown document outlining your responses. Point form and succinct responses are valued.

### Written
- Review the existing code, what is the work that needs to be done if you were to take this code into production in your opinion? How does this mock up / challenge exercise differ from what you would expect to see in the 'real world'? 
- Also use this document to highlight what you have done in the code (and why)

## Code
Please complete the following (based on and using the code provided)
- Right now the project uses mock data, please externalize this so that the program can injest both a JSON file and an external url as a source of mock transactions
- Add a command line option to be able to specify the data source file name, e.g: 
```
# json file
./main --transactions transactions.json

# external url
./main --transactions https://domain.com/transactions.json
```
- Instead of displaying the PAN with GetTransactions it is preferred to only display the last four digits and replace the rest of the PAN with `*`; create a function to achieve this and ensure that all output is handled in this way
- Create an endpoint that returns the transactions ordered by descending posted_timestamp 
- Create a test for GetTransactions and your new functions