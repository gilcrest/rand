# rand

Random string generators with solutions using the math/rand or the crypto/rand library. The math/rand solution is based on Jon Calhoun's [fantastic post](https://www.calhoun.io/creating-random-strings-in-go/) and the crypto/rand solution is based on Matt Silverlock's [equally helpful and insightful post](https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/).

## crypto/rand

To get a random string using the crypto/rand package, use the following example:

```go
package main

import (
    "fmt"

    "github.com/gilcrest/rand"
)

func main() {
    randomString, err := rand.CryptoString(15)
    if err != nil {
        panic(err)
    }
    l := len(randomString)
    fmt.Printf("random string = %s with a length = %d\n", randomString, l)
    // Output: random string = QkQbyxjV4lqGkTfim8VR with a length = 20
}
```

## math/rand

To get a random string using the math/rand package, use the following example:

```go
package main

import (
    "fmt"

    "github.com/gilcrest/rand"
)

func main() {
    randomString := rand.MathString(20)
    l := len(randomString)
    fmt.Printf("random string = %s with a length = %d\n", randomString, l)
    // Output: random string = sCo2RC8PxUYTkJlIzxyZ with a length = 20
}
```

If you want control over the character set used for the random string, use the `MathStringWCharset` method. The default character set is simple Roman alpha-numerics (a-z, A-Z, 0-9). If you want to add a new set, check out the constants and submit a pull request to add a new one (along with the string set in the switch/case statement). 

Example using `MathStringWCharset`:

```go
package main

import (
    "fmt"

    "github.com/gilcrest/rand"
)

func main() {
    randomString := rand.MathStringWCharset(20, rand.Default)
    l := len(randomString)
    fmt.Printf("random string = %s with a length = %d\n", randomString, l)
    // Output: random string = cJ1UDuUYrnmdrjGRlqG2 with a length = 20
}
```