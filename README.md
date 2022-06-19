# BB2 Alert
This tool captures live Blood Bowl 2 data stream and alerts the user on Discord that match has started.

## Configuration

1. Create discord webhook for a channel you want to receive the notifications into. See https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks

Example:

  ```
  https://discord.com/api/webhooks/<webhook_id>/<token>
  
  ```
2. Get your Discord UserID. See https://www.alphr.com/discord-find-user-id/

3. Update config.yaml
```
discord:
  webhook_id: <webhook_id>
  token: <token>
  user_id: <UserID>
```
4. Start bb2-alert
5. Start BB2 and spin
6. Get annoyed that you are 500TV down agaist killer chorfs