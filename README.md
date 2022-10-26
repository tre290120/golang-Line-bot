# golang-Line-bot
A simple Line bot with golang

## Tech stack
1. ngrok
2. docker
3. line sdk
4. golang
5. golang-vipre
6. golang-cobra
7. golang-mongo-diver
8. gin
9. mongo

## How to use

build golang code
```
./build.sh
```

Pull mongo docker image
```
docker pull mongo:4.4
```

Run mongo image
```
./mongo.sh
```

Run ngrok
```
ngrok http 4000
```
ngrok will give you a link and the link forward http request to localhost:4000
![ngrok](https://user-images.githubusercontent.com/28688049/198027967-e9339925-d219-47dc-9fbc-a1b947daee29.png)
Set the link/callback as webhook url in your [line developer](https://developers.line.biz/zh-hant/ "link") setting page


Get your line config including channel secret and channel token
Please refer to below link to get a token
https://developers.line.biz/en/reference/messaging-api/#issue-shortlived-channel-access-token
Get your channel secret in your basic setting section in the [line developer](https://developers.line.biz/zh-hant/ "link") sites
Create a yaml file "line_config.yaml" with the format
```
line:
  token: "xxxxx"
  secret: "xxxx"
```

Run golang binary
```
./bot -c $(pwd) -p 4000
```
-c:path to line_config 
-p:the port ngrok forward to

## Test
1. open your line and add the bot as freind
2. send any message
3. you will receive the echo message
4. send "#ls"
5. bot will reply you all the messages you have sent

## Reference
https://developers.line.biz/en/reference/messaging-api/#verfiy-channel-access-token-v2-1
https://www.learncodewithmike.com/2020/06/python-line-bot.html?fbclid=IwAR1yzMljhtza2d86jxW14tC6xsxlEEVR8aDN5Y229zylzWSkyP9yDLP5vXM
https://developers.line.biz/en/docs/
