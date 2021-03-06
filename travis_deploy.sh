docker build -t mikedutuandu/micro_teacher:latest -t mikedutuandu/micro_teacher:$SHA -f ./micro_teacher/Dockerfile ./micro_teacher
docker build -t mikedutuandu/micro_learner:latest -t mikedutuandu/micro_learner:$SHA -f ./micro_learner/Dockerfile ./micro_learner
docker build -t mikedutuandu/micro_booking:latest -t mikedutuandu/micro_booking:$SHA -f ./micro_booking/Dockerfile ./micro_booking
docker build -t mikedutuandu/gateway_teacher:latest -t mikedutuandu/gateway_teacher:$SHA -f ./gateway_teacher/Dockerfile ./gateway_teacher
docker build -t mikedutuandu/gateway_learner:latest -t mikedutuandu/gateway_learner:$SHA -f ./gateway_learner/Dockerfile ./gateway_learner

docker push mikedutuandu/micro_teacher:latest
docker push mikedutuandu/micro_learner:latest
docker push mikedutuandu/micro_booking:latest
docker push mikedutuandu/gateway_teacher:latest
docker push mikedutuandu/gateway_learner:latest

docker push mikedutuandu/micro_teacher:$SHA
docker push mikedutuandu/micro_learner:$SHA
docker push mikedutuandu/micro_booking:$SHA
docker push mikedutuandu/gateway_teacher:$SHA
docker push mikedutuandu/gateway_learner:$SHA

kubectl apply -f config_kubernetes
kubectl set image deployments/micro-teacher-deployment micro-teacher=mikedutuandu/micro_teacher:$SHA
kubectl set image deployments/micro-learner-deployment micro-learner=mikedutuandu/micro_learner:$SHA
kubectl set image deployments/micro-booking-deployment micro-booking=mikedutuandu/micro_booking:$SHA
kubectl set image deployments/gateway-teacher-deployment gateway-teacher=mikedutuandu/gateway_teacher:$SHA
kubectl set image deployments/gateway-learner-deployment gateway-learner=mikedutuandu/gateway_learner:$SHA
