pipeline {
  agent none
  stages {
    stage('Build') {
      agent any
      steps {
        dir(path: '/home/budi/development/build') {
          git(url: 'git@github.com:budiariyanto/build-testing', branch: '$GIT_BRANCH', changelog: true)
        }
        
      }
    }
  }
  environment {
    GIT_BRANCH = 'master'
  }
}