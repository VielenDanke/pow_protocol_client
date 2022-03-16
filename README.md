# Proof Of Work challenge-response protocol client

### Environment variables
CLIENT_ADDRESS - address of client server;   
CLIENT_USERNAME - username is using for accessing server and validating request;   
CLIENT_PASSWORD - password is using to accessing server and validating request;   

### Restrictions
Client default address - :8090, to change it - supply header above.   
Client should be provided by credentials using headers above, if not - client won't start.   

### Interaction with server
````
curl -v -H 'X-Redirect-To:<SERVER_ADDRESS>' <CLIENT_ADDRESS>/ping
````