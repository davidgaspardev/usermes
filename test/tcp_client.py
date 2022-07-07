import socket
import time

def create_tcp_client():
    HOST = 'localhost'
    PORT = 8888

    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    dest = (HOST, PORT)

    tcp.connect(dest)
    read(tcp)
    # index = 0
    # while True:
    #     if index == 20:
    #         tcp.send(b"David\n")
    #         break
    #     tcp.send(b"Hello")
    #     time.sleep(2)
    #     index += 1

    tcp.close()

def read(tcp):
    while True:
        data = tcp.recv(1024)
        print(data)

create_tcp_client()