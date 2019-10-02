# Code base micro services by golang

- gRpc,protobuf for communicate
- Gateway(2 gateway) to forward http to gRpc
- Pub/sub for event driven source
- Mongo for db(3 db)
- Discovery services from kubenetes
- Use docker, docker compose for dev
- Use docker, Kubenetes for production
- Use travis ci for CI,CD

# Note for development when run docker compose

    1. For the first time setup this project, must go to each folder and run command below for setup vendor inside this
        go mod vendor
    
    2. When change file docker-compose.yml must run these command below for affect
        docker-compose down
        docker-compose up -d
    
    3. When just change the codebase must run command below for affect
        docker-compose restart

# Note for development when run each micro by go run
    
    1. Use tool below for call api gRPC
        https://github.com/ktr0731/evans
        
    2. Some command from evans
        evans -p 3001 -r
        show package
        show service
        call [name service]
    
    