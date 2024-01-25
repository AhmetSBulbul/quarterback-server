# Quarterback 

## Backend Server

```sh
    ## PROD
    # Run both app and database
    docker-compose up -d  

    ## DEV
    # Run both app and database
    docker-compose -f docker-compose.dev.yml up -d  
    # Run database only
    docker-compose -f docker-compose.dev.yml up database
    # Re-build database
    docker-compose -f docker-compose.dev.yml up --build database

```

## API

Build proto files for API:

```sh
    make gopb
```

Build API binary
```sh
    make goapi_linux
    # Or
    make goapi_macos
```

Run API binary
```sh
    ./bin/api/bin
```

### Development

```sh
    cd api
    cp .env.example .env
    go run main.go
```

## To-Do

- [x] Court Create and Get
- [x] Courts Seed min 5-6 for main districts (Bornova, Karsiyaka). 1-2 is enough for others,
- [x] Comment on courts
- [?] Add media for court (@AhmetSBulbul)
- [ ] Comment to game(@AhmetSBulbul)
- [x] If possible court check-in(pazarlarken is yapar redis varsa ttl'li falan deneriz sunumluk olacak)
- [x] Get Courts by location to show on home page and start game screen(@okb)
- [ ] Create Game (@okb)
- [ ] Start Game (@okb)
- [ ] End Game (@okb)
- [ ] Cancel Game (@okb)
- [ ] Scores and points (@okb)

Not: docker compose dev dosyasinda calisiyorsan production dosyasini da guncellemeyi unuttma config degisikliginde.

### Good to have
- [x] Simple chat just for presentation don't need to be fully functional we can store the data on mongo, redis or sql doesn't matter.

- [ ] Create Team
- [ ] Team Match


## MUST HAVE IF WE HAVE PLAN TO MAINTAIN IT

- [ ] Allah rizasi icin katmanlama join atmaktan bileklerim kireclendi.
- [ ] Understandable and traceable error messages on returns
