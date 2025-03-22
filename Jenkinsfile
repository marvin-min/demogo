pipeline {
  agent any
  environment {
    WS = "${WORKSPACE}"
  }
  stages {
    stage('环境检查') {
      steps {
        sh 'printenv'
        sh 'hostname'
      }
    }

    stage('Test') {
      agent {
        docker { image 'golang:1.23.7-alpine3.21' }
      }
      steps {
        echo 'Testing..'
        sh 'go test ./...' 
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
        sh 'ls -l'  // 查看当前目录
        sh 'go clean'
        sh 'cd $WS && GOOS=linux GOARCH=amd64 go build -o hello .'  // 直接输出到当前目录
        sh 'ls -l'  // 检查 hello 是否生成
      }
    }

    stage('Deploy') {
  
      steps {
        sh 'docker version'
        echo 'Deploying....'
        sh 'ls -l'
        sh 'cd $WS && docker build -t hello .'  // 直接在 Jenkins 工作目录下构建
        sh 'docker stop demogo || true'  // 如果容器已运行，则先停止
        sh 'docker rm demogo || true'  // 删除旧容器
        sh 'docker run -d -p 8888:8080 --restart=always --name demogo hello'
      }
    }
  }
}
