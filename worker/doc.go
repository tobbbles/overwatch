// package worker provides a cron-like way of continually updating our database from the remote Overwatch API.
// Worker consumes data on a time interval, and will (currently) default to overwrite data in the database.
// NOTE: Whilst overwriting data is an option currently - only overwrite is implemented, to save on time
// due to this being for demonstration. In future, you'd likely want to implement diffing and partial updates to data.
package worker
