# AlertMetrics

## Config 
Config file must be storage in past ./config.yaml
### Settings
* db.connection - set postgres connection string
* period - set time for check period in seconds
* metrics
    * Name must be equal in DB
    * lowerBound - minimal bound of Value. If less than it, will be notify alert
    * upperBound - maximal bound of Value. If greater than it, will be notify alert
* notifiers
    * log - send notify at console
        * enable - if true, will be work

## Tests
You can use migration.sql for insert data (Schema name exist in Code!)