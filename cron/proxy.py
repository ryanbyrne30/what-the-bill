import subprocess
from subprocess import Popen 
import requests 
import time

def create_tor_proxy(id: int) -> Popen:
    file_dst = f"/tmp/torrc.{id}"
    file_contents = f'''
SocksPort {9060 + id}
ControlPort {9060 + id + 1}
DataDirectory /tmp/tor.{id} 
'''
    with open(file_dst, "w") as f:
        f.write(file_contents)

    print("Creating tor socket...")
    return subprocess.Popen(['tor', '-f', file_dst])


def print_ip():
    print("Fetching ip...")
    http_proxy  = "socks5://localhost:9060"
    proxies = { 
        "http"  : http_proxy, 
        "https" : http_proxy, 
    }
    response = requests.get("http://checkip.amazonaws.com/", proxies=proxies)
    ip = response.content.strip()
    print(ip)


# process = Process(target=create_tor_proxy, args=[0])
# process.start()
process = create_tor_proxy(0)
print("Sleeping 30...")
time.sleep(30)
print_ip()
process.terminate()
