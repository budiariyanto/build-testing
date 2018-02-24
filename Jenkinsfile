pipeline {
  agent any
  stages {
    stage('Test') {
      parallel {
        stage('Test') {
          steps {
            sh '/home/budi/development/sdk/go/bin/go test -v'
          }
        }
        stage('Check dir') {
          steps {
            sh 'pwd'
          }
        }
      }
    }
    stage('Build') {
      steps {
        sh '''/home/budi/development/sdk/go/bin/go build
ls -la'''
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploy with ansible'
      }
    }
  }
  environment {
    GIT_BRANCH = 'master'
  }
}