// cmd/artificialeaselsdknext/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaselsdknext/internal/artificialeaselsdknext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaselsdknext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
