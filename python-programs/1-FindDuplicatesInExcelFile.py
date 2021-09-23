import pandas as pd

data = pd.read_excel("/home/anil/Desktop/Employee.XLSX")
count = 0
# print(data.duplicated("Emp Name"))
print(type(data))
df = pd.DataFrame(data)
print(df)
print(df.groupby(["Emp Name", "Emp Salary"]).size())
print(data.drop_duplicates())
# print(data)
