# GoBlastApi

A sample implementation of how to use the GoBlast framework. This example showcases a simple arithmetic addition endpoint:
```cmd
curl --location 'http://localhost:1323/api/arithmetic/add' \
--header 'X-WellKnown-GoBlast-Reference-Id: 126NtJbDfM' \
--header 'Content-Type: application/json' \
--data '{
    "variableOne": 19.0,
    "variableTwo": 34.9
}'
```