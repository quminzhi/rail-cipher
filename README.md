# Rail Cipher Text

This is simple Rail Fence Cipher lib.

For example, the string "WEAREDISCOVEREDFLEEATONCE" could be represented in a three rail system as follows:

```
W       E       C       R       L       T       E
  E   R   D   S   O   E   E   F   E   A   O   C  
    A       I       V       D       E       N
```

The encoded string would be: "WECRLTEERDSOEEFEAOCAIVDEN"

## Algorithm

### Encryption

The rail cipher encryption process with a depth of "d" and repetition of "r"
involves the following steps:

- Initialize a 2D matrix with "d" rows and a length corresponding to the
plaintext message's character count.
- Begin at the top-left corner of the matrix and start filling it with
characters from the plaintext in a zigzag pattern, moving either down or
up and changing direction when reaching the top or bottom rail.
- Repeat this zigzag pattern until the entire plaintext is filled in the
matrix.
- Create the ciphertext by reading the characters from the matrix row by row.
- If the "repetitions" parameter is greater than 1, repeat the encryption
process "r" times by using the ciphertext obtained from the previous iteration
as the new plaintext.

The depth determines the number of rails or rows, and the repetitions control
how many times the encryption process is repeated. The resulting ciphertext is
the rail cipher, which appears as a series of characters that are encrypted
using this zigzag pattern.

### Decryption

A rail cipher decryption process involves reversing the encryption process.
The depth (d) and the number of repetitions (r) are key parameters in this
process. Below is a pseudocode for decrypting a rail cipher:

```go
Function Decode(ciphertext, key) {
    depth = key[0]
    repetitions = key[1]

    for i from 1 to repetitions {
        decryptedtext = DecodeOnce(ciphertext, depth)
        ciphertext = decryptedtext
    }
}

Function DecodeOnce(ciphertext, depth) {
    Create an empty 2D array railMatrix with depth rows and a length of ciphertext columns.
    Create variables currentRow and currentCol and initialize them to 0.
    Create a variable direction and initialize it to 1. # 1 for moving down, -1 for moving up.

    // Initialize railMatrix with placeholders
    For row in railMatrix {
        For column in railMatrix[row] {
            railMatrix[row][column] = '*'
        }   
    }
    
    // Fill railMatrix with ciphertext characters
    For each character in ciphertext {
        railMatrix[currentRow][currentCol] = character
        currentCol += 1
        If currentRow == 0 or currentRow == depth - 1 {
            direction *= -1
        }
        currentRow += direction
    }

    // Create a list of characters from railMatrix
    decryptedText = ""
    For row in railMatrix {
        For column in railMatrix[row] {
             If railMatrix[row][column] != '*' {
                decryptedText += railMatrix[row][column]
            }
        }
    }

    Return decryptedText
}
```

## How to use

Prerequisite:

- Basic environment for docker engine.
- Basic knowledge of Go.

To simplify use, we containerize the project and encapsulate useful commands
into a makefile.

```bash
> make up   # build an image and run rail container
> make test # run test for packages
> make run  # jump into container and run cipher command
> make down # clean rail container and release resources
```

Here's an example for run cipher command in the container:

```bash
# by make up, we build, run a container, and jump into it
> make up
root@97f6fca7915f:/workspace#
root@29c040e9c001:/workspace# ./cipher -help
Usage of ./cipher:
  -depth int
     depth of rail fence encryption (default 2)
  -help
     display usage information
  -message string
     input message (default "hello world")
  -option string
     encryption or decryption (default "encryption")
  -repeat int
     the number of times to repeat encryption or decryption (default 1)
root@29c040e9c001:/workspace# 

root@29c040e9c001:/workspace# ./cipher -option encryption -message "hello world" -depth 3 -repeat 2
Ciphertext: hlloe lwrod
root@29c040e9c001:/workspace# ./cipher -option decryption -message "hlloe lwrod" -depth 3 -repeat 2
Plaintext: hello world

# test rail package
root@29c040e9c001:/workspace# cd rail/
root@29c040e9c001:/workspace/rail# ls
go.mod rail.go  rail_test.go
root@29c040e9c001:/workspace/rail# go test -v .
=== RUN   TestEncode
--- PASS: TestEncode (0.00s)
=== RUN   TestDecode
--- PASS: TestDecode (0.00s)
PASS
ok   github.com/quminzhi/rail 0.004s
root@29c040e9c001:/workspace/rail# <ctrl-d>

# close the container and release resources
> make down
rail
rail
Container rail ended.
```

We have 4 tests as shown in `rail_test.go`,

```go
// Test cases for encryption
testcases := []struct {
  name      string
  plaintext string
  key       []int
  expected  string
 }{
  {"test0", "WEAREDISCOVEREDFLEEATONCE", []int{3, 1}, "WECRLTEERDSOEEFEAOCAIVDEN"},
  {"test1", "Hello, World!", []int{3, 1}, "Hoo!el,Wrdl l"},
  {"test2", "Hello, World!", []int{3, 2}, "Herlo!lWd o,l"},
  {"test_homework",
   "CRYPTOLOGY IS THE PRACTICE AND STUDY OF TECHNIQUES FOR SECURE COMMUNICATION IN THE PRESENCE OF THIRD PARTIES CALLED ADVERSARIES",
   []int{4, 5},
   "CT HEGCONAYQNENMDD FRLMUET EIIEE ISTOPCSRT CC FLRSIPVI TO IR  TH  OSFAAHNCOTC  DSRDRN SAIUY CDRTEAPHA CUPEEAEIOEOLSIETNEES YRRU"},
 }

// Test cases for decryption
testcases := []struct {
  name       string
  ciphertext string
  key        []int
  expected   string
 }{
  {"test0", "WECRLTEERDSOEEFEAOCAIVDEN", []int{3, 1}, "WEAREDISCOVEREDFLEEATONCE"},
  {"test1", "Hoo!el,Wrdl l", []int{3, 1}, "Hello, World!"},
  {"test2", "Herlo!lWd o,l", []int{3, 2}, "Hello, World!"},
  {"test_homework",
   "TPSNIONFRMHANOIREEOIBTSEAKLAPSEHISOBPEGTBRQREPTTMHRTHTTAWEAACTFVAECAOLHANSEEKFHALOITUEEAICNLEYOLTOLEPADFKATATSJMSIINSHROCTITRIEEAKYNHYUEOTTSTATSIIRSARERUYUEDPCLETEROINEIYEHC",
   []int{3, 3},
   "THERAILFENCECIPHERISANEASYTOAPPLYTRANSPOSITIONCIPHERTHATJUMBLESUPTHEORDEROFTHELETTERSOFAMESSAGEINAQUICKCONVENIENTWAYITALSOHASTHESECURITYOFAKEYTOMAKEITALITTLEBITHARDERTOBREAK"},
 }
```
