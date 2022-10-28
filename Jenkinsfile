pipeline {
  agent any

  stages {
    stage("检出") {
      steps {
        checkout(
          [$class: 'GitSCM',
          branches: [[name: GIT_BUILD_REF]],
          userRemoteConfigs: [[
            url: GIT_REPO_URL,
              credentialsId: CREDENTIALS_ID
            ]]]
        )
      }
    }
    stage('编译') {
      steps {
        sh "go env && unset GO111MODULE && go env -w GO111MODULE=on && unset GOPROXY && go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && go mod download && go build -o fwds-app ."
      }
    }
    stage(' 上传到 generic 仓库') {
      steps {
        codingArtifactsGeneric(
          files: "fwds-app",
          repoName: "fwds",
        )
      }
    }
  }
}