# Goatee

* Mustache templates for JSON structures. *

Goatee takes the form of the [mustache](https://mustache.github.io/) templating language, but applies it to nested JSON data structures. 
The goatee parser uses a template JSON structure, applying the mustache template parser to every string field. 

Goatee can also expand nested elements into arrays based on the template structure. If the goatee parser finds a lone `#each` command as a dictionary 
key, it will take the field value and replace dictionary with an expanded version of the data. 
When expanding a structure, such as an array of strings, the varible name becomes `this`.
For example the structure `{ "{{#each names}}" : { "name" : "{{this}}" } }` with the data `{"names" : ["Alice", "Bob", "Chuck"]}` would produce the output
`[{"name":"Alice"}, {"name":"Bob"}, {"name":"Chuck"}]`. 


## Example 1

### Example 1 Input 
```
{
    "names" : ["test1", "test2", "test3"]
}
```

### Example 1 Template
```
{
    "name": "job1",
    "inputs" : {
        "{{#each names}}" : "/home/{{this}}.txt"
    }
}
```

### Example 2 Template
```
{
    "name": "job1",
    "inputs" : [
        "/home/test1.txt",
        "/home/test2.txt",
        "/home/test3.txt"
    ]
}
```
