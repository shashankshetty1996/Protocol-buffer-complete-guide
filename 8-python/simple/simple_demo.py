import simple_pb2

simple_message = simple_pb2.SimpleMessage()
simple_message.id = 123
simple_message.is_simple = True
simple_message.name = "This is a simple message"

simple_list = simple_message.sample_list
simple_list.append(3)
simple_list.append(4)
simple_list.append(5)

print(simple_list)
print(simple_message)

with open("simple.bin", "wb") as f:
  print("write binary")
  bytesAsStrings = simple_message.SerializeToString()
  f.write(bytesAsStrings)

with open("simple.bin", "rb") as f:
  print("read binary")
  simple_message_read = simple_pb2.SimpleMessage().FromString(f.read())

print(simple_message_read)
print("Is simple? :", str(simple_message.is_simple))
print("Id? :", str(simple_message.id))