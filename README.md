[![CircleCI](https://circleci.com/gh/YasushiKobayashi/go-bot.svg?style=svg)](https://circleci.com/gh/YasushiKobayashi/go-bot)

### aboout
- slack notice by slack webhook
    - for Backlog

### env
- golang version 1.8

### set up golang & direnv

```bash
brew install go
brew install glide
brew install direnv
echo 'eval "$(direnv hook bash)"' >> ~/.bash_profile
```

### setup project
```bash
cd src/app
glide install
go run main.go
```

### godoc
`godoc -http=:6060`

### run with docker
```shell
docker-compose up --build

// update
docker-compose build --no-cache
docker-compose up
```
