# fjis
fjis is a command line program for verification that **standard input can be encoded to Shift_JIS**.


Character(s) cannot be encoded, output **highlighted** that character(s) and **Unicode code point** by the corresponding that character(s).


![console image](console.png)

## Usage

```
$ git clone https://github.com/ShingoYadomoto/fjis.git
$ go build -o fjis main.go
$ echo 'aÈaÄa' | ./fjis
```
