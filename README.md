# StreakAIAssessment
This assignment contains solution of Pair-Sum problem

# About 
Here we will be given nums array/slice which contains elements and we have to return indexs of array which are summing up to given target.

Here we have to api's
1> login
2> find-pairs

In this I have added Login authentication which returns JWT Token which is needed in find-pairs api for authentications.

## Future enhancment
Connect database for LoginID and Password verification 
Create Docker image of this code (Dockerise)

## How to run this application
Step1: 
call /login api
Login Curl :
            curl --location 'http://localhost:8080/login' \
            --header 'Content-Type: application/json' \
            --data-raw '{"username": "StarkAI", "password": "StarkAI@2017"}'

Step 2:
call /find-pairs api
curl :
    curl --location 'http://localhost:8080/find-pairs' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY2NTYwNjYsInVzZXJuYW1lIjoiU3RhcmtBSSJ9.c1YFJpz95dYso7NRecSKxTpJA1Lqe1EsoxjmN1yevhQ' \
    --data '{
        "numbers": [
            1,
            2,
            3,
            4,
            5,
            5
        ],
        "target": 6
    }'