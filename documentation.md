
# Bookstore API

This project is a simple **Bookstore API** that provides functionality for managing books, authors, customers, orders, and book sales in an e-commerce setting. It is designed using Go, with a focus on an in-memory data store, and follows a clean architecture with separation of concerns.

## Features

The API allows you to perform CRUD operations on various resources like books, authors, customers, orders, and book sales.

- **Books**: Add, retrieve, update, and delete books from the store.
- **Authors**: Add, retrieve, update, and delete authors.
- **Customers**: Manage customer information (CRUD operations).
- **Orders**: Create, retrieve, update, and delete orders.
- **Book Sales**: Record and search for book sales.

### API Endpoints

The following sections describe the API endpoints, based on the Swagger documentation.

#### Books

- **POST /books**: Create a new book.
- **GET /books/{id}**: Retrieve a book by its ID.
- **PUT /books/{id}**: Update a book by its ID.
- **DELETE /books/{id}**: Delete a book by its ID.
- **GET /books**: Search for books by filters.all books are are returned if not filters are provided with the json request 


#### Authors

- **POST /authors**: Create a new author.
- **GET /authors/{id}**: Retrieve an author by ID.
- **PUT /authors/{id}**: Update an author by ID.
- **DELETE /authors/{id}**: Delete an author by ID.
- **GET /authors**: Search for authors. all customers are are returned if not filters are provided with the json request 
 

#### Customers

- **POST /customers**: Create a new customer.
- **GET /customers/{id}**: Retrieve a customer by ID.
- **PUT /customers/{id}**: Update a customer by ID.
- **DELETE /customers/{id}**: Delete a customer by ID.
- **GET /customers**: get all customers.

#### Orders

- **POST /orders**: Create a new order.
- **GET /orders/{id}**: Retrieve an order by ID.
- **PUT /orders/{id}**: Update an order by ID.
- **DELETE /orders/{id}**: Delete an order by ID.
- **GET /orders**:get all orders

#### Book Sales

- **POST /bookSales**: Create a new book sale.
- **GET /bookSales/{id}**: Retrieve a book sale by ID.
- **PUT /bookSales/{id}**: Update a book sale by ID.
- **DELETE /bookSales/{id}**: Delete a book sale by ID.
- **GET /bookSales**: Search for book sales, all sales are are returned if not filters are provided with the json request 

## Project Structure

The project is structured as follows:

```
/bookstore
  /handlers        # HTTP handlers for handling API requests
  /memory          # In-memory store for handling the data
  /models          # Data models representing the entities
  /repositories    # Interfaces for interacting with the data store
  /services        # Business logic layer for handling CRUD operations
  openapi.yml         # Swagger configuration 
  main.go          # Entry point to run the application
```

### Directories and Files Breakdown

- **/handlers**: Contains the HTTP handler functions which process incoming requests, map them to the appropriate service methods, and return responses.
- **/memory**: Implements the in-memory data store using Go maps and sync mechanisms (mutexes). Singleton instances are used for managing resources like books, customers, and orders.
- **/models**: Defines the data models that represent entities such as books, authors, orders, and book sales.
- **/repositories**: Contains interfaces for data access layers, such as methods for creating, retrieving, and deleting entities from the data store.
- **/services**: Handles the business logic and interacts with the repositories for CRUD operations and data management.
- **openapi.yml** :Swagger configuration 

- **main.go**: The main entry point for the application, where the server is set up, and routing is initialized.




## Example Requests
refer to the swagger file, to explore different apis and there examples.

## Technologies Used

- **Go**: The programming language used for building the API.
- **Swagger/OpenAPI**: API documentation and testing interface.
- **In-Memory Store**: Simple in-memory storage for entities.
- **Goroutines & Mutexes**: For managing concurrency and thread safety.

## Contributing

Feel free to fork the repository, open issues, and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
