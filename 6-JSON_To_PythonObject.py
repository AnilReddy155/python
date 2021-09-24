
import json

jsonObj = '{"Id": 1 ,"Ename":"David", "salary":1000, "Age":6 }'
empObj = json.loads(jsonObj)
print("JSON data: ")
print(empObj)
print("Name: ", empObj["Ename"])
print("Class: ", empObj["salary"])
print("Age: ", empObj["Age"])
