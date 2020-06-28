# ugg-scrapper
u.gg scrapper to make winrate graphs

## Install
Requires go installed

```
go install github.com/j4rv/ugg-scrapper
```

## How to use
Nidalee on her default role:

```
ugg-scrapper -champ nidalee
```

![Nidalee winrates](images/nidalee-default.png)

Annie on support:

```
ugg-scrapper -champ annie -role support
```

![Annie support winrates](images/annie-support.png)

## Libraries credits
- [github.com/gocolly/colly](https://github.com/gocolly/colly)
- [github.com/wcharczuk/go-chart.](https://github.com/wcharczuk/go-chart)