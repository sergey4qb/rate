host=$(awk -F "=" '/POSTGRES_HOST/ {print $2}' ./configs/postgresql.conf)
port=$(awk -F "=" '/POSTGRES_PORT/ {print $2}' ./configs/postgresql.conf)
user=$(awk -F "=" '/POSTGRES_USER/ {print $2}' ./configs/postgresql.conf)
password=$(awk -F "=" '/POSTGRES_PASSWORD/ {print $2}' ./configs/postgresql.conf)
dbname=$(awk -F "=" '/POSTGRES_DB/ {print $2}' ./configs/postgresql.conf)

db_url="postgresql://$user:$password@$host:$port/$dbname?sslmode=disable"

if [ "$1" == "up" ]; then
    command="up"
elif [ "$1" == "down" ]; then
    command="down"
elif [ "$1" == "force" ]; then
      if [ -z "$2" ]; then
            echo "Usage: $0 force <version>"
            exit 1
        fi
        force="force $2"
else
    echo "Usage: $0 {up|down} or force <version>"
    exit 1
fi

echo "Executing: migrate -path ./migrations -database \"$db_url\" $command $force"

migrate -path ./migrations -database "$db_url" $command