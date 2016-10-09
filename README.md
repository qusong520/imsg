# imsg
Append message to image(or any other file)

[TOC]



## Usage

### Encrypt

Append text message to a file



Example:

```shell
imsg -i a.jpg -o b.jpg "What is the time?"
```



### Decrypt

Get the message from the file



Example:

```shell
imsg -d b.jpg
```

or

```shell
imsg -d -i b.jpg
```



## Flag

| name | description                              | e.g.     |
| ---- | ---------------------------------------- | -------- |
| -d   | decrypt mode (without this, default encrypt mode) |          |
| -i   | the input file path                      | a.jpg    |
| -o   | the output file path                     | b.jpg    |
| -s   | the split text between the origin file text and message(default: #########) | $$$$$$$$ |