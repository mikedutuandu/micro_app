pipeline {
    agent any

    environment {
        SHA=sh 'git rev-parse HEAD'
        CLOUDSDK_CORE_DISABLE_PROMPTS=1
        CLOUDSDK="test"
    }

    stages {
        stage('Cloning Git') {
          steps {
            git 'https://github.com/mikedutuandu/micro_app.git'
          }
        }
        stage('Build') {
            steps {
                echo 'Building ${CLOUDSDK}'
                echo "Database engine is ${SHA}"
            }
        }
        stage('Test') {
            steps {
                echo 'Testing'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying'
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