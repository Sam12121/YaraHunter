# Jenkins example for Toae YaRadare

This project demonstrates using Toae YaRadare in Jenkins build pipeline.
After customer's image is built, Toae YaRadare is run on the image and Malware results are either shown in STDOUT(build logs) or in a JSON file.
Refer [this](https://github.com/Sam12121/YaraHunter#command-line-options) for command line options


## Steps
- Ensure `toaeio/toae-yaradare:latest` image is present in the vm where jenkins is installed.
```shell script
docker pull toaeio/toae-yaradare:latest
```
### Scripted Pipeline
```
    stage('Run Toae YaRadare'){
        ToaeAgent = docker.image("toaeio/toae-yaradare:latest")
        try {
            c = ToaeAgent.run("--name=toae-yaradare -v /var/run/docker.sock:/var/run/docker.sock", "--image-name ${full_image_name}")
            sh "docker logs -f ${c.id}"
            def out = sh script: "docker inspect ${c.id} --format='{{.State.ExitCode}}'", returnStdout: true
            sh "exit ${out}"
        } finally {
            c.stop()
        }
    }
```
### Declarative Pipeline
```
stage('Run Toae Vulnerability Mapper'){
    steps {
        script {
            ToaeAgent = docker.image("toaeio/toae-yaradare:latest")
            try {
                c = ToaeAgent.run("--name=toae-yaradare -v /var/run/docker.sock:/var/run/docker.sock", "--image-name ${full_image_name}")
                sh "docker logs -f ${c.id}"
                def out = sh script: "docker inspect ${c.id} --format='{{.State.ExitCode}}'", returnStdout: true
                sh "exit ${out}"
            } finally {
                c.stop()
            }
        }
    }
}
```