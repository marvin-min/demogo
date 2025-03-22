pipeline {
  // 任何一个代理可用就可以执行（集群模式）
  agent any
  //定义一些环境信息
  //定义流水线的加工流程
  stages {
    stage('环境检查') {
      steps {
        sh 'printenv'
        sh  'hostname'
        }
    }
       // 2. 单元测试
    stage('Test') {
      steps {
        echo 'Testing..'
      }
    }
    // 1. 拉取代码
    // 3. 编译
    stage('Build') {
      agent {
          docker {image 'golang:1.23.7-alpine3.21'}
      }
      steps {
        sh 'go version'
        sh 'go env'
        echo 'Building..'
        sh 'pwd & ls -l'
        sh 'cd ${WORKSPACE} & go clean & GOOS=linux GOARCH=amd64 go build -o hello .'
        sh 'pwd & ls -l'
      }
    }
    // 4. 部署
    stage('Deploy') {
      steps {
        sh 'docker version'
        echo 'Deploying....'
        sh 'pwd & ls -l'
        sh 'pwd & docker build -t hello .'
        sh 'docker run -d -p 8080:8080   --restart=always --name demogo demogo'
      }
    }
  }
}