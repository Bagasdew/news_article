##Description
This repo is an app to post and get news article using cache
##How to use
- run the `init.sql` file in mysql db
- run the app
- cURL example of GET request `curl --location --request GET 'localhost:8080/article?author=tron'`
- cURL example of POST request `curl --location --request POST 'localhost:8080/article' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "author" : "John Tron",
  "title" : "Test Article ",
  "body" : "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
  }'`