pipeline {
  // 任何一个代理可用就可以执行（集群模式）
  agent any
  //定义一些环境信息
  //定义流水线的加工流程
  stages {
    // 1. 拉取代码
    // 3. 编译
    stage('Build') {
      steps {
        echo 'Building..'
        sh 'pwd & ls -l'
      }
    }
    // 2. 单元测试
    stage('Test') {
      steps {
        echo 'Testing..'
      }
    }
    // 4. 部署
    stage('Deploy') {
      steps {
        echo 'Deploying....'
        sh 'go version'
      }
    }
  }
}