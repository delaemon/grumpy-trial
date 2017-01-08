import socket
import select

read_waits = {}
write_waits = {}

def wait_read(con, callback):
    read_waits[con.fileno()] = callback

def wait_write(con, callback):
    write_waits[con.fileno()] = callback

def evloop():
    while 1:
        rs, ws, xs, = select.select(read_waits.keys(), write_waits.keys(), [])
        for rfd in rs:
            read_waits.pop(rfd)()
        for wfd in ws:
            write_waits.pop(wfd)()

class Server(object):
    def __init__(self, con):
        self.con = con

    def start(self):
        wait_read(self.con, self.on_acceptable)

    def on_acceptable(self):
        try:
            while 1:
                con, _ = self.con.accept()
                con.setblocking(0)
                Client(con)
        except IOError:
            wait_read(self.con, self.on_acceptable)

class Client(object):
    def __init__(self, con):
        self.con = con
        self.on_readable()

    def on_readable(self):
        data = self.con.recv(32 * 1024)
        if not data:
            wait_read(self.con, self.on_readable)
            return
        self.buf = b"""HTTP/1.1 200 OK \r
Content-Type: text/plain\r
Content-Length: 6\r
\r
hello
"""
        self.on_writable()

    def on_writable(self):
        wrote = self.con.send(self.buf)
        self.buf = self.buf[wrote:]
        if self.buf:
            wait_write(self.con, self.on_writable)
        else:
            self.con.close()

def server():
    sock = socket.socket()
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.setblocking(0)
    sock.bind(('', 8000))
    sock.listen(128)
    server = Server(sock)
    server.start()
    evloop()

if __name__ == '__main__':
    server()
