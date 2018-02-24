pipeline {
  agent {
    node {
      label 'Test label'
    }
    
  }
  stages {
    stage('Test') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go test -v'
      }
    }
    stage('Build') {
      steps {
        sh '/home/budi/development/sdk/go/bin/go build -o build-testing'
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