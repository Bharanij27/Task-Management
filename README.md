# Task Management API

A simple task management application helps to organize and prioritize teh tasks. 

Task management API is an backend API service build using golang used to perform CRUD Operation to create a task, update it, read it and delete the tasks.


# Local Development

### Prerequisites
Before you begin, ensure you have the following installed:

- **Go (1.18 or later)**: [Download Go](https://golang.org/dl/)

- **MYSQL**: Database for storing tasks and user data. (This API uses MYSQL to store the task details)

- **cURL** or **Postman**: For testing the API.



### Steps

1. Navigate to the project directory
   ```shell
   git clone https://github.com/BharaniJ27/Task-Management-API.git
   cd Task-Management-API
   ```
2. Install all the dependencies
   ```shell
   go mod tidy
   ```
3. **Set up environment variables**:
   Create a `.env` file in the root of the project:
    ```env
    PORT=3000
    DB_PROTOCOL=tcp
    DB_USER=[DB_USER_NAME]
    DB_PASSWORD=[DB_PASSWORD]
    DB_NAME=[DB_NAME]
    DB_HOST=localhost
    DB_PORT=3306
    ```
4. Initialize the database
    ```
    CREATE DATABASE task_management;

    CREATE TABLE tasks (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        status VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON DELETE CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
    ```
5. Run the application
    ```shell
    CompileDaemon --build="go build -o main ./cmd" -command="./main"
    ```
6. If faced any issue on step 5
    ```shell
        go install github.com/githubnemo/CompileDaemon@latest   
        echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc   
        source ~/.zshrc
    ```

# Testing
### Create a task

#### Sample Payload
```json
{
    "Title": "Go Lang",
    "Description": "Learn go lang",
    "Status": "low"
}
```
#### Sample Response for task creation
```json
{
    "message": "Successfully created a task with id 1"
}
```

### Update a task

#### Sample Payload
```json
{
    "Title": "Guitar",
    "Description": "Learn go lang",
    "Status": "medium"
}
```

#### Sample Response for task updation

```json
{
    "message": "Successfully updated a task 1"
}
```

### Delete a task
##### Sample Request

```
Method: DELETE
http://localhost:3000/api/tasks/{taskID} - pass the task id the url 
```

#### Sample Response for task deletion

```json
{
    "message": "Successfully deleted a task 5"
}
```