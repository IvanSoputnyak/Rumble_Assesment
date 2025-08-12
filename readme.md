# How to run the code

## Requirenments
- Go 1.20+ 

## How To Run
- For local file, put your JSON file into the roort project folder and run the following code, where 'transactions.json' is your filename:
```
make run TRANS=./transactions.json
```
- For URL should be something like this, you can utilize this Gist or your URL:
```
make run TRANS="https://gist.githubusercontent.com/IvanSoputnyak/a1bd89db58de230169e643b74b7d77b8/raw/d5ede34930ea8e10299777e4bb41133584fd5a7c/transactions.json"
```

- You should be able to see something like this is everything is good:
```
2025/08/12 11:20:03 listening on :8000
```

- Try to curl the data:
```
curl http://localhost:8000/transactions | jq .
curl http://localhost:8000/transactions/ordered | jq .
```

- If you want to run tests, simply do this:
```
make test
```

# Written part

## Go To prod:
- Env vars, it's not good to keep sensible project data in your code
- do we need to encrypt data in any way as financial operations is involved?
- Tests should be both unit and E2E tests
- logs/monitoring good to have also not just general errors/warnings logs but also audit logs for sensible operations.
- CI/CD flow should be described same to dockerfile, linters etc.
- project structure update no all in one file
- I would use stream to pass json files not getting from localhost. Blob storages can stream data out of the box.
- need to add server timeouts
- security(tokens)


## What was done in code:
- put code to seperate files by it's purpose(easier to read follow and update)
- add mask function(we need this functionality)
- add seperate type for return masked PAN(it's better to keep input and output datatypes different datastructure even if it's looks the same)
- add functions for load file(should be a separate function for all types of getting files, seperation of concerns)
- add handler for sorted transactions by timestamp(requirenment)
- implement sort function(also better to have it seperate as it's easier to modify/reuse)
- add tests for handlers(need a prove that code works)
- added download file by chunks(we do not know the input file size. So based on that we can assume it can be any size which means we can download a huge file in RAM which can cause inefficent resource usage/potential fails)

## What should be done
- For TransactionMasked redifined there should be better way to do it?
- what is postedTimestamp??
- tests just cover basic scenarios
- Server termination is not implemented
- Need to add server timeouts
- inside the system PostedTimeStamp should be time.Time
- it's not good to return the whole array of transactions as being said we do not know file size.
- based on the above statement we need pagination, filters etc
- error handling, retries is not there
- invalid input data edge cases is not covered

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