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

blablabla

### Decryption

bala

## How to use

Prerequisite:

    - Basic environment for docker engine.
    - Basic knowledge of Go.

To simplify use, we containerize the project and encapsulate useful commands
into a makefile.

```bash
> make up   # build an image and run rail container
> make test # run test
> make down # clean rail container and release resources
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
