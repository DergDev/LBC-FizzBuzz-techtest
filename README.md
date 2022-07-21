# LBC-FizzBuzz-techtest
Technical Test for Le Bon Coin - Backend Dev in Golang

Julien Guillochon's (https://www.linkedin.com/in/julien-guillochon/) Attempt to solve the Fizz-Buzz Technical Test for Le Bon Coin.
This is a project created from the ground up using :
- Project Templating: https://github.com/golang-standards/project-layout
- ORM:                https://github.com/go-gorm/gorm
- Router:             https://github.com/gorilla/mux
------------------------------------------------------------------------------------
 
Prompt:
 
  "Exercise: Write a simple fizz-buzz REST server. 
  The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".
  
  Your goal is to implement a web server that will expose a REST API endpoint that: 
  - Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
  - Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
  - The server needs to be:
    - Ready for production
    - Easy to maintain by other developers
  
  - Add a statistics endpoint allowing users to know what the most frequent request has been. 
    This endpoint should:
      - Accept no parameter
      - Return the parameters corresponding to the most used request, as well as the number of hits for this request"

------------------------------------------------------------------------------------

To run this project you will need:
- docker :  https://www.docker.com/
------------------------------------------------------------------------------------

To test This project you can use this postman collection:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.getpostman.com/collections/587033dd476bdfa406bf)

------------------------------------------------------------------------------------

To run the project :
Start the dev docker environment on your workspace
  
  Use the docker command: 
  ```docker-compose --file docker-compose.dev.yml up```

Then run the different postman commands to test the exposed endpoints

Once you're done use the command:
``` docker-compose -f docker-compose.dev.yml stop ```

OR

Use the make command to start the dev / prod environmement: 
  ```make dev``` | ```make prod```
  
Then run the different postman commands to test the exposed endpoints
  
Finally use the make command to stop the dev / prod environmement: 
  ```make stop_dev``` | ```make stop_prod```

------------------------------------------------------------------------------------

Thanks for your time!
