# Attendance Service

This is a RESTful API service built with Go-Gin for managing attendance records.

## Prerequisites

- Go 
- Go-Gin 
- Mongo DB

## Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/masnasri-a/attendance-svc.git
    ```

2. Install the dependencies:

    ```shell
    go mod download
    ```

3. Set up the database:

    - Create a Mongodb database named `attendance`.
    - Update the database connection details in the `.env` file.

4. Build and run the application:

    ```shell
    go run main.go
    ```

## API Endpoints
- `/v1/auth/login` - POST request to authenticate a user.
- `/v1/auth/create-workspace` - POST request to create a new workspace.
- `/v1/auth/register` - POST request to register a new user.

For detailed information about each endpoint and the expected request/response formats, please refer to the [API documentation](./docs/api.md).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](./LICENSE).