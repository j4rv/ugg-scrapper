# ugg-scrapper
u.gg scrapper to make winrate graphs

## Install
Requires go installed

```
go install github.com/j4rv/ugg-scrapper/bin/ugg-scrapper
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

Volibear at patch 10.12 (default role: jungle):

```
ugg-scrapper -champ volibear -patch 10_12
```

![Volibear winrates at patch 10.12](images/volibear-default-10_12.png)

## Libraries credits
- [github.com/gocolly/colly](https://github.com/gocolly/colly)
- [github.com/wcharczuk/go-chart.](https://github.com/wcharczuk/go-chart)