# Golang client for Spring Cloud Config
## Introduction
### Features
* Able to connect and get application profile from Spring Cloud Config Server
* Stream configuration data to Viper
* Total compatible with Viper

### About SCC
Remotly configuration management for cloud application.
Separate application config into multi profiles in order to deploy one application package to multi enviroment using only one SCC server.

### SCC & Viper relationship?
SCC is not a on of Viper remote provider. Currently, Viper only supports 2 key-value etcd and consul.
SCC is compatible client of SCC and use Viper as configuration getter, setter. 

## Flow
SCC -> Viper read local config -> SCC get app-name & active profile
SCC get config server endpoint
SCC get full config of app-name profile: endpoint/app-name/active-profile
SCC override default config by new config that given from config server

SCC wrapper all Get-Set functions of Viper