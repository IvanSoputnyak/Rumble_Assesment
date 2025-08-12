Go To prod:
- Env vars, it's not good to keep sensible project data in your code
- do we need to encrypt data in any way as financial operations is involved?
- Tests should be both unit and E2E tests
- logs/monitoring good to have also not just general errors/warnings logs but also audit logs for sensible operations.
- CI/CD flow should be described same to dockerfile.
- project structure update no all in one file
- I would use stream to pass json files not getting from localhost. Blob storages can stream data out of the box.

What was done in code:
- put code to seperate files by it's purpose(easier to read follow and update)
- add mask function(we need this functionality)
- add seperate type for return masked PAN(it's better to keep input and output datatypes different datastructure even if it's looks the same)
- add functions for load file(should be a separate function for all types of getting files, seperation of concerns)
- add handler for sorted transactions by timestamp(requirenment)
- implement sort function(also better to have it seperate as it's easier to modify/reuse)
- add tests for handlers(need a prove that code works)

What should be done
- For TransactionMasked redifined there should be better way to do it?
- what is postedTimestamp??
- tests just cover basic scenarios
- 