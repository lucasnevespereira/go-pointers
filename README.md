# go-pointers

This is a repo about pointers and derefrence operators in Golang.

## Value and Reference

Let's take this snippet as example.

```
myValue := 7
```

Here the value is 7 an integer.
And the reference here is the location where that value is being stored.

To see the reference of 7 (exact location) we are going to use the `&` operator.

```
fmt.Println(&myValue) // output: 0xc000018098
```

The result of this print statement (`0xc000018098`) is a memory location where the computer stores our value, this is our reference.

## Dereference

Let's take another code example.

```
myValue := 7
adressOfMyValue := &myValue
```

This means that the `adressOfMyValue` here it's equal to the pointer/reference of `myValue`.
<b>Attention: </b> This doesn't mean that `adressOfMyValue` it's equal to `7`, but that `adressOfMyValue` it's equal to the pointer/reference of 7, so `adressOfMyValue` it's equal to `0xc000018098`.

What if I wanted to change the value stored in that pointer (7) to 8 for example.

Well to modify the value of a pointer/reference, we need to use the `*` operator to <b> derefrence</b>.

```
myValue := 7
adressOfMyValue := &myValue

fmt.Println(myValue, adressOfMyValue) // outputs: 7, 0xc000018098

*y = 8

fmt.Println(myValue, adressOfMyValue) // outputs: 8, 0xc000018098
```

`myValue` is now <b>8</b>.

<hr>

### Operators

### &

The `&` operator stands for <b>get the pointer</b>.

### \*

THE `*` operator stands for <b>derefrence</b>.
