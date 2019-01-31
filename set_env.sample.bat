:: Name:     set_env.sample.bat
:: Purpose:  Configures enviroment variables
:: Author:   Leandro Souza
:: Revision: January 2019 - initial version

@ECHO OFF

:: variables

:: default profile
SET DATABASE_NAME=dbname
SET DATABASE_USERNAME=username
SET DATABASE_PASSWORD=password
SET DATABASE_SERVER=localhost
SET DATABASE_PORT=3306

echo All set!


:END
ENDLOCAL
ECHO ON
@EXIT /B 0