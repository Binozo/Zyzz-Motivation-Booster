# [Zyzz](https://youtu.be/yDqk6KJVyP8) Motivation Booster

A Bot which sends you a motivational zyzz video every 3 hours through telegram.

**Note:** My opinion does not necessarily reflect the opinions of the video creators.

## Setup
```bash
docker run --env telegrambottoken={$YOUR_TOKEN} --env telegramchatid={$YOUR_CHAT_ID} -d --restart unless-stopped --name zyzz_motivation_booster ghcr.io/binozo/zyzz-motivation-booster:latest
```