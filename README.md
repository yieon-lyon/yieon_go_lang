# yieon의 Golang Study

###[Heroku](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)

###[go framework stars](https://github.com/mingrammer/go-web-framework-stars)

###[gin-gonic](https://github.com/gin-gonic/gin)

##[1주차 스터디 - 가볍게 Go 입문하기](https://www.notion.so/spacewalkcorp/1-Go-ad168db2bd554142a0d56e307a1d9202)

##[2주차 스터디 - Github Action을 통한 Go-Heroku 배포 따라하기](https://www.notion.so/spacewalkcorp/2-Github-Action-Go-Heroku-eb303c252b7d446fa694afe0bdee213a)

1. Heroku 회원가입
2. Github 회원가입
3. Git Organization 등록하기
4. Go Project 만들기
5. Github에 Project 올리기
6. Github Action 설정하기
    1. Github Secrets `HEROKU_API_KEY` 추가하기
7. 오류 발생
    ![img.png](doc/study2weeks-issue.png)
8. 오류 해결하기
    1. .github/workflows/myAction.yml파일에서 다음과 같이 수정합니다.

        ```yaml
        name: myAction
        
        on:
          push:
            branches: [main]
          pull_request:
            branches: [main]
        
        jobs:
          build:
            runs-on: ubuntu-latest
            steps:
              - uses: actions/checkout@v2
        
              - name: Set up Go
                uses: actions/setup-go@v2
                with:
                  go-version: 1.16
        
              - name: Build
                run: go build -v -o ./bin/[Heroku에등록할이름(중복불가)]
        
              - uses: akhileshns/heroku-deploy@v3.12.12 # This is the action
                with:
                  heroku_api_key: ${{secrets.HEROKU_API_KEY}}
                  heroku_app_name: [Heroku에등록할이름(중복불가)] #Must be unique in Heroku
                  heroku_email: [Heroku계정]
        ```

       다음 예시와 같이 수정하시면 됩니다.

        ```yaml
        - name: Build
          run: go build -v -o ./bin/yieon-go-heroku
        
        - uses: akhileshns/heroku-deploy@v3.12.12 # This is the action
                with:
                  heroku_api_key: ${{secrets.HEROKU_API_KEY}}
                  heroku_app_name: "yieon-go-heroku" #Must be unique in Heroku
                  heroku_email: "parrotbill@naver.com"
        ```

9. Action 상태 확인하기
![img.png](doc/study2weeks-action.png)
10. 배포 확인하기
    ![img.png](doc/study2weeks-deploy.png)
    1. [https://yieon-go-heroku.herokuapp.com/](https://yieon-go-heroku.herokuapp.com/)
    


##[3주차 스터디 - Templates 동작](https://www.notion.so/spacewalkcorp/3-Templates-ba8d873a2f5f4b879b3610e15696a51b)


