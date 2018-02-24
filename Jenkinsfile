pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        sh 'go test -v'
      }
    }
    stage('Build') {
      steps {
        sh 'go build'
      }
    }
    stage('Deploy') {
      steps {
        sh 'cp build-testing /home/budi/Desktop'
      }
    }
  }
  environment {
    GIT_BRANCH = 'master'
  }
}