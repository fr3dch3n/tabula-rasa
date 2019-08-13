# tabula-rasa
> Filter a list of tabs regarding the user-roles.

[![Build Status][travis-image]][travis-url]

## Installation

* with go-get: `go get github.com/fr3dch3n/tabula-rasa`
* with dep: `dep ensure --add github.com/fr3dch3n/tabula-rasa`


## Usage example

```go
package main

import (
    "fmt"
    tabularasa "github.com/fr3dch3n/tabula-rasa"
)

func main() {
    userRoles := []string{"my-awesome-role", "another-awesome-role"}
    allTabs := tabularasa.Tabs{Tabs:[]tabularasa.Tab{
        {"title", "esiLink", "initialState", []string{"my-aweseom-role"}},
        {"title", "esiLink", "initialState", []string{"not-allowed"}},
    }}
    
    tabsToRender := tabularasa.Tabs{Tabs: tabularasa.AllowedTabs(allTabs.Tabs, userRoles)}

    fmt.Println(tabsToRender)
}
```

## Release History

* 0.0.1
    * initial release

## Meta

[@fr3dch3n](https://twitter.com/fr3dch3n)

Distributed under the Apache 2.0 license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/fr3dch3n/tabula-rasa/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[travis-image]: https://img.shields.io/travis/fr3dch3n/tabula-rasa/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/fr3dch3n/tabula-rasa
