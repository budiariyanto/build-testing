pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go test -v'
      }
    }
    stage('Build And Deploy') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go build -o /home/budi/Desktop'
      }
    }
  }
  environment {
    GIT_BRANCH = 'master'
  }
}