# Overview

The config package provides a singleton configuration loader that reads configuration data from a JSON file and optionally overrides it with environment variables. This package ensures that the configuration is loaded only once and can be accessed safely from different parts of your application.

# Folder Structure

The folder structure in project should look like this:
```
go

/
├── config/
│   ├── config-local.json
│   └── config-dev.json
├── src/
│   └── src.go
└── main.go
```

# Configuration File

The configuration file should be a JSON file located in a /config directory. The file name should follow the pattern config-<environment>.json, where <environment> is the environment for which the configuration is being loaded (e.g., dev, prod).

Example config-dev.json:
```
json

{
    "DatabaseURL": "localhost:5432",
    "APIKey": "my-secret-api-key"
}
```
# Usage
1. Define Your Configuration Structure

Define a struct that represents your configuration in the example/config package.
```
go

package config

type MyConfig struct {
    DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
    APIKey      string `json:"api_key" env:"API_KEY"`
}
```
2. Initialize the Configuration

In your main function, initialize the configuration by specifying the environment.
```
go

package main

import (
    "fmt"
    "github.com/kowiste/config"
    conf "example/config"
    pkg1 "example/pkg1"
    pkg2 "example/pkg2"
)

func main() {
    // Initialize the configuration
    err := config.New[conf.MyConfig](config.GetPathEnv("dev"))
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    // Get the loaded configuration
    cfg, err := config.Get[conf.MyConfig]()
    if err != nil {
        fmt.Println("Error getting config:", err)
        return
    }
    fmt.Println("main", cfg)

    // Use the configuration in other packages
    pkg1.Test()
    pkg2.Test()
}
```

3. Access the Configuration in Other Packages

In other packages, you can access the configuration using the config.Get function.

```
go

package pkg1

import (
    "fmt"
    "github.com/kowiste/config"
    conf "example/config"
)

func Test() {
    cfg, err := config.Get[conf.MyConfig]()
    if err != nil {
        fmt.Println("Error getting config in pkg1:", err)
        return
    }
    fmt.Println("pkg1", cfg)
}
```

# Environment Variable Overrides

You can override configuration values with environment variables by specifying the env tag in the struct fields. The environment variable should match the key specified in the env tag.

Example:
```
go

type MyConfig struct {
    DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
    APIKey      string `json:"api_key" env:"API_KEY"`
}
```
If you set the environment variable DATABASE_URL, it will override the value from the JSON configuration file.
Error Handling

The config.New function initializes the configuration and should be called only once, typically in the main function. If the configuration fails to load, it returns an error.

The config.Get function retrieves the loaded configuration. If the configuration has not been loaded or there is a type mismatch, it will returns an error.

