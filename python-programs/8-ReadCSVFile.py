import csv

lineList = []
with open("/home/anil/Desktop/csvfile.csv") as file_data:
    csv_file_data = csv.reader(file_data, delimiter=',')
    print(type(csv_file_data))
    # headers
    print(next(csv_file_data))
    for line in csv_file_data:
        lineList.append(line)
        print("id : "+line[0])
        print("Name : "+line[1])
        print("Age : " + line[2])
        print(" ")

# with open("/home/anil/Desktop/read.csv") as file_data:
#    for line in file_data:
#        lineList.append(line.rstrip())

print(lineList)
