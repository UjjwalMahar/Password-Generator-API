# Password Generator API

This is a Password Generator built using GoLang ,it generates password according to the user input.

## API Endpoints

### Get Root

- **URL**: `/`
- **Method**: GET
- **Description**: This endpoint returns a friendly message to let you know that the Password Generator API was created with love.

#### Example

```json
{
    "message": "Password Generator Created with ❤️ by Ujjwal Mahar"
}
```

### Generate Password

- **URL**: `/generate-password`
- **Method**: POST
- **Description**: This endpoint generates a random password based on the user's input.

#### Request Body

- `Length` (int): The desired length of the password.
- `ReqSpecialChar` (bool): A flag indicating whether special characters should be included in the password.
- `ReqDigit` (bool): A flag indicating whether digits (numbers) should be included in the password.

#### Example

```http
POST /generate-password HTTP/1.1
Content-Type: application/json

{
    "Length": 12,
    "ReqSpecialChar": true,
    "ReqDigit": true
}
```

#### Response

If the password is generated successfully, you will receive a response like this:

```json
{
    "GeneratedPassword": "5g!vK#p7Zf@2"
}
```

If there is an issue with generating the password, you might receive an error response:

```json
{
    "error": "Failed to generate a password"
}
```

## How to Run

To run this application, you need to have Go installed. Here are the steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/UjjwalMahar/Password-Gen.git
   ```

2. Change to the project directory:

   ```sh
   cd Password-Gen
   ```

3. Build and run the application:

   ```sh
   go run main.go
   ```

The server will start, and you can access the API endpoints as described above.


## License

This project is open-source and available under the [MIT License](LICENSE). Feel free to use and modify it as needed.