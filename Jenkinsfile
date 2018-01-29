/**
 * @author [Gopi Karmakar]
 * @email [gopi.karmakar@monstar-lab.com]
 * @create date 2018-01-26 03:00:43
 * @modify date 2018-01-26 03:00:43
 * @desc [description]
*/

pipeline {
    agent any
    stages {
        stage('build base image') {
            steps {
                echo 'Building Base Docker Image...'
                sh 'make build-base'
            }
        }
        
        stage('run tests') {
            steps {
                echo 'Running Test...'
                sh 'make build-test'
                sh 'make test-unit'
                sh 'ls'
                junit 'report/report.xml'
            }
        }

        stage('build image') {
            steps {
                echo 'Building Docker Image...'
                sh 'make build'
            }
        }
    }
}