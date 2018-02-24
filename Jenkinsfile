pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go test -v'
      }
    }
    stage('Build') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go build'
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