pipeline {
  agent any
  parameters {
        string(name: 'GIT_REVISION', defaultValue: 'master', description:
        'Fuck Jenkins')
    }
  stages {
    state('Clone') {
      steps {
        git clone git@github.com:budiariyanto/build-testing.git
        git checkout $(params.GIT_REVISION)
      }
    }
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
