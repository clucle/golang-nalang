### 완성본

![gif](https://github.com/clucle/golang-nalang/blob/master/src/console/snake-game/resource/snake.gif)

### 개발 과정

콘솔에서 게임을 진행하기 위해 타이머를 돌며 게임을 진행하고 출력해야한다

`channels` 기능을 사용해보자

업데이트 마다 스크린을 정리해줘야 한다.
`fmt.Print("\033[H\033[2J")`

일단 키이벤트 받아와서 키 상태값은 저장 했는데, 출력이 지금은 부자연스러움
exe 로 만들어봐야하나