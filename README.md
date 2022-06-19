# BB2 Alert
This tool captures live Blood Bowl 2 data stream and alerts the user on Discord that match has started.

## Configuration

1. Create discord webhook for a channel you want to receive the notifications into. See https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks

Example:

  ```
  https://discord.com/api/webhooks/988079967420166194/J76ho2FPfAqX_NoXzGu5Ec0ZSP-upvCgZ5T3adwKQB72S7zyZU0xrTM8muwMmiwojjFJ
  ServerID = 988079967420166194
  Token = J76ho2FPfAqX_NoXzGu5Ec0ZSP-upvCgZ5T3adwKQB72S7zyZU0xrTM8muwMmiwojjFJ
  
  ```
2. Get your Discord UserID. See https://www.alphr.com/discord-find-user-id/

3. Update config.yaml
```
discord:
  server_id: <ServeID>
  token: <Token>
  user_id: <UserId>
```
4. Start bb2-alert
5. Start BB2 and spin
6. Get annoyed that you are 500TV down agaist killer chorfs