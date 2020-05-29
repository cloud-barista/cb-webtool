cb-webtool
==========
cb-webtool은 Multi-Cloud Project의 일환으로 다양한 클라우드를 cb-webtool에서 처리해 <br>
사용자로 하여금 간단하고 편안하게 클라우드를 접할 수 있게 해준다.
***
## [Index]
1. [설치 환경](#설치-환경)
2. [의존성](#의존성)
3. [소스 설치](#소스-설치)
3. [환경 설정](#환경-설정)
4. [서버 실행](#서버-실행)
***
## [설치 환경]
cb-webtool은 1.12 이상의 Go 버전이 설치된 Windows & Linux & Mac OS 등에서 실행 가능합니다.<br>
Windows 8/10 & Mac OS 10.14.6(Mojave) 환경에서 개발을 진행했으며 최종 전체 동작을 검증한 OS는 Ubuntu 18.0.4입니다.<br>
<br>

## [의존성]
cb-webtool은 내부적으로 cb-tumblebug & cb-spider & cb-dragonfly 프로젝트를 이용하며, <br>
개별 프로젝트들의 문서를 통해서 동일 서버 또는 개별 서버에 설치 및 실행합니다.<br>
- [https://github.com/cloud-barista/cb-tumblebug](https://github.com/cloud-barista/cb-tumblebug) README 참고하여 설치 및 실행
- [https://github.com/cloud-barista/cb-spider](https://github.com/cloud-barista/cb-spider) README 참고하여 설치 및 실행
- [https://github.com/cloud-barista/cb-dragonfly](https://github.com/cloud-barista/cb-dragonfly) README 참고하여 설치 및 실행

## [소스 설치]
- Git 설치
  - `# apt update`
  - `# apt install git`

- Go 설치
  - https://golang.org/doc/install <br>
    (2020년 05월 현재 `apt install golang` 명령으로 설치하면 1.10 버전이 설치되므로 위 링크에서 1.12 이상의 버전으로 설치할 것)
  - `wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz`
  - `tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz`
  - `.bashrc` 파일 하단에 다음을 추가: 
  ```
  export PATH=$PATH:/usr/local/go/bin
  export GOPATH=$HOME/go
  ```

- `.bashrc` 에 기재한 내용을 적용하기 위해, 다음 중 하나를 수행
  - bash 재기동
  - `source ~/.bashrc`
  - `. ~/.bashrc`

 - echo 설치
    ````bash
      $ go get -u -v github.com/labstack/echo
    ````
 
 - echo-session 설치
     ````bash
       $ go get -u -v github.com/go-session/echo-session
     ````

 - reflex 설치 (Windows 미지원 / Windows에 bash 설치 시 사용 가능)
     ````bash
       $ go get github.com/cespare/reflex 
     ````

 - cb-webtool 설치
     ````bash
       $ go get github.com/cloud-barista/cb-webtool
     ````

## [환경 설정]
   - 의존성 프로젝트를 다른 서버에 설치한 경우 URL 설정<br>
     conf/setup.env 파일에 cb-tumblebug & cb-spider & cb-dragonfly의 URL 정보를 수정합니다.
   
   - 초기 Data 구축<br>
https://github.com/cloud-barista/cb-spider의 [API규격] 및 [활용 예시]를 참고해서 CLI및 json 방식의 웹 호출로 데이터 구축이 가능합니다.


## [서버 실행]
- Linux & Mac OS에서 실행
    ````bash (Linux & Mac OS)
    $ cd github.com/cloud-barista/cb-webtool
    $ run.sh
    ````

- Bash를 설치하지 않은 Windows 환경에서는 reflex를 사용할 수 없으므로 직접 구동해야 합니다.
    ````bash (Windows)
    $ cd github.com/cloud-barista/cb-webtool
    $ source ./conf/setup.env
    $ go run main.go
    ````
