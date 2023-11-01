pipeline {
    agent any
    options {
        skipStagesAfterUnstable()
    }
    environment {
        GOOGLE_CLIENT_ID = credentials('google_client_id')
    }
    stages {
        stage("Add env") {
            steps {
                sh 'echo "GOOGLE_CLIENT_ID=${MYSQL_ROOT_LOGIN}" | tee -a .app.production.env'
            }
        }

        stage('Build') { 
            steps { 
                script {
                    app = docker.build("underwater")
                }
            }
        }
        stage('Test') {
            steps {
                 echo 'Empty'
            }
        }
        stage('Deploy') {
            steps {
                script{
                    docker.withRegistry('https://475211048958.dkr.ecr.ap-southeast-1.amazonaws.com/asana-be', 'ecr:ap-southeast-1:aws-credentials') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }
    }
}