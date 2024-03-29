from package.services import bill_scraper, bill_sitemap_scraper


URL = "https://www.congress.gov/search?q=%7B%22source%22%3A%22legislation%22%2C%22congress%22%3A118%7D"
list_selector = "ol.basic-search-results-lists"

urls = bill_sitemap_scraper.get_bill_urls("https://www.govinfo.gov/sitemap/bulkdata/BILLSUM/118hres/sitemap.xml")
for url in urls:
    print(url)
