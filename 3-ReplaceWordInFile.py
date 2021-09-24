with open("/home/anil/Desktop/wordr.txt") as file_data:
    file_content = file_data.read()

print(file_content.count("NAME"))
print(file_content.count("COMPANY"))
file_content = file_content.replace("NAME", "Anil")
file_content = file_content.replace("COMPANY", "IBM")

with open("/home/anil/Desktop/wordr.txt", 'w') as file_data:
    file_data.write(file_content)

print(file_content)
