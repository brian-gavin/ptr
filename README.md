# ptr

ptr is a go library for creating pointers to objects inline.

```go
import (
    "fmt"

    "github.com/brian-gavin/ptr"
)

func main() {
    type s struct {
        p *int
    }
    myS := s{
        p: ptr.To(10)
    }
    fmt.Println(s.p)
}
```