pipeline {
    environment {
        VERSION = "v1.0.2"
        PROJECT_NAME = JOB_NAME.split('/')
        IMAGE_NAME = "unjxde/${PROJECT_NAME[0]}"
        IMAGE = ''
    }
    agent any
    stages {
        stage('Building') {
            steps {
                script {
                    IMAGE = docker.build IMAGE_NAME
                }
            }
        }
        stage('Deploying') {
            steps {
                script {
                    docker.withRegistry( 'https://registry.hub.docker.com', 'dockerHub' ) {
                        IMAGE.push("${VERSION}")
                        if (BRANCH_NAME == "main") {
                            IMAGE.push("latest")
                        }
                    }
                }
            }
        }
    }
}
