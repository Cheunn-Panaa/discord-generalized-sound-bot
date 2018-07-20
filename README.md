# Discord Generalized Sound Bot

Discord Generalized Sound Bot is a bot that can play sound on your Discord server.

## Features

* Playing a sound on your Discord server.
* You can add sounds with web interface.
* You can search sounds on the web.
* No limit of the number of sounds.

## Usage

1. Create config files

      `.env`

      ```shell
      BACKEND_URL=http://localhost:3000
      BACKEND_VIRTUAL_HOST=backend.example.com
      BACKEND_LETSENCRYPT_HOST=backend.example.com
      BACKEND_LETSENCRYPT_EMAIL=example.com
      ADMIN_VIRTUAL_HOST=admin.example.com
      ADMIN_LETSENCRYPT_HOST=admin.example.com
      ADMIN_LETSENCRYPT_EMAIL=example.com
      WEB_VIRTUAL_HOST=example.com
      WEB_LETSENCRYPT_HOST=example.com
      WEB_LETSENCRYPT_EMAIL=example.com
      ADMIN_TITLE=Admin title here
      WEB_TITLE=Web title here
      WEB_DESCRIPTION=Description
      WEB_OG_DESCRIPTION=Og description
      WEB_OG_URL=Og url
      WEB_OG_IMAGE=Og image
      WEB_TWITTER_CARD=summary_large_image
      ```

      `backend/config/environment/production.yaml`

      ```yaml
      db:
        url: "data/production/sounds.db"
      server:
        port: "3000"
        cors:
          - "http://example.com"
          - "http://admin.example.com"
      data:
        path: "data/production"
        prefix: "sounds_dca"
      auth:
        secret: "secret"
        username: "admin"
        password: "Hashed password"
      ```

      You can make password using genpass

      ```shell
      docker build -t genpass genpass/
      docker run --rm -it genpass genpass
      ```

      `bot/config.yml`

      ```yaml
      botToken: [Discord bot token]
      dbUrl: /sounds_dca/sounds.db [Where the sounds.db is]
      botPrefix: [Bot command prefix]
      botHello: [Bot hello message]
      botPlaying: [Bot status]
      botNotFound:  [The message when sound not found]
      soundDir: /
      name: [Bot name]
      applicationId: [Discord application id]
      soundCacheSize: 100 [The number of cached sounds]
      MaxQueueSize: 6 [The size of queue per guild]
      ```

1. `docker-compose up -d`


## Development

I'm maintaining backend with test.

```shell
docker-compose -f docker-compose-test.yml run --rm backend-test
```

### TODO

- [x] Making a bot.
- [x] Making a web site to search sound name.
- [x] Making a admin site to add sounds.
- [x] Add Docker things.
- [ ] Implement session shading.
- [ ] Implement a command to search sounds.