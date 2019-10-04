export AAA=$(git rev-parse HEAD)
echo $AAA
#openssl aes-256-cbc -K $encrypted_0c35eebf403c_key -iv $encrypted_0c35eebf403c_iv -in service-account.json.enc -out service-account.json -d
#curl https://sdk.cloud.google.com | bash > /dev/null;
#source $HOME/google-cloud-sdk/path.bash.inc
#gcloud components update kubectl
#gcloud auth activate-service-account --key-file service-account.json
#gcloud config set project learn-drive-253814
#gcloud config set compute/zone asia-southeast1-a
#gcloud container clusters get-credentials standard-cluster-1
docker build -t mikedutuandu/micro_teacher_test -f ./micro_teacher/Dockerfile.dev ./micro_teacher