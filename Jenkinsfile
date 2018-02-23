pipeline {
  agent none
  stages {
    stage('Build') {
      steps {
        sh 'echo "Hello world"'
      }
    }
  }
  environment {
    GIT_BRANCH = 'master'
  }
}