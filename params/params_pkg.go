package params

import (
    "strings"
    "strconv"
    "io/ioutil"
)

type Params map[string]string

func Open(pathToFile string) (Params, error) {

    params := Params{}

    b, err := ioutil.ReadFile(pathToFile)
    if err != nil {
        return nil, err
    }

    blob := string(b)

    for _, line := range strings.Split(blob, "\n") {

        line = strings.TrimSpace(line)

        parts := strings.Split(line, "#")

        if len(parts[0]) == 0 { continue }

        kv := strings.Split(strings.TrimSpace(parts[0]), "=")

        k := strings.TrimSpace(kv[0])
        v := strings.TrimSpace(kv[1])

        params[k] = v

    }

    return params, nil
}

// Params methods

func (params Params) Bool(key string) bool {

    value := params[key]

    switch value {

        case "true":

            return true

        case "false":

            return false

    }

    panic("Invalid BOOL value for key: "+key+" - "+value)

    return false
}

func (params Params) Int(key string) int {

    value := params[key]

    i, err := strconv.Atoi(value)
    if err != nil {
        panic(err)
    }

    return i
}

func (params Params) Float64(key string) int {

    value := params[key]

    i, err := strconv.ParseFloat(value, 64)
    if err != nil {
        panic(err)
    }

    return i
}

func (params Params) String(key string) string {

    return params[key]
}
