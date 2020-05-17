import complex_pb2

complex_message = complex_pb2.ComplexMessage()

# This is not required
# one_dummy_message = complex_pb2.DummyMessage()
# one_dummy_message.id = 123
# one_dummy_message.name = "I am a dummy message"
# complex_message.one_dummy = one_dummy_message

complex_message.one_dummy.id = 123
complex_message.one_dummy.name = "I am a dummy message"

first_multiple_message = complex_message.multiple_dummy.add()
first_multiple_message.id = 345
first_multiple_message.name = "I am first element on an array"

complex_message.multiple_dummy.add(id=567, name="Second element")

# Doesn't do copy
third_dummy_message = complex_pb2.DummyMessage()
third_dummy_message.id = 999
third_dummy_message.name = "I'm last one!"
complex_message.multiple_dummy.extend([third_dummy_message])

third_dummy_message.id = 111
print(third_dummy_message)

print(complex_message)