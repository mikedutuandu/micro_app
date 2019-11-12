pipeline {
    agent any

    environment {
        CLOUDSDK_CORE_DISABLE_PROMPTS=1
    }

    stages {
        stage('Cloning Git') {
          steps {
            git 'git@gitlab.com:mikedutuandu/micro_app.git'
          }
        }
        stage('Build') {
            steps {
                echo "Building ==============================================="
                //sh 'bash ./jenkins_build.sh'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing ================================================='
                //sh 'bash ./jenkins_test.sh'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying ================================================'
                //sh 'bash ./jenkins_deploy.sh'
            }
        }
    }

    post {
            always {
                echo 'This will always run'
            }
            success {
                echo 'This will run only if successful'
            }
            failure {
                echo 'This will run only if failed'
            }
            unstable {
                echo 'This will run only if the run was marked as unstable'
            }
            changed {
                echo 'This will run only if the state of the Pipeline has changed'
                echo 'For example, if the Pipeline was previously failing but is now successful'
            }
        }
}