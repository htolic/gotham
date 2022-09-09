### Build app source and create container image
    docker build -t getweather:dev .

### Run the weather service
    docker run --rm -e OWM_API_KEY="put_api_key_here" -e OWM_CITY="Honolulu" getweather:dev

### Check syslog for container output
    grep -i honolulu /var/log/syslog
