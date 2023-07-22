# GoKit

## Installation

```
go get github.com/robertgontarski/gokit
```

## What it contains

### ENV

Retrieves and saves configurations found in **.env** files or its override **.env.local**

The entire configuration is available under the **Env** variable

```go
var (
    Env map[string]string
)
```

Example of use

```go

Env["APP_ENV"]

```

### Logger

Provides several types of predefined logs that write to files

In order to set the path as well as the name of the file, we need to have the corresponding parameters in env

The basic name is **file** in the **main folder** of the project

```dotenv
APP_LOG_FILE="file_name"
APP_LOG_PATH="file_path"
```

To use the logs, call one of the predefined variables

```go
var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)
```

Example of use 

```go
ErrorLogger.Println(err)
```