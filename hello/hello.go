package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    }
    return getLangaugePrefix(language) + name
}

func getLangaugePrefix(language string) string {
    switch language {
    case "spanish":
        return "Hola, "
    default:
        return "Hello, "
    }
    if language == "spanish"{
        return "Hola, "
    } else {
        return "Hello, "
    }
}

func main() {
    fmt.Println(Hello("world", ""))
}
