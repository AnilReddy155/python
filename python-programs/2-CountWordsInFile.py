# search and count the words in file

count = 0
with open("/home/anil/Desktop/py_file.txt") as file_data:
    file_content = file_data.read()
    file_data.seek(0)
    for line in file_data:
        for s in line.rstrip().split(" "):
            count += 1
#       print(line.rstrip())
#   for line in file_content:
#      print(line)

# print(file_content)

# for s in file_content.split(" "):
#    count += 1

print(count)
