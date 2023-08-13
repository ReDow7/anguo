# Anguo[WIP]

## Target
Some tools for company analysis to make a successful investment in Chinese stock market.

## Preparation
Some method implemented in this code is according to 《Financial Statement Analysis and Security Valuation》 written by Stephen H. Penman.

It's good if you are familiar with Chinese stock market, because the tool only works there.

It's important to have a token to query tushare database, which provides api for the tool. You can learn more about tushare on their website https://www.tushare.pro/.

## Usage

It's very simple to use it. Till now, we have three scene.

* All
```
go run main.go -token <your tushre token> -scene all
```
The program will scan all stocks in the market, compare the market value with evaluation, then output the details, like this:
```
Code    Name    Ratio   priceValue      Industry        Dividend        AlertInfo       saturation
601991.SH       大唐发电        87.80   54594.80m       火力发电        0.01    1_2     -0.03
600805.SH       悦达投资        71.37   3829.03m        综合类  0.00    1_2     -0.00
000725.SZ       京东方A 44.76   151216.82m      元器件  0.02    2       -0.01
002432.SZ       九安医疗        24.26   17309.73m       医疗保健        0.01    1       0.02
601919.SH       中远海控        20.51   157174.87m      水运    0.03    1       0.05
```
In the table output, we will show you the code&name&industry with some other details as well

* One

```
go run main.go -token <your tushre token> -scene one -code 002028.SZ
```
The same function as all, but just evaluate one company given by the code.
```
Code    Ratio   priceValue      odds    
002028.SZ       0.28    39.41b  -0.19
```

* Daily
```
go run main.go -token <your tushare token> -scene daily
```
The same function as all, but just evaluate companys given by the file "myholder.sec" in the root dictionary.
```
Code    Ratio   priceValue      odds    Name    
000651.SZ       1.47    202.22b 0.00    格力电器
000858.SZ       1.04    669.03b 0.00    五粮液
002032.SZ       0.96    40.02b  -0.00   苏泊尔
```

# End

Hope you like it, moreover, make more money by the tools. :)
