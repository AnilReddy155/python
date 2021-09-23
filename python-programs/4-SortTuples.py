subMarks = [('English', 88), ('Science', 90), ('Maths', 97), ('Social sciences', 82)]

print("Before Sorting the Tuple! ")
print(subMarks)


# sorting Tuple using Lambda
subMarks.sort(key = lambda x: x[1])
print("After Sorting the Tuples")
print(subMarks)

# without Lambda
# def sortTuple(tuples):
#     length = len(tuples)
#     for i in range(0, length):
#
#         for j in range(0, length - i - 1):
#             if tuples[j][-1] > tuples[j + 1][-1]:
#                 temp = tuples[j]
#                 tuples[j] = tuples[j + 1]
#                 tuples[j + 1] = temp
#
#     return tuples
#
#
# sortTuple(subMarks)
print("After Sorting")
print(subMarks)