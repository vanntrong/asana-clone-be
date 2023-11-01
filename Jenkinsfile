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
                    app = docker.build("asana-be")
                }
            }
        }
        stage('Test') {
            steps {
                 echo 'Empty'
            }
        }
        stage('Push image to ECR') {
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
    post {
        // Clean after build
        always {
            cleanWs()
        }
    }
}