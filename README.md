# battlerite-node
A Go SDk for accessing the official Battlerite API.
## Getting started
To start using battlerite-go, you just have to install the package:
```
go get github.com/sime1/battlerite-go
```
if you want to update the package:
```
go get -u github.com/sime1/battlerite-go
```
## Basic usage
First you need to import the package:
```
import(
  "github.com/sime1/golang/battlerite"
)
```
then inside your function you create a Client using your API key:
```
client := battlerite.Client{"your_API_key_goes_here"}
```
To retrieve a list of matches:
```
matches,err := client.GetMatchList(nil, "", nil)
if err != nil {
  fmt.Printf("Error: %+v", err)
} else {
  doSomethingWithMatches(matches)
}
```
you can also choose which matches to retrieve:
```
matches,err := client.GetMatchList(&battlerite.Page{Limit:2},
                                   "-createdAt",
                                   &battlerite.MatchFilter{PlayerIds: []string{"playerName1", "another_name"}})
...
```
To retrieve a single match:
```
match,err := client.GetMatch("matchIdHere")
```
To fetch telemetry data:
```
match,err := client.GetMatch(id)
if err == nil {
  telemetry,err := match.GetTelemetry()
  if err == nil {
    doSomethingWithTelemetry(telemetry)
  }
}
```
## Documentation
For detailed information about the usage of battlerite-go, you can just run godoc and look at the generated documentation in your browser:

```
godoc -http=:8080
```
then browse to http://localhost:8080
## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
## Changelog
0.1.0: initial Release

0.2.0: Player functionalities + README

0.2.1: minor bug fixes
