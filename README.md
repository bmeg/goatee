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

### Example 2 Output
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


## Example 2

Because the pattern for unrolling input arrays returns a dictionary, there needs to be a way to merge multiple sub-structure outputs into a single dictionary.

For the input 
```
{
    "inputs" : [
        {"name" : "job1", "state": "queued"},
        {"name" : "job2", "state": "waiting"},
        {"name" : "job3", "state": "running"},
        {"name" : "job4", "state": "done"}
    ]
}
```

and the template 

```
{
    "{{#each inputs}}" : {
		"task_{{name}}" : {
			"state" : "{{state}}"
		}
	}
}
```

we get the results

```
[
    {"task_job1" : {"state": "queued"} },
    {"task_job2" : {"state": "waiting"} },
    {"task_job3" : {"state": "running"} },
    {"task_job4" : {"state": "done"} }
]
```

If theese elements belong to a `{{#merge}}` they will be merged into a single dictionary.

So the template 
```
{   
    "{{#merge}}" : {
        "{{#each inputs}}" : {
            "task_{{name}}" : {
                "state" : "{{state}}"
            }
        }
    }
}
```

Results in:
```
{
    "task_job1" : {"state": "queued"},
    "task_job2" : {"state": "waiting"},
    "task_job3" : {"state": "running"},
    "task_job4" : {"state": "done"}
}

```
