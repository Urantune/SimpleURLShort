##Tech 
- LanguagePrograming: Go (1.25)
- Framework: Gin

Simple URL Shortener.

+ Create Code: 
curl http://localhost:8080/api/short `
-Method POST `
-Headers @{ "Content-Type" = "application/json" } `
-Body '{ "url": "link here" }'
(use power shell)


+ Use code 
../:code


+ Full list 
../api/links

+ Check code
/api/stats/:code

+ Create Link by broser
../api/shortByLink
