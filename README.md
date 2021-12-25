# Wallet

Wallet Coding Challange

## How To Run

First install dependencies

``go mod tidy -v``

Secondly, we need a couchbase server

``
docker run -t --name db1 -p 8091-8096:8091-8096 -p 11210-11211:11210-11211 couchbase/server:enterprise-7.0.0    
``

Then configure couchbase

``
./configure.sh
``

If this configuration file doesn't work for some reason, you can configure couchbase by yourself by 
following these steps;

```
1. Go to localhost:8091
2. Choose create new cluster
3. cluster name: wallet, username: Administrator, password: password
4. give memory as much as you want
5. then you need to add bucket in the bucket section, click add bucket and name it as wallet
```

Finally, you can run the server

`` go run cmd/main.go``

## How To Interact

Application runs on localhost:8080 and there are 4 different endpoints; 

You can view postman documentation in here https://documenter.getpostman.com/view/3718926/UVREj4bx

CREATE
```
POST /wallet
RESPONSE {id: "uuid"}
```

GET WALLET
```
GET /wallet/:id
RESPONSE {balance:amount}
```

DEPOSIT
```
POST /wallet/:id/deposit
REQUEST BODY {"amount": "float64AsString"}
RESPONSE {status: "200"}
```

WITHDRAW
```
POST /wallet/:id/withdraw
REQUEST BODY {"amount": "float64AsString"}
RESPONSE {status: "200"}
```

