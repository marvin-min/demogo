pipeline {
  agent any

  stages {
    stage('环境检查') {
      steps {
        sh 'printenv'
        sh 'hostname'
      }
    }

    stage('Test') {
      steps {
        echo 'Testing..'
      }
    }

    stage('Build') {
      agent {
        docker { image 'golang:1.23.7-alpine3.21' }
      }
      steps {
        sh 'go version'
        sh 'go env'
        echo 'Building..'
        sh 'pwd && ls -l'
        sh 'go clean && GOOS=linux GOARCH=amd64 go build -o hello .'
        sh 'pwd && ls -l'
      }
    }

    stage('Deploy') {
      environment {
        DOCKER_BUILDKIT = 1
      }
      steps {
        sh 'docker version'
        echo 'Deploying....'
        sh 'pwd && ls -l'
        sh 'cp ${WORKSPACE}/hello .'  // 将 hello 文件复制到当前目录
        sh 'docker build -t hello .'
        sh 'docker run -d -p 8080:8080 --restart=always --name demogo hello'
      }
    }
  }
}