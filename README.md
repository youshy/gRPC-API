# gRPC-API

how to make an API with gRPC

## Usage

* Make sure you have `Go` installed: `go version`
* Setup protobuf
  * On MAC: `brew install protobuf`
  * On Linux: 
    ```
    # Make sure you grab the latest version
    curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
    # Unzip
    unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
    # Move protoc to /usr/local/bin/
    sudo mv protoc3/bin/* /usr/local/bin/
    # Move protoc3/include to /usr/local/include/
    sudo mv protoc3/include/* /usr/local/include/
    # Optional: change owner
    sudo chown [user] /usr/local/bin/protoc
    sudo chown -R [user] /usr/local/include/google
    ```
    (or run `./proto.sh` and change your user)
