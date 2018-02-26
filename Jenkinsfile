pipeline {
  agent any
  parameters {
        string(name: 'Git SHA1 Hash', defaultValue: 'master', description:
        'Fuck Jenkins')
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
