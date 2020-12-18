# Accounting #

## Introduction ##

### Vendors ###

Following packages are imported and need to be installed

```sh
go get -u github.com/google/uuid
go get -u github.com/imdario/mergo
go get -u github.com/stretchr/testify
go get -u gorm.io/gorm
go get -u gorm.io/datatypes
go get -u gorm.io/driver/postgres

```

### Testing ###

Following functions below for testing different methodolgies are installed:

* TestOrderTypesValues
* TestStruct2Map
* TestMakeMaps
* TestLoopMap
* TestForLoopWithOrderTypes

#### TestOrderTypesValues ####

Use this command:

```sh
go test -v accounting_test.go -run TestOrderTypesValues
```

#### TestStruct2Map ####

Use this command:

```sh
go test -v accounting_test.go -run TestStruct2Map
```

#### TestMakeMaps ####

Use this command:

```sh
go test -v accounting_test.go -run TestMakeMaps
```

#### TestLoopMap ####

Use this command:

```sh
go test -v accounting_test.go -run TestLoopMap
```

#### TestForLoopWithOrderTypes ####

Use this command:

```sh
go test -v accounting_test.go -run TestForLoopWithOrderTypes
```
