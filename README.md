# GI Joe PSA Moral API

### Returns a certain number of GI Joe PSA lessons

*All lessons sourced from [this page](http://www.joeheadquarters.com/joeendings.shtml)*

### GET `/`

Requesting the root path will return the most recent version of the API response

### GET `/v1` / GET `/v1.json`

#### Params

`:num` - Number of responses. Defaults to 1 and will return the max if a greater number is specified

#### Returns

* `:text` - The text of the lesson
* `:season` - The season in which this PSA aired
* `:code_name` - The name of the GI Joe who delivered this lesson
* `:situation` - The main situation for this PSA

**Note: These attributes will be in an array if `num` is greater than 1**

### Running

```bash
$ go build && PORT=5000 ./gi_joe_psa
```
