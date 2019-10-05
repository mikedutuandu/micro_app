# Code base micro services by golang

- gRpc,protobuf for communicate
- Gateway(2 gateway) to forward http to gRpc
- Pub/sub for event driven source
- Mongo for db(3 db)
- Discovery services from kubenetes
- Use docker, docker compose for dev
- Use docker, Kubenetes for production
- Use travis ci or jenkins for CI,CD

# A. Note for development when run docker compose

    1. For the first time setup this project, must go to each folder and run command below for setup vendor inside this
        go mod vendor
    
    2. When change file docker-compose.yml must run these command below for affect
        docker-compose down
        docker-compose up -d
    
    3. When just change the codebase must run command below for affect
        docker-compose restart

# B. Note for development when run each micro by go run
    
    1. Use tool below for call api gRPC
        https://github.com/ktr0731/evans
        
    2. Some command from evans
        evans -p 3001 -r
        show package
        show service
        call [name service]
    
# C. Note for build CI & CD with jenkins
    1. Build and run jenkins with docker
    docker build -t jenkins_micro_app -f Dockerfile.jenkins .
    docker run -d -p 8081:8080 -v /var/run/docker.sock:/var/run/docker.sock -v /Users/mikedu/Jenkins:/var/jenkins_home jenkins_micro_app
    
    2. Set ssh key for connect to git lap
    - Log in to container jenkins and generate ssh key
    - Add ssh key to git lap
    - On container run to verify and confirm connect to git lap: ssh -T git@gitlab.com
    
    3. Install EnvInject plugin(for set env to inject to pipeline) and Gitlap plugin for jenkins
    
    4. Create pipeline one jenkins
    - New Item > Pipeline
    - Build Triggers > Poll SCM: * * * * *
    - Set env for connect to docker hub
        DOCKER_USERNAME="[your username]"
        DOCKER_PASSWORD="[your password]"
    - Pipeline 
        + Config repo point to gitlap
        + Config point to Jenkinsfile of repo to run pipeline
   