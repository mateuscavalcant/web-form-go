# web-form-go
This is a web application for users to signup and login. built with Golang, HTML, CSS, Javascript and MySQL.

## Description

The User Registration Application is a web-based application that allows users to register and login to the system. It provides a secure and user-friendly interface for managing user accounts.

## Features

- Signup: Users can create an account by providing their username, email and password.
- Login: Registered users can login to access their account.
- Username validation: Checks if an username already exists.
- Email validation: Checks if an email already exists.
- Password encryption: User passwords are securely encrypted using bcrypt for enhanced security.
- Error handling: Appropriate error messages are displayed to users for invalid input or unauthorized access.

## Usage
1. Clone the repository: ```https://github.com/mateuscavalcant/web-form-go```
2. Create your file `.env`, add the database set up.
3. Run ```main.go``` to start the application.
4. Access the application in your web browser at `http://localhost:8080`
5. Access the endpoints `http://localhost:8080/signup` and `http://localhost:8080/login`
6. Login using your registered email or username and password.
7. Explore the features and manage your user account.


## Signup form
![signup](https://github.com/mateuscavalcant/web-form-go/blob/main/screenshots/signup_form.jpeg?raw=true)

## Login form
![login](https://github.com/mateuscavalcant/web-form-go/blob/main/screenshots/login_form.jpeg?raw=true)



