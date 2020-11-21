# go-pointers

This is a repo about pointers and derefrence operators in Golang.

## Intro - Value and Reference

Let's take this snippet as example.

```
x := 7
```

Here the value is 7 an integer.
And the reference here is the location where that value is being stored.

To see the reference of 7 (exact location) we are going to use the `&` operator.

```
fmt.Println(&x) // output: 0xc000018098
```

The result of this print statement (`0xc000018098`) is a memory location where the computer stores our value, this is our reference.

## Operators

### &

The `&` operator stands for <b>get the pointer</b>.

### \*

THE `*` operator stands for <b>derefrence</b>.
