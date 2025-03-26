# DefaultConfig - Go Configuration Loader with Default Values

`DefaultConfig` is a library for loading YAML configuration files and mapping data into a struct in Golang, using `Viper`. It supports environment variables and default values.

## 🚀 Installation

Add the library to your project:
```sh
go get -u github.com/anhnmt/go-defaultconfig
```

## 🛠️ Usage

### 1️⃣ **Define the Configuration Struct**
```go
type Config struct {
    Name  string `mapstructure:"name" default:"default_name"`
    Debug bool   `mapstructure:"debug" default:"false"`
}
```

### 2️⃣ **Create a Configuration File** (e.g., `config/dev.yml`)
```yaml
name: "my_app"
debug: true
```

### 3️⃣ **Load Configuration with Load**
```go
package main

import (
    "fmt"
    "log"
)

func main() {
    var cfg Config
    err := defaultconfig.Load("./config", "dev", &cfg)
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }
    fmt.Printf("Loaded config: %+v\n", cfg)
}
```

## 🔧 Load Function Details
```go
func Load(dir, env string, cfg any) error
```
### ✅ **Parameters**
- `dir`: Path to the directory containing the configuration files.
- `env`: Environment name (corresponding to `config/{env}.yml`).
- `cfg`: Pointer to the struct that receives the configuration data.

### 🔄 **How It Works**
1. **Determines the configuration file path** (`config/{env}.yml`).
2. **Reads the YAML file** and maps it into `cfg`.
3. **Supports environment variables** (converts `.` to `_`).
4. **Supports default values** from the struct tag `default`.

## ✅ Environment Variables Example
If environment variables are set:
```sh
export NAME="env_name"
export DEBUG=true
```
The program will use these values instead of those in the YAML file.

## 🧪 Running Unit Tests
```sh
# Run all tests
go test -v ./...
```

## 📜 License
MIT License
