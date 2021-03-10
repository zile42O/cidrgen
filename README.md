[![cidrgen](https://img.shields.io/badge/cidrgen-2f2f2f.svg?style=for-the-badge)](https://github.com/zile42O)

Simple programs for generating CDIR ip list

##  How to use?

- Clone this repository with simple command
```bash
git clone https://github.com/zile42O/cidrgen
```
- Go to `cidrgen/asn`, you need to have file.json, in which be lookup from (i recommend) shodan.io API, Request (https://api.shodan.io/shodan/host/search?key={YOUR_API_KEY}&query={query}&facets={facets}
), :warning: **You need Shodan API Key for this**, btw there is already file which you can test how it works
- Example query for ASN ( https://api.shodan.io/shodan/host/search?key={YOUR_API_KEY}&query=linksys ) - I use this for colleting ASN from linksys, when you get your json request, download it, and replace or create file in `cidrgen/asn/file.json`

- Then you need to run this commands:

```bash
cd .\cidrgen\asn\ (if you are not already in this dir)
.\main.exe (this will create asn_list.txt without duplicates)
```
- After that you need to change dir, use this commands:
```bash
cd ..\lookup_asn_cidr\
.\main.exe (after change dir, start the program and you will get output cidr_list.txt)
```
