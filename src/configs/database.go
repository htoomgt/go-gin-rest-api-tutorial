package configs

/*DBConf is configuration parameters for database connections*/
type DBConf struct {
}

/*MySQLDbHost is a constant for MySQL Database host IP address */
const MySQLDbHost = "127.0.0.1"

/*MySQLDbPort is a constant for default MySQL port*/
const MySQLDbPort = "3306"

/*MySQLDbUsername is a constant for default MySQL connection username*/
const MySQLDbUsername = "root"

/*MySQLDbPassword is a constant for deafult MySQL connection password*/
const MySQLDbPassword = "abcd1234"

/*MySQLDbNAME is a constant for default MySQL database name*/
const MySQLDbNAME = "db_test"

/*MySQLDbTimeout is a constant for maximum query execution time in second*/
const MySQLDbTimeout = 30
