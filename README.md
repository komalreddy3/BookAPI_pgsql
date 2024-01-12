## Command for REST

``` API_TYPE=rest DB_NAME=book_keeper USER_NAME=postgres PASSWORD='abc@123' go run main.go ```


``` router.HandleFunc("/api/books", createBook).Methods("POST") ```

![1](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/959a3273-d58c-4ff9-a9fe-96e2b6c3b866)

![2](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/b70ef5a1-aec9-4f63-9192-29402c653805)

![3](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/d6cf5570-89ed-4685-b689-69bc2131a869)


``` router.HandleFunc("/api/books", createBook).Methods("POST") ```

![4](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/97164f90-61d8-45ea-ae8b-2ea0cc30d94b)

![5](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/99df1888-8b13-408c-a04e-1a21025b5233)

![6](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/be8414cd-36c0-4a8f-b155-4675d7d11672)

``` router.HandleFunc("/api/books", getBooks).Methods("GET") ```

![7](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/375cfe07-5963-4b6c-8c24-fbe1a1c55537)


![8](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/6fc1fcb6-011f-46a6-bf09-552a47d5c75f)

``` router.HandleFunc("/api/books/{id}", getBook).Methods("GET") ```


![9](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/e04bba2d-3de6-4c42-97d1-2ceedf213e3d)

``` router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT") ```

![10](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/47c9346f-793e-4c34-9f1e-00686c1f2da2)



![11](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/6c52789a-c9ad-4622-b654-2eeebff13b18)

![12](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/c0e06b81-3a19-4346-b207-f0a9a723d2e2)

``` router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") ```

![13](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/e5b8c61f-e8e8-4c05-bddf-63739f880bc4)

![14](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/c8d6b3c1-8485-4a9a-a50c-0fe91972859e)

## Command for gRPC


``` API_TYPE=grpc DB_NAME=book_keeper USER_NAME=postgres PASSWORD='abc@123' go run main.go ```




![16](https://github.com/komalreddy3/BookAPI_pgsql/assets/82363361/0f6a2464-f52a-4a22-a7bc-b071f4464cd3)
