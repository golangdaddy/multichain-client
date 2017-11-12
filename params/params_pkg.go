package params

import (
    "strings"
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

        if len(line) == 0 { continue }

        if string(line[0]) == "#" { continue }

        kv := strings.Split(line, "=")

        k := strings.TrimSpace(kv[0])
        v := strings.TrimSpace(kv[1])

        params[k] = v

    }

    return params, nil
}
