from scrape import BillsSitemapScraper, Fetch

if __name__ == "__main__":
    ofile = "billing-urls.txt"
    url = "https://www.govinfo.gov/sitemap/BILLS_2024_sitemap.xml"

    fetch = Fetch()
    scraper = BillsSitemapScraper(fetch)    

    urls = scraper.get_bill_urls(url)

    print('\n'.join(urls))
