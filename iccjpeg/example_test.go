package iccjpeg_test

import (
    "iccjpeg"
    "log"
    "os"
)

func ExampleGetICCBuf() {
    /* Open input. */
    input, err := os.Open("/path/to/my.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer input.Close()

    /* Grab the profile buffer. */
    profile, err = iccjpeg.GetICCBuf(input)
    if err != nil {
        log.Println("ICC profile search failed: ", err)
    }

    if len(profile) == 0 {
        log.Println("No ICC profile found.")
    } else {
        log.Println("ICC profile of length ", len(profile), " found.")
    }
}
