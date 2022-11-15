## Nginx parsing

## Goals
* Get logs of all docker containers
* If they match the nginx format then parse it according to the given regex.

## Steps to run
* Build the log generator `docker build -t tlog .`
* Start Signoz `docker-compose up -d`
* open the otel-collector-config.yaml inside `clickhouse-setup` for the configuration. 
* Run the generator `docker run tlog`
* Open SigNoz on `localhost:3301`
* Create an account and go to logs.
* Select the time range to be one week.
  
 ## Regex testing https://regex101.com/r/2HSyk2/1

## Note
* For testing purposes `add-filter-attr` operator `filter/logs` is added. It can be removed once testing is done
