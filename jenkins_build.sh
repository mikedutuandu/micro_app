gcloud auth activate-service-account --key-file service-account.json
gcloud config set project learn-drive-253814
gcloud config set compute/zone asia-southeast1-a
gcloud container clusters get-credentials standard-cluster-1
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t mikedutuandu/micro_teacher_test -f ./micro_teacher/Dockerfile.dev ./micro_teacher
docker build -t mikedutuandu/micro_learner_test -f ./micro_learner/Dockerfile.dev ./micro_learner
docker build -t mikedutuandu/micro_booking_test -f ./micro_booking/Dockerfile.dev ./micro_booking
docker build -t mikedutuandu/gateway_teacher_test -f ./gateway_teacher/Dockerfile.dev ./gateway_teacher
docker build -t mikedutuandu/gateway_learner_test -f ./gateway_learner/Dockerfile.dev ./gateway_learner