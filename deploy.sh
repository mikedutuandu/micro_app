docker build -t mikedutuandu/multi-client:latest -t mikedutuandu/multi-client:$SHA -f ./client/Dockerfile ./client
docker build -t mikedutuandu/multi-server:latest -t mikedutuandu/multi-server:$SHA -f ./server/Dockerfile ./server
docker build -t mikedutuandu/multi-worker:latest -t mikedutuandu/multi-worker:$SHA -f ./worker/Dockerfile ./worker

docker push mikedutuandu/multi-client:latest
docker push mikedutuandu/multi-server:latest
docker push mikedutuandu/multi-worker:latest

docker push mikedutuandu/multi-client:$SHA
docker push mikedutuandu/multi-server:$SHA
docker push mikedutuandu/multi-worker:$SHA

kubectl apply -f k8s
kubectl set image deployments/server-deployment server=mikedutuandu/multi-server:$SHA
kubectl set image deployments/client-deployment client=mikedutuandu/multi-client:$SHA
kubectl set image deployments/worker-deployment worker=mikedutuandu/multi-worker:$SHA