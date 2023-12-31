# Init backend
sudo docker compose up --build
sudo docker compose up --build -d db
sudo docker compose up --build api

# db
psql -d numbers -U postgres

# SH
sudo docker exec -ti odd-or-even-api-1 /bin/bash
sudo docker exec -ti odd-or-even-db-1 /bin/bash
