# pagee

usage

```go
w := pagee.Walk{
    Uri:  "https://gocn.vip/topics/excellent",
    Next: ".pagination .next a",
    Item: ".topics .topic .title a",
    LimitItems: 61,
    LimitPages: 3,
}

w.Start(func(e *colly.HTMLElement) {
    // ....
})
```