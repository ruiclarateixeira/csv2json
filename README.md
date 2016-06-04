# CSV2JSON - GO Library

**Turn csv like**
```csv
header1,header2,header3
cell1,cell2,cell3
cell4,cell5,cell6
```

**Into a json array like**
```json
[
	{
		"header1": "cell1",
		"header2": "cell2",
		"header3": "cell3"
	},
	{
		"header1": "cell4",
		"header2": "cell5",
		"header3": "cell6"
	}
]
```

**By calling**
```GO
csv2json.ReadFile(fileName)
```