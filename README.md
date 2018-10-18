# go-seo
This tool will mock browser to click your page that found in result of Google.

### How to work
If your config.json has "items' property, this tool will use your setting to open "google search page". Then it will use the "pattern" property of the item to search matched url. 
If it can find the url, it will click this url like browser's behavior. So this will have chance to improve google result ranking of your web page.

### The Content of Config.json
```
{
    "arg" : {
        "agent" : "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
        "clickSleepSecs" : 5,
        "nextPagePattern" : "google.*&start=[0-9]+&sa="
    },
    "items" : [
        {
            "search" : "https://www.google.com.tw/search?q=x86+assembly+%E7%8C%9C%E6%95%B8%E5%AD%97&oq=x86+assembly+%E7%8C%9C%E6%95%B8%E5%AD%97",
            "pattern" : "http://tylpk.blogspot.com/2018/09/x86-assembly-in-linux-part-5.html"
        },
        {
            "search" : "https://www.google.com.tw/search?q=uefi+python&oq=uefi+python",
            "pattern" : "uefi-application-python-272.html"
        }
    ]
}
```

### Screenshot
```
page pattern :  google.*&start=[0-9]+&sa=
search  : https://www.google.com.tw/search?q=x86+assembly+%E7%8C%9C%E6%95%B8%E5%
AD%97&oq=x86+assembly+%E7%8C%9C%E6%95%B8%E5%AD%97
pattern : http://tylpk.blogspot.com/2018/09/x86-assembly-in-linux-part-5.html
result  : clicked
error   : <nil>

search  : https://www.google.com.tw/search?q=uefi+python&oq=uefi+python
pattern : uefi-application-python-272.html
result  : clicked
error   : <nil>
```
