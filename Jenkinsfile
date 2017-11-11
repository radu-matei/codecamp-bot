podTemplate(label: 'mypod', containers: [
    containerTemplate(name: 'docker', image: 'docker:17.09', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'kubectl', image: 'lachlanevenson/k8s-kubectl:v1.8.0', command: 'cat', ttyEnabled: true)
  ],
  volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
  ]) {

    node('mypod') {

        checkout scm

        stage('build and push the k8s client') {
            container('docker') {

                withCredentials([[$class: 'UsernamePasswordMultiBinding', 
                        credentialsId: 'dockerhub',
                        usernameVariable: 'DOCKER_HUB_USER', 
                        passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
                    
                    sh """
                        cd go-client
                        docker build -t ${env.DOCKER_HUB_USER}/codecamp-client:${env.BUILD_NUMBER} .
                        """
                    sh "docker login -u ${env.DOCKER_HUB_USER} -p ${env.DOCKER_HUB_PASSWORD} "
                    sh "docker push ${env.DOCKER_HUB_USER}/codecamp-client:${env.BUILD_NUMBER} "
                }
            }
        }


        stage('build and push the bot image') {
            container('docker') {

                withCredentials([[$class: 'UsernamePasswordMultiBinding', 
                        credentialsId: 'dockerhub',
                        usernameVariable: 'DOCKER_HUB_USER', 
                        passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
                    
                    sh """
                        cd bot
                        docker build -t ${env.DOCKER_HUB_USER}/codecamp-bot:${env.BUILD_NUMBER} .
                        """
                    sh "docker login -u ${env.DOCKER_HUB_USER} -p ${env.DOCKER_HUB_PASSWORD} "
                    sh "docker push ${env.DOCKER_HUB_USER}/codecamp-bot:${env.BUILD_NUMBER} "
                }
            }
        }


        stage('update kubernetes deployment for bot and k8s client') {
            container('kubectl') {

                withCredentials([[$class: 'UsernamePasswordMultiBinding', 
                        credentialsId: 'dockerhub',
                        usernameVariable: 'DOCKER_HUB_USER',
                        passwordVariable: 'DOCKER_HUB_PASSWORD']]) {

                    sh "kubectl set image deployment/go-client go-client=${env.DOCKER_HUB_USER}/codecamp-client:${env.BUILD_NUMBER} "
                    sh "kubectl set image deployment/bot-deployment bot-deployment=${env.DOCKER_HUB_USER}/codecamp-bot:${env.BUILD_NUMBER} "
                }
            }
        }
    }
}
