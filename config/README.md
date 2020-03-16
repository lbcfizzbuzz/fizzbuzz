# Config
The server will find its configuration in a JSON file containing the following properties:
- `datastore` (`string`): the type of the datastore used, the `Datastore` interface (`datastore/datastore.go`) must be implemented
to be able to use a new datastore
- `dsn` (`string`): the connection string of the datastore
- `port` (`int`): the port that the server should listen on (must be in the range 1024 - 65534)
- `log_in_json` (`bool`): if true, the logger will use JSON as output format, otherwise plain text
- `log_file_path` (`string`): if empty, the logger will use stdout as output, otherwise the given file