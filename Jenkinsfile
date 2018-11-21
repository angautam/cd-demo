  env.DOCKERHUB_USERNAME = 'angautam'

  node("docker-prod") {
    stage("Production") {
      try {
        // Create the service if it doesn't exist otherwise just update the image
        sh '''
          SERVICES=$(docker service ls --filter name=cd-demo --quiet | wc -l)
          if [[ "$SERVICES" -eq 0 ]]; then
            docker network rm cd-demo || true
            docker network create --driver overlay --attachable cd-demo
            docker service create --replicas 3 --network cd-demo --name cd-demo -p 8090:8080 ${DOCKERHUB_USERNAME}/cd-demo:${BUILD_NUMBER}
          else
            docker service update --image ${DOCKERHUB_USERNAME}/cd-demo:${BUILD_NUMBER} cd-demo
          fi

      }catch(e) {
        sh "docker service update --rollback  cd-demo"
        error "Service update failed in production"
      }finally {
        sh "docker ps -aq | xargs docker rm || true"
      }
    }
  }
