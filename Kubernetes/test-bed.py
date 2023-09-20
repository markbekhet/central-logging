import requests
import threading
import sys

add = sys.argv[1]
print(add)
def hitEndpoint():
    requests.get(add)

threads = []
for i in range(0, 10000000):
    thread = threading.Thread(target=hitEndpoint)
    thread.start()
    threads.append(thread)

for thread in threads:
    thread.join()