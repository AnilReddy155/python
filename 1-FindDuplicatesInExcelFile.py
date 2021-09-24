import pandas as pd

data = pd.read_excel("/home/anil/Desktop/Employee.XLSX")
count = 0
# print(data.duplicated("Emp Name"))
print(type(data))
df = pd.DataFrame(data)
# print(df.set_index("Emp Name"))
# print(df.set_index("Emp Name").T.to_dict())
print(df.groupby(["Emp Name", "Emp Salary"]).size())
print(data.drop_duplicates())
print(data)
