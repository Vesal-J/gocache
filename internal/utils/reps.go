package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ToRESP(args ...string) []byte {
	var b []byte
	b = append(b, '*')
	b = append(b, []byte(fmt.Sprintf("%d\r\n", len(args)))...)
	for _, arg := range args {
		b = append(b, '$')
		b = append(b, []byte(fmt.Sprintf("%d\r\n%s\r\n", len(arg), arg))...)
	}
	return b
}

func ParseRESP(input string) ([]string, error) {
	var result []string
	reader := bufio.NewReader(strings.NewReader(input))

	line, err := reader.ReadString('\n')
	if err != nil || len(line) == 0 || line[0] != '*' {
		return nil, fmt.Errorf("invalid RESP array header")
	}

	count, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return nil, fmt.Errorf("invalid array length: %v", err)
	}

	for i := 0; i < count; i++ {
		prefix, err := reader.ReadByte()
		if err != nil || prefix != '$' {
			return nil, fmt.Errorf("expected bulk string")
		}

		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		strLen, err := strconv.Atoi(strings.TrimSpace(lenLine))
		if err != nil {
			return nil, err
		}

		buf := make([]byte, strLen+2) // +2 for \r\n
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return nil, err
		}

		result = append(result, string(buf[:strLen]))
	}

	return result, nil
}

func EncodeRESP(value any) ([]byte, error) {
	switch v := value.(type) {

	case string:
		// Treat as Bulk String
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)), nil

	case []string:
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("*%d\r\n", len(v)))
		for _, item := range v {
			buf.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(item), item))
		}
		return buf.Bytes(), nil

	case int, int64:
		return []byte(fmt.Sprintf(":%v\r\n", v)), nil

	case error:
		return []byte(fmt.Sprintf("-%s\r\n", v.Error())), nil

	case nil:
		return []byte("$-1\r\n"), nil // RESP Null Bulk String

	case []any:
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("*%d\r\n", len(v)))
		for _, item := range v {
			encoded, err := EncodeRESP(item)
			if err != nil {
				return nil, err
			}
			buf.Write(encoded)
		}
		return buf.Bytes(), nil

	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}
