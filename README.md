# Bob

Simple bot for Telegram which listens for events on an REST API from your CI system (presumably, Jenkins) and announces them on a channel.

Right now it only speaks Spanish.

## Config
Two environment variables:
- `TG_TOKEN`: Bot token from the Botfather *(required)*
- `TG_GROUP_ID`: Telegram Group ID where to send notifications *(optional)*

### Bot token
To get one (and therefore a new bot), talk to the @BotFather

### Group ID
To figure out a Group ID, start the bot without the `TG_GROUP_ID` variable, and add it to a group. The ID will be printed in stdout.  
From then on, you can provide it via environment. Keep in mind that the bot must be in a group in order for it to talk in it, even if you provide the ID.

## Dev/Dependencies
`dep` is used for dependency management. Do a `dep ensure` (or `make prepare`) to download dependencies.

Dep: https://github.com/golang/dep

### Release
Run `make release`. Compilation will be made inside the provided `Dockerfile` and a new image will be created.

## Events
The bot listens on a REST API, on port 8888.

The URLs are:
- `/notify/build`: Notify of new builds.
- `/notify/pr`: Notify of new Pull Requests

### Payloads

Sample build event payload:

    {
        "project": "test project",
        "result": "SUCCESS",
        "phase": "finished",
        "build_url": "http://reddit.com"
    }

Supported build phases are: `started`, `finished`, `aborted`, `waiting`. All others are printed as "unknown"
    
Sample Pull Request event payload:

    {
        "project": "test project",
        "target": "master",
        "changeId": "00",
        "author": "Some One",
        "changeUrl": "http://imgur.com"
    }

Examples are provided in the `send-events.sh` file, which is obviously a shitty way to test messages.

### Sample implementation
In your Jenkinsfile you could declare the following method:

    def notifyBuild(String event, String result = null) {
        httpRequest(url: "${botUrl}/build", contentType: 'APPLICATION_JSON', httpMode: 'POST', requestBody: """
        {
            "project": "${JOB_NAME}",
            "result": "${result != null ? result : "-"}",
            "phase": "${event}",
            "build_url": "${BUILD_URL}"
        }
        """)
    } 

Then, in the proper pipeline steps:
    
    stage('Testing & analysing') {
        steps {
            notifyBuild "started"
            withSonarQubeEnv('Sonar') {
                ansiColor('xterm') {
                    sh "sbt clean coverage test"
                }
            }
        }
    }
    
    [...]
    
    post {
        always {
            notifyBuild "finished" "${currentBuild.currentResult}" 
        }
    }

## TODO
- Add other languages
- Do proper tests
- Accept commands
