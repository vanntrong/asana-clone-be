export DC_PROJECT=shope-nest
export ENV_FILE=docker.env

backup_db() {
    TIME=$(date '+%Y-%m-%d')
    echo "Current time : $now"

    echo "Backing up database..."
    docker-compose -p $DC_PROJECT --env-file $ENV_FILE -f docker-compose.yml exec -u root db pg_dump --inserts -U ${DB_USERNAME} -d ${DB_DATABASE} -f /var/backups/$TIME.sql
    docker-compose -p $DC_PROJECT --env-file $ENV_FILE -f docker-compose.yml cp db:/var/backups/$TIME.sql ${BACKUP_DIR}/$TIME.sql
}

backup_db