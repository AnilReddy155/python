
import json
# jsonObj = '{"employees": ["name": "Anil", "Age": 24 },  {"name": "Anil", "Age": 24  }, { "name": "harish",
# "Age": 26 ' \ '}, {"name": "Tony", "Age": 30 }]} '
jsonObj = '{"employees": [{"name":"Anil","Age":24},{"name":"Anil", "Age": 23},{"name":"harish", "Age": 23}]}'
empList = json.loads(jsonObj)
uniqueNames = []
duplicateNames = []
size = len(empList["employees"])
for i in range(0, size):
    if empList["employees"][i]["name"] not in uniqueNames:
        uniqueNames.append(empList["employees"][i]["name"])
    else:
        duplicateNames.append(empList["employees"][i]["name"])

print("unique Names : ")
print(uniqueNames)

print("duplicate Names : ")
print(duplicateNames)
