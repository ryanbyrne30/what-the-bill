from datetime import datetime
from pack.fetch import Fetch
import logging
import urllib.parse


class API:
    def __init__(self, api_key: str, fetch: Fetch) -> None:
        self.fetch = fetch
        self.url_base = "https://api.govinfo.gov"
        self.headers = {"X-Api-Key": api_key}

    def __url(self, path: str, query: dict[str, any]) -> str:
        return self.url_base + path + "?" + urllib.parse.urlencode(query)

    def fetch_bills(
        self, from_date: datetime, limit: int = 1000, offset: int = 0
    ) -> list[any]:
        url_path = f"/collections/BILLS/{from_date.strftime('%Y-%m-%dT%H:%M:%SZ')}"
        url_query = {"pageSize": limit, "offset": offset}
        url = self.__url(url_path, url_query)
        data = self.fetch.json_request(url, headers=self.headers)
        if "packages" not in data:
            logging.error("[us_bills] 'packages' not found in result")
            return []
        return data["packages"]
