pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'pwd'
        dir(path: '/home/budi/development/build')
        git(url: 'https://github.com/budiariyanto/build-testing', branch: 'master', changelog: true)
      }
    }
  }
}