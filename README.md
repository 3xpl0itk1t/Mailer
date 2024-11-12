# Mass Mailing API with Golang and Mailgun

## Overview

This project is a Mass Mailing API built using Golang and integrated with the Mailgun service to send bulk emails efficiently. It provides an endpoint for sending mass emails to a list of recipients with customizable content. This document covers the installation, usage, and endpoint details.

---
## Prerequisites

- [Golang](https://golang.org/dl/) (version 1.18 or higher)
- [Mailgun Account](https://www.mailgun.com/) with an active API key
- Basic knowledge of REST API and HTTP methods

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd mass-mailing-api
   
2. **Install dependencies**:
    ```bash
    go mod tidy
    ```
3. **Set up environment variables** (see below).

## Environment Variables

Create a .env file in the project root directory and add the following configuration:


  `
  MAILGUN_DOMAIN=your-mailgun-domain
  `
  <br>
  `
  MAILGUN_API_KEY=your-mailgun-api-key
  `
 <br>
  `
  SENDER_EMAIL=your-sender-email@example.com
  `

## Endpoints
1. Signup <br>
Endpoint: /signup <br>
Method: POST <br>
```json
{
  "auth" : "auth code",
  "email" : "your email",
  "password" : "your password"
}
```

Example Success Response:

```json
{
  "status": "success",
  "message": "Signed Up Succesfully."
}
```

Example Error Response:

```json
{
  "status": "error",
  "message": "Description of the error."
}
```


2. Login <br>
Endpoint: /login <br>
Method: POST <br>
```json
{
  "email" : "your email",
  "password" : "your password"
}
```

Example Success Response:

```json
{
  "status": "success",
  "message": "Logged in Succesfully."
}
```

Example Error Response:

```json
{
  "status": "error",
  "message": "Description of the error."
}
```


3. Send Mass Email <br>
Endpoint: /sendmail <br>
Method: POST <br>
```json
{
  "subject": "Your Email Subject",
  "body": "Your email content here",
  "recipients": ["recipient1@example.com", "recipient2@example.com"]
}
```

Example Success Response:

```json
{
  "status": "success",
  "message": "Emails sent successfully."
}
```

Example Error Response:

```json
{
  "status": "error",
  "message": "Description of the error."
}
```
- Usage Start the server:

```bash

go run main.go
```
Send a POST request to the /sendmail endpoint with the required request body (see example below).

```bash
curl -X POST http://localhost:8080/send-mass-email \
-H "Content-Type: application/json" \
-d '{
      "subject": "Monthly Newsletter",
      "body": "<h1>Hello, Subscribers!</h1><p>This is our latest newsletter.</p>",
      "recipients": ["user1@example.com", "user2@example.com"]
    }'
```
Example Success Response:

```json
{
  "status": "success",
  "message": "Emails sent successfully."
}
```

Example Error Response:

```json
{
  "status": "error",
  "message": "Invalid recipient email address."
}
```

## Additional Notes
This API is set up to send emails asynchronously to avoid blocking the main thread.
Rate limits on Mailgun may affect the delivery speed if the recipient list is too large.
Ensure SENDER_EMAIL is verified in Mailgun for successful delivery.


## License
This project is licensed under the MIT License. See the LICENSE file for details.


`https://github.com/3xpl0itk1t/Mailer`
