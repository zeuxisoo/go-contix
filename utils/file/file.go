package file

import (
    "os"
    "io"
    "bufio"
    "bytes"
)

func ReadByLines(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    var buffer bytes.Buffer
    var lines []string
    for {
        line, prefix, err := reader.ReadLine()
        if err != nil {
            break
        }

        buffer.Write(line)

        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }

    if err == io.EOF {
        return nil, err
    }

    return lines, nil
}
