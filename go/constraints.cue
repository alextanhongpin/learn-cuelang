// https://cuelang.org/docs/howto/

import "list"
import "time"
import "strings"

let url = =~ "^https://(w+)"

#User: close({
    name!: string & strings.MinRunes(2) & strings.MaxRunes(8)
    age!: >= 13
    hobby!: [...string] & list.MinItems(1) & list.MaxItems(1)
    birthday!: string & time.Format("2006-01-02")
    imageURL: string & url
})

{
    name: "john",
    age: 13,
    hobby: ["programming"],
    birthday: "1970-01-01",
    imageURL: "https://test.com"
}  & #User
