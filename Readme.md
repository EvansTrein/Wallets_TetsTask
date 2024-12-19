Language: [EN](https://github.com/EvansTrein/Wallets_TetsTask/blob/main/Readme.md), [RU](https://github.com/EvansTrein/Wallets_TetsTask/blob/main/RU.md)

# A test assignment from one of the companies
The task looks like this:

**Write an application that accepts a request of the following type via REST**

POST api/v1/wallet

{<br>
walletId: UUID, <br>
operationType: DEPOSIT or WITHDRAW, <br>
amount: 1000 <br>
}

after executing the logic to change the account in the database

**it is also possible to get the wallet balance**
GET api/v1/wallets/{WALLET_UUID}

stack:
- Golang
- Postgresql
- Docker

Pay special attention to problems when working in a competitive environment (1000 RPS per wallet). No request should be unprocessed (50X error)

the application should run in a docker container, the database too, the whole system should be brought up using docker-compose.

It is necessary to cover the application with tests
Upload the solved task to github, provide a link

Environment variables should be read from the config.env file.

All arising questions on the assignment should be solved independently, at your own discretion.


# How do I start it up? 
By default, run from Docker (to make the database). Clone the repository --> run the command `docker compose --env-file config.env up --build -d`. 

**Important!** If the server container does not connect to the database after the build, just restart it. As I realized, this is a common problem, but I haven't figured out how to fix it yet

To run without Docker -->  you need to remove comments from code lines in <u>envs.go</u> file (it is marked where exactly) and change POSTGRES_HOST to localhost in <u>config.env</u>.

After launching, to view the documentation, go to http://localhost:8000/swagger/index.html


# What didn't I know from this?
**this part:** "Pay special attention to problems when working in a competitive environment (1000 RPS per wallet). No request should be unprocessed (50X error)"

As I understand it, we're talking about transactions here - dirty reads, lost writes, and phantoms.

Before that, I read about ACID, but I have not reflected it in the code yet. And I haven't worked with SQL much. 

**and more:** "It is necessary to cover the application with tests"

Before, tested the http server myself, through Postman, not through tests in Go.

# What were you thinking of doing? 
I decided to do it through the GIN framework. As for the database and SQL, the assignment was sent just when I was studying this topic. I had a book on Postgresql from the official documentation site 

From myself, I added a swagger and a query to create a wallet.

# Fulfillment
**POST api/v1/wallet**
- there was an understanding of how to write a handler, judging by the task, the necessary data comes in the body of the request.
- The whole difficulty for me was in the interaction with the database. I solved it through `FOR UPDATE` + transaction with Read Committed isolation level (default for Postgres). I thought about Serializable isolation, but I didn't, the task doesn't say that it's necessary to be so strict. **Summary** - a transaction is in progress and if a GET request arrives at that moment, it will return the balance that was BEFORE the transaction started. If another transaction request comes, it will queue up and wait until the transaction before it is completed. 
- How did I test it? I added `time.Sleep(time.Second * 10)` to the transaction and via Postman I ran 5 requests for the transaction and 2 more to get the balance. POST requests worked one by one, and GET requests showed the last known balance (no transactions were used in GET request).

**GET api/v1/wallets/{WALLET_UUID}**
- Here the necessary data is passed in the query itself, we extract and search in the database. Transactions are not used here.  

**No request should be unprocessed (50X error)** - I understood it this way: the server should not return 500 codes, i.e. each SQL query should work, following the logic - the balance allows, debit; the balance does not have the required amount for the operation, return 400 server response.. 

**The application must run in a docker container** - there is a package in Docker. Two containers, a database and a server. The server has a multi-stage build (to reduce the image size).

**It is necessary to cover the application with tests** - There is a package of tests. The first time I made tests for the http server. It was the tests that took more time than anything else. **Important!** For the tests to be correct, the server must be started at least once, so that the migration takes place and the table is created. A separate table is needed for testing, but here everything is in one table.

**Environment variables must be read from the config.env file** - there is a config.env file.
