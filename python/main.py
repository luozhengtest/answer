# !user/bin
# -*- coding: utf-8 -*- 
# @Time : 2023/4/21 上午11:39 
# @Author : luozhengtest 
# @File : main.py
from python import person_pb2


address_book = person_pb2.AddressBook()
person = address_book.people.add()
person.id = 123123
person.name = "zheng.luo"
person.email = "123@test.com"
phone = person.phones.add()
phone.number = "17711077426"
phone.type = person_pb2.Person.HOME

binary = "binary/persion"

f = open(binary, "wb")
f.write(address_book.SerializeToString())
f.close()

f = open(binary, "rb")
info = address_book.ParseFromString(f.read())
f.close()
