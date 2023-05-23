<div align="center" width="100%">
    <img src="https://i.toaaa.de/i/f9dq9.png" width="128"/>
</div>

<div align="center" width="100%">
    <h2>SchneileTV VOD Archiv</h2>
    <p>Stack: Go, Gin, Gorm, FFmpeg</p>
</div>

## 🐳 Deploy

Copy `.env.sample` to `.env` and replace the required variables.

### Example `docker-compose.yml`

```
version: '3'
services:
  api:
    container_name: archiv-api
    build: .
    restart: unless-stopped
    env_file: .env
    ports:
      - 127.0.0.1:5000:5000
    volumes:
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
      - /path/to/media/:/var/www/media/
    depends_on:
      - db
  db:
    container_name: archiv-db
    image: postgres:15-alpine
    restart: unless-stopped
    env_file: .env
    volumes:
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
      - /path/to/postgres:/var/lib/postgresql/data
```

## 🚪 Reverse Proxy

My preferred way is to use NGINX Proxy Manager with a `Advanced Configuration`.

```
location / {
    proxy_pass http://192.168.10.36:5000;
    add_header Access-Control-Allow-Origin "*";
    add_header X-Content-Type-Options "nosniff";
    add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; object-src 'none'; base-uri 'self'; form-action 'self'; frame-ancestors 'none';";
    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-XSS-Protection "1; mode=block";
    add_header Referrer-Policy "same-origin";
    add_header Cache-Control "no-store, no-cache, must-revalidate, proxy-revalidate";
    gzip on;
}
```

## Backup and restore database

Required to upgrade major postgres versions.

**Backup**
`docker exec -t archiv-db pg_dumpall -c -U YOUR_DB_USER > /path/to/backup/dump_$(date +%Y-%m-%d"_"%H_%M_%S).sql`

**Restore**
`docker exec -i archiv-db psql -d YOUR_DB_NAME -U YOUR_DB_USER < /path/to/backup/dump_<some-date>.sql`
