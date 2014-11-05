package iccjpeg

import (
    "crypto/md5"
    "fmt"
    "net/http"
    "testing"
)

func TestICCGetBuf(t *testing.T) {
    tests := map[string]string {
        "http://www.color.org/Upper_Left.jpg" : "099e4048d0bfd95338b4bdf38ac1eaf4",
        "http://www.color.org/Upper_Right.jpg": "bfe2da5cfcffbf3ebdd7d096bfe430c0",
        "http://www.color.org/Lower_Left.jpg" : "32d74636855936433e215571df4b6173",
        "http://www.color.org/Lower_Right.jpg": "50b9125494306a6fc1b7c4f2a1a8d49d",
    }

    for url, hash := range tests {
        req, err := http.Get(url)
        if err != nil {
            t.Errorf("URL Unavailable: %s", err)
        }

        buf, err := GetICCBuf(req.Body)
        if err != nil {
            t.Error(err)
        }
        req.Body.Close()

        hashstr := fmt.Sprintf("%x", md5.Sum(buf))
        if hashstr != hash {
            t.Errorf("ICC buffer doens't match: %s != %s", hashstr, hash)
        }
    }
}
