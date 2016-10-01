# LINEMSGBot
This is Echo Bot using LINE Messaging API on GAE/Go  

This program needs the Go SDK for the LINE Messaging API https://github.com/line/line-bot-sdk-go .  
To deploy Google App Engine, app.yaml file likes below is needed.  

--- app.yaml ---  
application: YOURGAEPROJECTNAME  
version: 1  
runtime: go  
api_version: go1  
  
handlers:  
- url: /.*  
  script: _go_app  
  
env_variables:  
  CHANNEL_SECRET: 'YOURSECRET'  
  CHANNEL_TOKEN: 'YOURTOKEN'  
--- app.yaml ---  

YOURSECRET is 'Channel Secret' in Basic information of LINE Developers.  
YOURTOKEN is 'Channel Access Token' in Basic information of LINE Developers.  
LINE Developers are in LINE BUSINESS CENTER.  

This program is branched from https://github.com/line/line-bot-sdk-go/blob/master/examples/echo_bot/server.go  
Thanks.  
