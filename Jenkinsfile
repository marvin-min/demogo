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
        sh 'pwd && ls -l'  // 查看当前目录
        sh 'cd ${WORKSPACE} && go clean && GOOS=linux GOARCH=amd64 go build -o ${WORKSPACE}/hello .'  // 显式指定输出路径
        sh 'pwd && ls -l ${WORKSPACE}'  // 检查 hello 文件是否生成
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
        sh 'cp ${WORKSPACE}/hello .'  // 使用绝对路径复制文件
        sh 'docker build -t hello .'
        sh 'docker run -d -p 8080:8080 --restart=always --name demogo hello'
      }
    }
  }
}