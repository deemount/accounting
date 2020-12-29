# Accounting #

A online banking related application focussed on transactions
with different order types.

## Introduction ##

Creating different types of lists for exchange orders by quering transactions and
customer data in a PostgreSQL Database. This combination is mostly used on transactions
with ATM's

The Exchange Order Types, within a transaction, are:

* withdrawal:0
* buy:1
* spread:2
* fee:3

## Requirements ##

* Go 1.15.x

### Dependencies ###

Following packages are imported and need to be installed

```sh
go get -u github.com/google/uuid
go get -u github.com/imdario/mergo
go get -u github.com/stretchr/testify
go get -u gorm.io/gorm
go get -u gorm.io/datatypes
go get -u gorm.io/driver/postgres

```

### To Do's ###

* add database driver (GORM)
* add calculations to the list for spread, fee and withdrawal
