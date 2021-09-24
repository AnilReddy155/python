lis = [{'make': 'Nokia', 'model': 216, 'color': 'Black'},
       {'make': 'Mi Max', 'model': 2, 'color': 'Gold'},
       {'make': 'Samsung', 'model': 7, 'color': 'Blue'},
       {'make': 'Samsung', 'model': 71, 'color': 'Blue'}
       ]

print("The list printed sorting by Make: ")
print(sorted(lis, key=lambda i: i['make']))

print("The list printed sorting by model: ")
print(sorted(lis, key=lambda i: i['model']))

print("The list printed sorting by color: ")
print(sorted(lis, key=lambda i: i['color']))

print("The list printed sorting by make and model: ")
print(sorted(lis, key=lambda i: (i['make'], i['model'])))
