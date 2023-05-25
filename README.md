<h1 align="center">Chain Reorganisation Tracker</h1>

## About The Project

Check for chain reorganisation on ethereum mainet and logs them.

## Test Logs & expected output

 ```
Block number: 17319058
Block hash: 0xdc1f81ad9a0c1014baf50d9c964720779570f178d752c4c0cb59b5154564baad
Block number: 17319059
Block hash: 0x475c69ddfa416acb5393342b03b0b2898035999a9241ea9091f9f94de5d91a45
Block number: 17319060
Block hash: 0x5a2bc136e0cb5fd348c72e91ffce2656428756273afd1e4f306589d1f7f19fbc
Chain reorganization detected!
Discarded blocks:
Discarded Block Number: 17319060, Discarded Block Hash: 0xa7d6927deb3a1d37599c691c5df12bf344ce1232e98ac1fd5650dab770350a4b
Block number: 17319060
Block hash: 0xa7d6927deb3a1d37599c691c5df12bf344ce1232e98ac1fd5650dab770350a4b
Block number: 17319061
Block hash: 0x82394dfd7eb8bed314d7d0b47ee85029a5c2afc7bac57927e35de699ee20cd0e
 ```

## Prerequisites

Download and install [Golang 1.20](https://go.dev/doc/install) (or higher).  

## How To Use?

1. Navigate to chainReorg/:
   ``` 
   cd /path/to/folder/sarva-assignment/chainReorg/
   ``` 
2. Get dependencies:
   ``` 
   go mod tidy
   ```
3. Create .env from .env.example and update URL with Infrua Websocket
    ```
   URL = <INFURA-WEBSOCKET>
    ```
4. Run the app:
   ``` 
   go run chainReorg.go 
   ```
