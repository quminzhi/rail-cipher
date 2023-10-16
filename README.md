# Rail Cipher Text

This is simple Rail Fence Cipher lib.

For example, the string "WEAREDISCOVEREDFLEEATONCE" could be represented in a three rail system as follows:

```
W       E       C       R       L       T       E
  E   R   D   S   O   E   E   F   E   A   O   C  
    A       I       V       D       E       N
```

The encoded string would be: "WECRLTEERDSOEEFEAOCAIVDEN"

## How to use

```go
     got := Encode("WEAREDISCOVEREDFLEEATONCE",3) // got = "WECRLTEERDSOEEFEAOCAIVDEN"
        got = Decode("WECRLTEERDSOEEFEAOCAIVDEN",3) // got = "WEAREDISCOVEREDFLEEATONCE"
```
