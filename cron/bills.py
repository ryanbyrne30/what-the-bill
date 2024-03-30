from scrape import Fetch, Proxy, BillScraper
import sys
import time
import re

with open("bills.txt", "r") as f:
    content = f.read()

saved_urls = re.findall("'url': '([^']+)'", content)

if __name__ == "__main__":
    proxy = Proxy(id=0, change_ip_after=5)
    proxy.run()
    time.sleep(10)

    fetch = Fetch(proxy=proxy) 
    scraper = BillScraper(fetch)   

    for line in sys.stdin:
        if line in saved_urls:
            continue
        bill = scraper.scrape(line.strip())
        print(bill.__dict__())

    proxy.kill()